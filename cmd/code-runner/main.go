package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type SubmissionRun struct {
	ID 		   uint   `json:"id"`
	Code           string `json:"code"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	TimeLimit      int    `json:"timeLimit"`
	MemoryLimit    int    `json:"memoryLimit"`
}

type SubmissionRunResp struct {
	Submissions []SubmissionRun `json:"submissions"`
}

type SubmissionResult struct {
	ID 		   uint   `json:"id"`
	Status         int `json:"status"`
}

// RunGoCode runs the provided Go code and checks the output against the expected result
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

	// Create a Dockerfile in the temporary directory to run the Go code
	dockerfile := tempDir + "/Dockerfile"
	dockerfileContent := fmt.Sprintf(`
	FROM golang:1.23
	WORKDIR /app
	COPY . .
	RUN go mod init code_runner
	RUN go mod tidy
	EXPOSE 8080
	CMD ["go", "run", "main.go"]
	`)
	err = os.WriteFile(dockerfile, []byte(dockerfileContent), 0644)
	if err != nil {
		return "Compile error"
	}
	fmt.Println("Dockerfile created")

	// Build the Docker image
	cmd := exec.Command("docker", "build", "-t", "go_code_runner", tempDir)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println()
		return "Compile error"
	}
	fmt.Println("Docker image built")

	// Run the Docker container with the provided memory and time limits
	runCmd := exec.Command(
		"docker", "run", 
		"--memory", fmt.Sprintf("%d", memoryLimit), // Set the memory limit
		"--rm", 
		"-i", 
		"--name", "go_code_container", 
		"go_code_runner", 
		"--", "main.go",
	)
	runCmd.Stdin = strings.NewReader(input)
	var runOut, runErr bytes.Buffer
	runCmd.Stdout = &runOut
	runCmd.Stderr = &runErr

	// Set a timeout for the execution
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	// Run the Go code inside the container
	done := make(chan error)
	go func() {
		done <- runCmd.Run()
	}()

	select {
	case <-done:
		// Check if the output matches the expected output
		if runOut.String() == expectedOutput {
			return "OK"
		} else {
			return "Wrong output"
		}
	case <-timer.C:
		// Kill the container if time exceeds
		cmd := exec.Command("docker", "kill", "go_code_container")
		cmd.Run()
		return "Time limit exceeded"
	}
	// select {
    // case err := <-done:
    //     if err != nil {
    //         fmt.Printf("Error running Docker container: %v\n", err)
    //         fmt.Printf("Container stderr: %s\n", runErr.String())
    //         return "Runtime error"
    //     }
    //     fmt.Printf("Container stdout: %s\n", runOut.String())
    //     fmt.Printf("Container stderr: %s\n", runErr.String())
    //     if strings.TrimSpace(runOut.String()) == expectedOutput {
    //         return "OK"
    //     }
    //     return "Wrong output"
    // case <-timer.C:
    //     exec.Command("docker", "kill", "go_code_container").Run()
    //     return "Time limit exceeded"
    // }
}

// SendResultToServer sends the result of the code execution to the main server
func SendResultToServer(result SubmissionResult) {
	// The server URL to send the result
	serverURL := "http://localhost:8000/runner/result" // Change this to your server's actual URL

	// Convert the result to JSON
	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error marshaling result: %v\n", err)
		return
	}

	// Send a POST request to the server
	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(resultJSON))
	if err != nil {
		fmt.Printf("Error sending result: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Print the server's response (optional)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading server response: %v\n", err)
		return
	}
	fmt.Printf("Server response: %s\n", body)
}

func PollForCode() {
	for {
		// Make a request to the main program's /get-code endpoint
		resp, err := http.Get("http://localhost:8000/runner/get")
		if err != nil {
			fmt.Printf("Error fetching code: %v\n", err)
			time.Sleep(2 * time.Minute) // Wait before retrying
			continue
		}

		// Decode the response into CodeRequest struct
		var srr SubmissionRunResp
		if err := json.NewDecoder(resp.Body).Decode(&srr); err != nil {
			fmt.Printf("Error decoding response: %v\n", err)
			time.Sleep(2 * time.Minute) // Wait before retrying
			continue
		}

		// Run the Go code
		for _, codeRequest := range srr.Submissions {
			fmt.Printf("Running code: %s\n", codeRequest.Code)
			fmt.Printf("Input: %s\n", codeRequest.Input)
			fmt.Printf("Expected Output: %s\n", codeRequest.ExpectedOutput)
			fmt.Printf("Time Limit: %d\n", codeRequest.TimeLimit)
			go func(codeRequest SubmissionRun) {
				// Run the code with the provided input and expected output
				// result := RunGoCode(codeRequest.Code, codeRequest.Input, codeRequest.ExpectedOutput, codeRequest.TimeLimit, codeRequest.MemoryLimit)
				result := "OK"
				var status int
				switch result {
				case "OK":
					status = 3
				case "Wrong output":
					status = 4
				case "Time limit exceeded":
					status = 5
				case "Compile error":
					status = 6
				case "Memory limit exceeded":
					status = 7
				case "Runtime error":
					status = 8
				}

				// Send the result to the server
				submissionResult := SubmissionResult{
					ID: 		  codeRequest.ID,
					Status:         status,
				}
				SendResultToServer(submissionResult)
			}(codeRequest)
		}

		// Wait for the next cycle (2 minutes)
		time.Sleep(2 * time.Minute)
	}
}

func main() {
	// Start the polling loop
	PollForCode()
	// fmt.Println(RunGoCode(`package main
	// import "fmt"
	// func main() {
	// 	time.Sleep(20 * time.Second)
	// 	fmt.Println("Hello, World!")
	// 	}`, "", "Hello, World!", 2, 128))
}