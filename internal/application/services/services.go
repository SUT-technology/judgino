package services

import (
	auth "github.com/SUT-technology/judgino/internal/application/services/authsrvc"
	prof "github.com/SUT-technology/judgino/internal/application/services/profilesrvc"
	"github.com/SUT-technology/judgino/internal/application/services/questionssrvc"
	"github.com/SUT-technology/judgino/internal/application/services/runnersrvc"
	"github.com/SUT-technology/judgino/internal/application/services/submissionssrvc"
	"github.com/SUT-technology/judgino/internal/application/services/usersrvc"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/config"
)

func New(db repository.Pool, cfg config.Config) service.Service {
	return service.Service{
		AuthSrvc:       auth.NewAuthSrvc(db, cfg.Server.SecretKey),
		PrflSrvc:       prof.NewPrflSrvc(db),
		QuestionsSrvc:  questionssrvc.NewQuestionsSrvc(db),
		UserSrvc:       usersrvc.NewUserSrvc(db),
		SubmissionSrvc: submissionssrvc.NewSubmissionSrvc(db),
		RunnerSrvc:     runnersrvc.NewRunnerSrvc(db),
	}
}
