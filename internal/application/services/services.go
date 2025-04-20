package services

import (
	auth "github.com/SUT-technology/judgino/internal/application/services/authsrvc"
	prof "github.com/SUT-technology/judgino/internal/application/services/profilesrvc"
	"github.com/SUT-technology/judgino/internal/application/services/questionssrvc"
	"github.com/SUT-technology/judgino/internal/application/services/submissionssrvc"
	"github.com/SUT-technology/judgino/internal/application/services/usersrvc"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/domain/service"
)

func New(db repository.Pool) service.Service {
	return service.Service{
		AuthSrvc:       auth.NewAuthSrvc(db),
		PrflSrvc:       prof.NewPrflSrvc(db),
		QuestionsSrvc:  questionssrvc.NewQuestionsSrvc(db),
		UserSrvc:       usersrvc.NewUserSrvc(db),
		SubmissionSrvc: submissionssrvc.NewSubmissionSrvc(db),
	}
}
