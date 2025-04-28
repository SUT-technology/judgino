package coderunner

import (
	"github.com/SUT-technology/judgino/coderunner/services"
)

func Start(srvc services.RunnerServices){
	srvc.StartProcessing()
}