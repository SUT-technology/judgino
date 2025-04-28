package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SUT-technology/judgino/coderunner/dto"
)

func (c RunnerServices) FetchSubmissions() (*dto.SubmissionRunResp, error) {
	req, err := http.NewRequest("GET", c.cfg.ApiUrl+"/get", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("RUNNER-API-KEY", c.cfg.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submissions dto.SubmissionRunResp
	json.NewDecoder(resp.Body).Decode(&submissions)

	return &submissions, nil
}

func (c RunnerServices) StartProcessing() {
	for {
		submissions, err := c.FetchSubmissions()
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
			go c.RunCode(submission)
		}

		time.Sleep(c.cfg.TimeInterval) // Poll every second
	}
}
