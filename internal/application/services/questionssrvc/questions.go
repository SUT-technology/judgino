package questionssrvc

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type QuestionsSrvc struct {
	db repository.Pool
}

func NewQuestionsSrvc(db repository.Pool) QuestionsSrvc {
	return QuestionsSrvc{
		db: db,
	}
}

func (c QuestionsSrvc) GetDate(ctx context.Context, questionsDto dto.QuestionsDto) ([]*entity.Question, error) {
	var (
		questions []*entity.Question
		err  error
	)

	queryFuncFindQuestions := func(r *repository.Repo) error {
		questions, err = r.Tables.Questions.GetQuestionByFilter(ctx, questionsDto.SearchValue, questionsDto.QuestionValue, questionsDto.SortValue, int(questionsDto.PageParam), questionsDto.UserId)
		if err != nil {
			return fmt.Errorf("failed to get questions: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindQuestions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}
