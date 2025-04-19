package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type SubmissionService interface {
	GetSubmissions(ctx context.Context, submissionDto dto.SubmissionRequest) (dto.SubmissionsResponse, error)
	SubmissionsCount(ctx context.Context, submissionDto dto.SubmissionRequest) (int, error)
}
