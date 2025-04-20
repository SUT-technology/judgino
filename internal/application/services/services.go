package services

import (
	prof "github.com/SUT-technology/judgino/internal/application/services/profilesrvc"
	auth "github.com/SUT-technology/judgino/internal/application/services/authsrvc"
	"github.com/SUT-technology/judgino/internal/application/services/authsrvc"
	"github.com/SUT-technology/judgino/internal/application/services/questionssrvc"
	"github.com/SUT-technology/judgino/internal/application/services/submissionssrvc"
	"github.com/SUT-technology/judgino/internal/application/services/usersrvc"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/domain/service"
)

func New(db repository.Pool) service.Service {
	return service.Service{
		AuthSrvc: auth.NewAuthSrvc(db),
		PrflSrvc: prof.NewPrflSrvc(db),
		AuthSrvc: authsrvc.NewAuthSrvc(db),
		QuestionsSrvc: questionssrvc.NewQuestionsSrvc(db),
		UserSrvc: usersrvc.NewUserSrvc(db),
		SubmissionSrvc: submissionssrvc.NewSubmissionSrvc(db),
	}
}
