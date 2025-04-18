package repository

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type QuestionRepository interface {
	GetQuestionById(ctx context.Context, id uint)(*entity.Question , error)
	GetQuestionByFilter(ctx context.Context, searchFilter string, questionFilter string, sortFilter string, pageParam int, userId uint) ([]*entity.Question, error)
}