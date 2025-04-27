package service

import (
	"context"
	"mime/multipart"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type SubmissionService interface {
	GetSubmissions(ctx context.Context, submissionDto dto.SubmissionRequest, userId uint, isAdmin bool, questionId int) (dto.SubmissionsResponse, error)
	SubmissionsCount(ctx context.Context, submissionDto dto.SubmissionRequest, userId uint, questionId int) (int, error)
	SubmitQuestion(ctx context.Context, file *multipart.FileHeader, userId int64, questionId int) error
}
