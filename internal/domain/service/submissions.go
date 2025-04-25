package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type SubmissionService interface {
	GetSubmissions(ctx context.Context, submissionDto dto.SubmissionRequest, userId uint, isAdmin bool, questionId int) (dto.SubmissionsResponse, error)
	SubmissionsCount(ctx context.Context, submissionDto dto.SubmissionRequest, userId uint, questionId int) (int, error)
	SubmitQuestion(ctx context.Context, submitDto dto.SubmitRequest, userId int64, questionId int) error
	SendSubmissions(ctx context.Context) (dto.SubmissionRunResp, error)
	SubmitResult(ctx context.Context, result dto.SubmissionResult) error
}
