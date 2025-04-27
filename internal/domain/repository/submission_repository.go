package repository

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type SubmissionRepository interface {
	GetSubmissionById(ctx context.Context, id uint)(*entity.Submission , error)
	GetSubmissionsByFilter(ctx context.Context, userId uint, questionId uint, submissionFilter string, finalFilter bool, pageParam uint) ([]*entity.Submission, error)
	GetSubmissionsCount(ctx context.Context, userId uint, questionId uint, submissionFilter string, finalFilter bool) (int, error)
	CreateSubmission(ctx context.Context,  submission entity.Submission) error
	UpdateSubmission(ctx context.Context, submission *entity.Submission) error
	GetUnjudgedSubmissions(ctx context.Context) ([]*entity.Submission, error)
}