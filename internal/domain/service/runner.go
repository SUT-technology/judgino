package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type RunnerService interface {
	SendSubmissions(ctx context.Context) (dto.SubmissionRunResp, error)
	SubmitResult(ctx context.Context, result dto.SubmissionResult) error
}
