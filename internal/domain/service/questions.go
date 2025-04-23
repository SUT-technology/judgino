package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type QuestionsService interface {
	GetQuestions(ctx context.Context, questionsDto dto.QuestionRequest, userId uint) (dto.QuestionsResponse, error)
	GetQuestion(ctx context.Context, questionId uint) (*entity.Question, error)
	QuestionsCount(ctx context.Context, questionsDto dto.QuestionRequest, userId uint) (int, error)
	PublishQuestion(ctx context.Context, questionId uint) error
}
