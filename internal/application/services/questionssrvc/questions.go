package questionssrvc

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/repository"
)

type QuestionsSrvc struct {
	db repository.Pool
}

func NewQuestionsSrvc(db repository.Pool) QuestionsSrvc {
	return QuestionsSrvc{
		db: db,
	}
}

func (c QuestionsSrvc) GetQuestions(ctx context.Context, questionsDto dto.QuestionRequest, userId uint) (dto.QuestionsResponse, error) {
	var (
		questions []*entity.Question
		err       error
	)

	if questionsDto.QuestionValue == "" {
		questionsDto.QuestionValue = "all"
	}
	if questionsDto.SortValue == "" {
		questionsDto.SortValue = "deadline"
	}
	if questionsDto.PageParam == 0 {
		questionsDto.PageParam = 1
	}

	questionsCount, err := c.QuestionsCount(ctx, questionsDto, userId)
	if err != nil {
		return dto.QuestionsResponse{Error: err}, err
	}

	totalPages := questionsCount/10 + 1
	if questionsDto.PageAction == "next" && questionsDto.PageParam < (totalPages) {
		questionsDto.PageParam++
	}
	if questionsDto.PageAction == "prev" && questionsDto.PageParam > 1 {
		questionsDto.PageParam--
	}
	if questionsDto.PageParam > (totalPages) {
		questionsDto.PageParam = (totalPages)
	}

	queryFuncFindQuestions := func(r *repository.Repo) error {
		questions, err = r.Tables.Questions.GetQuestionByFilter(ctx, questionsDto.SearchFilter, questionsDto.QuestionValue, questionsDto.SortValue, int(questionsDto.PageParam), userId)
		if err != nil {
			return fmt.Errorf("failed to get questions: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindQuestions)
	if err != nil {
		return dto.QuestionsResponse{Error: err}, err
	}
	// Create the data to pass to the template
	questionsData := make([]dto.Question, len(questions))
	for i, question := range questions {
		questionsData[i] = dto.Question{
			Title:       question.Title,
			PublishDate: question.PublishDate.Format("2006-01-02 15:04:05"),
			Deadline:    question.Deadline.Format("2006-01-02 15:04:05"),
		}

	}

	totalPages = questionsCount/10 + 1

	resp := dto.QuestionsResponse{
		Questions:      questionsData,
		TotalPages:     totalPages,
		CurrentPage:    (questionsDto.PageParam),
		SearchFilter:   questionsDto.SearchFilter,
		QuestionFilter: questionsDto.QuestionValue,
		SortFilter:     questionsDto.SortValue,
		Error:          nil,
	}
	fmt.Println(resp.CurrentPage)

	return resp, nil
}

func (c QuestionsSrvc) GetQuestion(ctx context.Context, questionId uint) (dto.Question, error) {
	var (
		question *entity.Question
		err      error
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
		return dto.Question{}, err
	}
	return dto.Question{
		Title:       question.Title,
		PublishDate: question.PublishDate.Format("2006-01-02 15:04:05"),
		Deadline:    question.Deadline.Format("2006-01-02 15:04:05"),
	}, nil
}

func (c QuestionsSrvc) QuestionsCount(ctx context.Context, questionsDto dto.QuestionRequest, userId uint) (int, error) {

	var (
		count int
		err   error
	)

	queryFuncFindQuestions := func(r *repository.Repo) error {
		count, err = r.Tables.Questions.GetQuestionsCount(ctx, questionsDto.SearchFilter, questionsDto.QuestionValue, userId)
		if err != nil {
			return fmt.Errorf("failed to get questions count: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindQuestions)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c QuestionsSrvc) PublishQuestion(ctx context.Context, questionId uint) error {

	queryFunc := func(r *repository.Repo) error {
		err := r.Tables.Questions.PublishQuestion(ctx, questionId)
		if err != nil {
			return fmt.Errorf("failed to get questions count: %w", err)
		}
		return nil
	}
	err := c.db.Query(ctx, queryFunc)
	if err != nil {
		return err
	}

	return nil
}
