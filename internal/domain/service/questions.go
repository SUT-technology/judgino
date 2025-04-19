package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type QuestionsService interface {
	GetQuestions(ctx context.Context, questionsDto dto.QuestionRequest) (dto.QuestionsResponse, error)
	GetQuestion(ctx context.Context, questionId uint) (dto.Question, error)
	QuestionsCount(ctx context.Context, questionsDto dto.QuestionRequest) (int, error)
}
