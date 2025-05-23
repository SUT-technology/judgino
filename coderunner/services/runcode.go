package services

import (

	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"

	"time"

	"github.com/SUT-technology/judgino/coderunner/dto"
)

const (
    StatusOK = iota + 3
    StatusWrongOutput
    StatusCompileError
    StatusRuntimeError
    StatusTimeLimitExceeded
    StatusMemoryLimitExceeded
)

func (c RunnerServices) RunCode(submission dto.SubmissionRun) error {
	id := submission.ID
	code := submission.Code
	input := submission.Input
	wantOutput := submission.ExpectedOutput
	timeLimit := time.Duration(submission.TimeLimit) * time.Second
	memLimitMB := submission.MemoryLimit

	wantOutput += "\n"

    tmpDir, err := ioutil.TempDir("", "go-run-")
    if err != nil {
        c.SendResult(id, StatusRuntimeError)
		return err
    }
    defer os.RemoveAll(tmpDir)

    srcPath := filepath.Join(tmpDir, "main.go")
    if err := ioutil.WriteFile(srcPath, []byte(code), 0644); err != nil {
		c.SendResult(id, StatusRuntimeError)
        return err
    }
	

	srcPath2 := filepath.Join(tmpDir, "input.in")
    if err := ioutil.WriteFile(srcPath2, []byte(input), 0644); err != nil {
		c.SendResult(id, StatusRuntimeError)
        return err
    }



    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
    
    resp, err := cli.ContainerCreate(context.Background(), &container.Config{
        Image: "golang:latest",
        WorkingDir: "/workspace",
        Cmd: []string{"sh", "-c", "go mod init app\ngo mod tidy\ngo build main.go\ndate > start.txt\n./main < input.in > output.txt\ndate > end.txt"},
        Tty:   true,
        OpenStdin: true,
        StdinOnce: true,
        AttachStdin: true,
        AttachStdout: true,
        AttachStderr: true,
    }, &container.HostConfig{
        AutoRemove: false,
        Resources: container.Resources{
            Memory: int64(memLimitMB) * 1024 * 1024,
        },
        Binds: []string{fmt.Sprintf("%s:/workspace", tmpDir)},
    }, &network.NetworkingConfig{}, &ocispec.Platform{}, "C1")
    if err != nil { 
		c.SendResult(id, StatusCompileError)
        return err
    }
    defer func() {
        if err := cli.ContainerRemove(context.Background(), resp.ID, container.RemoveOptions{
            Force: true,
        }); err != nil {
            panic(err)
        }
    }()

    if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
        c.SendResult(id, StatusCompileError)
		return err
    }




    for {
        cn, err := cli.ContainerInspect(context.Background(), "C1")
        if err != nil || cn.State.Status != "running" {
            if cn.State.OOMKilled {
				c.SendResult(id, StatusMemoryLimitExceeded)
                return fmt.Errorf("memory limit exceeded")
            }
            break
        }
        time.Sleep(50 * time.Millisecond)
    }


    outputPath := filepath.Join(tmpDir, "output.txt")
    startPath := filepath.Join(tmpDir, "start.txt")
    endPath := filepath.Join(tmpDir, "end.txt")
    data, err := ioutil.ReadFile(outputPath)
    if err != nil {
        panic(err)
    }
    startData, err := ioutil.ReadFile(startPath)
    if err != nil {
        panic(err)
    }
    endData, err := ioutil.ReadFile(endPath)
    if err != nil {
        panic(err)
    }
    outputStr := string(data)

    st := strings.Split(string(startData), " ")[3]
    et := strings.Split(string(endData), " ")[3]
    startTime, _ := time.Parse("15:04:05", st)
    endTime, _ := time.Parse("15:04:05", et)
    
    duration := endTime.Sub(startTime)
    if outputStr == "" {
		c.SendResult(id, StatusRuntimeError)
        return fmt.Errorf("runtime error")
    }



    if duration > timeLimit {
		c.SendResult(id, StatusTimeLimitExceeded)
        return fmt.Errorf("time limit exceeded")
    }

	fmt.Println("Output: ", outputStr)

    if outputStr == wantOutput {
		c.SendResult(id, StatusOK)
        return nil
    }
    if outputStr != wantOutput {
		c.SendResult(id, StatusWrongOutput)
        return fmt.Errorf("wrong output: got %q, want %q", outputStr, wantOutput)
    }
    
    
    return nil
}
