package repository

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type SubmissionRepository interface {
	GetSubmissionById(ctx context.Context, id uint)(*entity.Submission , error)
	GetSubmissionsByFilter(ctx context.Context, userId uint, questionId uint, submissionFilter string, finalFilter bool, pageParam uint) ([]*entity.Submission, error)
}