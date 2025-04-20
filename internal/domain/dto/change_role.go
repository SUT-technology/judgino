package dto

import "github.com/SUT-technology/judgino/internal/domain/model"

type ChangeRoleRequest struct {
	ID                    uint   `json:"ID"`
	// FirstName             string `json:"FirstName"`
	// Email                 string `json:"Email"`
	// Phone                 string `json:"Phone"`
	// Username              string `json:"Username"`
	// Password              string `json:"Password"`
	Role                  string `json:"Role"`
	// CreatedQuestionsCount int64  `json:"Created"`
	// SolvedQuestionsCount  int64  `json:"Solved"`
	// SubmissionsCount      int64  `json:"Submissions"`
	err model.UserMessage
}

type ChangeRoleResponse struct {
	err model.UserMessage
}
