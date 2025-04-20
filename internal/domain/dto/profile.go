package dto

import "github.com/SUT-technology/judgino/internal/domain/model"

type ProfileRequest struct {
	UserId int64 `json:"userId" validate:"required"`
}

type ProfileRespone struct {
	UserId    int64
	CurrentUserId int64
	Username string
	Phone string
	Email string
	Role string
	NotAccepted int64
	Accepted int64
	Total int64
	SolvedPercentage int
	IsCurrentUserAdmin bool
	err model.UserMessage
}