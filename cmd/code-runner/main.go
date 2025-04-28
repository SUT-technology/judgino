package main

import (
	"context"
	"fmt"
	"strings"

	"io/ioutil"
	"os"

	// "os/exec"
	"path/filepath"
	// "strings"
	"time"
	// "encoding/json"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	// "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	// "github.com/google/uuid"
	// "github.com/docker/docker/api/types"
)

const (
    StatusOK = iota
    StatusWrongOutput
    StatusCompileError
    StatusRuntimeError
    StatusTimeLimitExceeded
    StatusMemoryLimitExceeded
)
func RunCodeInContainer(code, input, wantOutput string, timeLimit time.Duration, memLimitMB int) (int, error) {
    wantOutput += "\n"

    tmpDir, err := ioutil.TempDir("", "go-run-")
    if err != nil {
        return StatusRuntimeError, err
    }
    defer os.RemoveAll(tmpDir)

    srcPath := filepath.Join(tmpDir, "main.go")
    if err := ioutil.WriteFile(srcPath, []byte(code), 0644); err != nil {
        return StatusRuntimeError, err
    }


    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
    
    resp, err := cli.ContainerCreate(context.Background(), &container.Config{
        Image: "golang:latest",
        WorkingDir: "/workspace",
        Cmd: []string{"sh", "-c", "go mod init app\ngo mod tidy\ngo build main.go\ndate > start.txt\n./main > output.txt\ndate > end.txt"},
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
        return StatusCompileError, err
    }
    defer func() {
        if err := cli.ContainerRemove(context.Background(), resp.ID, container.RemoveOptions{
            Force: true,
        }); err != nil {
            panic(err)
        }
    }()

    if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
        return StatusCompileError, err
    }




    for {
        cn, err := cli.ContainerInspect(context.Background(), "C1")
        if err != nil || cn.State.Status != "running" {
            if cn.State.OOMKilled {
                return StatusMemoryLimitExceeded, fmt.Errorf("memory limit exceeded")
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
        return StatusRuntimeError, fmt.Errorf("runtime error")
    }



    if duration > timeLimit {
        return StatusTimeLimitExceeded, fmt.Errorf("time limit exceeded")
    }


    if outputStr == wantOutput {
        return StatusOK, nil
    }
    if outputStr != wantOutput {
        return StatusWrongOutput, fmt.Errorf("wrong output: got %q, want %q", outputStr, wantOutput)
    }
    
    

    return StatusOK, nil
}

func main() {
	// Start the polling loop
	// PollForCode()
	fmt.Println(RunCodeInContainer(`
    package main
    import "fmt"
	func main() {
        fmt.Println("")
	}`, "", "", 2, 1024))
    
}