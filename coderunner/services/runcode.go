package services

import (
	"log"
	"time"

	"github.com/SUT-technology/judgino/coderunner/dto"
)

func (c RunnerServices) RunCode(submission dto.SubmissionRun) error {
	log.Printf("recieve submission for run code: %v , input: %v, output: %v, memoryLimit: %v, timeLimit: %v", submission.Code, submission.Input, submission.ExpectedOutput, submission.MemoryLimit, submission.TimeLimit)
	time.Sleep(5 * time.Second)
	c.SendResult(submission.ID, 3)
	return nil
}
