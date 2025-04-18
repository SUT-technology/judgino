package services

import (
	"github.com/SUT-technology/judgino/internal/application/services/authsrvc"
	profilesrvc "github.com/SUT-technology/judgino/internal/application/services/profilesrvc"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/domain/service"
)

func New(db repository.Pool) service.Service {
	return service.Service{
		AuthSrvc: authsrvc.NewAuthSrvc(db),
		PrflSrvc: profilesrvc.NewPrflSrvc(db),
	}
}
