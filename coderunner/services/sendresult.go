package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (c RunnerServices) SendResult(submissionID, result int) {
	payload := map[string]int{"id": submissionID, "status": result}
	body, _ := json.Marshal(payload)

	_, err := http.Post(c.cfg.ApiUrl+"/result", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error sending result: %v", err)
	}

	log.Printf("send result submission: %v with status: %v", submissionID, result)
}
