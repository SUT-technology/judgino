package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type QuestionsService interface {
	GetQuestions(ctx context.Context, questionsDto dto.QuestionRequest) ([]*entity.Question, error)
	GetQuestion(ctx context.Context, questionId uint) (*entity.Question, error)
	GetSubmissions(ctx context.Context, submissionDto dto.SubmissionRequest) ([]*entity.Submission, error)
}
