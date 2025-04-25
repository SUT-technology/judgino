package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunGoCode(code string, input string, expectedOutput string, timeLimit int, memoryLimit int) string {
    // Create a temporary directory for the Go code
    tempDir, err := ioutil.TempDir("", "go_code")
    if err != nil {
        return "Compile error"
    }
    defer os.RemoveAll(tempDir)

    // Write the provided Go code into a temporary Go file
    tempFile := tempDir + "/main.go"
    err = os.WriteFile(tempFile, []byte(code), 0644)
    if err != nil {
        return "Compile error"
    }

    // Run the Go code using the `go run` command
    cmd := exec.Command("go", "run", tempFile)
    cmd.Stdin = strings.NewReader(input) // Pass the input to stdin

    // Capture the output and error
    var out, stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    // Set a timeout for the execution
    timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

    // Run the Go code
    done := make(chan error)
    go func() {
        done <- cmd.Run()
    }()

    select {
    case <-done:
        // Check if the output matches the expected output
        if out.String() == expectedOutput {
            return "OK"
        } else {
            return "Wrong output"
        }
    case <-timer.C:
        // Timeout, kill the process if time exceeds
        cmd.Process.Kill()
        return "Time limit exceeded"
    }
}


func PollForCode() {
}

func main() {
	// Start the polling loop
	// PollForCode()
	fmt.Println(RunGoCode(`package main
	import "fmt"
	func main() {
		time.Sleep(20 * time.Second)
		fmt.Println(2)
		}`, "", "2", 2, 128))
}