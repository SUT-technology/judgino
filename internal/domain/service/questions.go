package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type QuestionsService interface {
	GetDate(ctx context.Context, questionsDto dto.QuestionsDto) ([]*entity.Question, error)
}
