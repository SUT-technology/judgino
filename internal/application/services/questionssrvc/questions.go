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

func (c QuestionsSrvc) GetQuestions(ctx context.Context, questionsDto dto.QuestionsDto) ([]*entity.Question, error) {
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

func (c QuestionsSrvc) GetQuestion(ctx context.Context, questionId uint) (*entity.Question, error) {
	var (
		question *entity.Question
		err  error
	)

	queryFuncFindQuestion := func(r *repository.Repo) error {
		question, err = r.Tables.Questions.GetQuestionById(ctx, questionId)
		if err != nil {
			return fmt.Errorf("failed to get question: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindQuestion)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (c QuestionsSrvc) GetSubmissions(ctx context.Context, submissionDto dto.SubmissionsDto) ([]*entity.Submission, error) {
	var (
		submissions []*entity.Submission
		err  error
	)
	

	queryFuncFindSubmissions := func(r *repository.Repo) error {
		submissions, err = r.Tables.Submissions.GetSubmissionsByFilter(ctx, submissionDto.UserId, submissionDto.QuestionId, submissionDto.SubmissonValue, submissionDto.FinalValue, submissionDto.PageParam)
		if err != nil {
			return fmt.Errorf("failed to get submissions: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindSubmissions)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}
