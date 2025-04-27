package command

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SUT-technology/judgino/coderunner"
	"github.com/SUT-technology/judgino/coderunner/config"
	"github.com/SUT-technology/judgino/coderunner/services"
)

type SubmissionRun struct {
	ID             int    `json:"id"`
	Code           string `json:"code"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	TimeLimit      int    `json:"timeLimit"`
	MemoryLimit    int    `json:"memoryLimit"`
}

type SubmissionRunResp struct {
	Submissions []SubmissionRun `json:"submissions"`
}

func fetchSubmissions() (*SubmissionRunResp, error) {
	resp, err := http.Get("http://localhost:8000/api/runner/get")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submissions SubmissionRunResp
	json.NewDecoder(resp.Body).Decode(&submissions)
	return &submissions, nil
}

func sendResult(submissionID, result int) {
	payload := map[string]int{"id": submissionID, "status": result}
	body, _ := json.Marshal(payload)

	_, err := http.Post("http://localhost:8000/api/runner/result", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error sending result: %v", err)
	}

	log.Printf("send result submission: %v with status: %v", submissionID, result)
}

func RunCode(submission SubmissionRun) error {
	time.Sleep(5 * time.Second)
	sendResult(submission.ID, 3)
	return nil
}
func startProcessing() {
	for {
		submissions, err := fetchSubmissions()
		if err != nil {
			log.Printf("Error fetching submissions: %v", err)
			time.Sleep(time.Second)
			continue
		}
		if submissions == nil {
			time.Sleep(time.Second)
			continue
		}
		new_submissions := *submissions

		for _, submission := range new_submissions.Submissions {
			go RunCode(submission)
		}

		time.Sleep(1 * time.Second) // Poll every second
	}
}

func Run() error {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/code-runner.yaml", "Configuration File")
	flag.Parse()
	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	srvc := services.NewRunnerService(c)
	coderunner.Start(srvc)

	return nil
}
