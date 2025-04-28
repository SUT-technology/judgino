package coderunner

import (
	"github.com/SUT-technology/judgino/coderunner/dto"
	"github.com/SUT-technology/judgino/coderunner/services"
)

func Start(srvc services.RunnerServices){
	// srvc.StartProcessing()
	srvc.RunCode(dto.SubmissionRun{
		ID: 1,
		Code: `
		package main

import "fmt"

func main() {
  var name string
  _, err := fmt.Scanln(&name)
  if err != nil {
    fmt.Println("Error reading input:", err)
    return
  }
  fmt.Print("1")
}
		`,
		Input: "1",
		ExpectedOutput: "1",
		MemoryLimit: 1024,
		TimeLimit: 4,
	})
}