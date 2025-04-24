package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type QuestionsService interface {
	GetQuestions(ctx context.Context, questionsDto dto.QuestionSummeryRequest, userId uint) (dto.QuestionsSummeryResponse, error)
	GetQuestion(ctx context.Context, questionId uint) (dto.QuestionSummery, error)
	QuestionsCount(ctx context.Context, questionsDto dto.QuestionSummeryRequest, userId uint) (int, error)
	CreateQuestion(ctx context.Context, questionsDto dto.CreateQuestionRequest,currentUserId int64) (dto.CreateQuestionResponse, error)
}
