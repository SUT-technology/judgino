package questionssrvc

import (
	"context"
	"fmt"
	"time"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/model"
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

func (c QuestionsSrvc) CreateQuestion(ctx context.Context, createQuestionDto dto.CreateQuestionRequest, currentUserId int64) (dto.CreateQuestionResponse, error) {
	var (
		response dto.CreateQuestionResponse
		question *entity.Question
		err      error
	)

	if createQuestionDto.Title == "" {
		response.Title = true
		response.Error = true
	}
	if createQuestionDto.Body == "" {
		response.Body = true
		response.Error = true
	}
	if createQuestionDto.TimeLimit == 0{
		response.TimeLimit = true
		response.Error = true
	}
	if createQuestionDto.MemoryLimit == 0 {
		response.MemoryLimit = true
		response.Error = true
	}
	if createQuestionDto.InputURL == "" {
		response.InputURL = true
		response.Error = true
	}
	if createQuestionDto.OutputURL == "" {
		response.OutputURL = true
		response.Error = true
	}

	if response.Error {
		response.Status = model.UserMessage("error creating question")
		return response, nil
	}
	time,err := time.Parse("2006-01-02T15:04:05.000Z",createQuestionDto.Deadline)
	if err!= nil {
		fmt.Errorf(err.Error())
	}

	createQuestionDto.UserID = currentUserId
	question = &entity.Question{
		UserID:      uint(createQuestionDto.UserID),
		Status:      "draft",
		Title:       createQuestionDto.Title,
		Body:        createQuestionDto.Body,
		TimeLimit:   createQuestionDto.TimeLimit,
		MemoryLimit: createQuestionDto.MemoryLimit,
		InputURL:    createQuestionDto.InputURL,
		OutputURL:   createQuestionDto.OutputURL,
		Deadline:    time,
	}

	queryFuncCreateQuestion := func(r *repository.Repo) error {
		err = r.Tables.Questions.CreateQuestion(ctx, question)
		if err != nil {
			return err
		}
		user, err := r.Tables.Users.GetUserById(ctx, currentUserId)
		if err != nil {
			return err
		}
		count := user.CreatedQuestionsCount + 1
		err = r.Tables.Users.FindAndUpdateUser(ctx, currentUserId, entity.User{CreatedQuestionsCount: count})
		return err
	}

	err = c.db.Query(ctx, queryFuncCreateQuestion)
	if err != nil {
		return dto.CreateQuestionResponse{}, err
	}

	return dto.CreateQuestionResponse{
		Status:     model.UserMessage("question created successfully"),
		UserID:     currentUserId,
		QuestionID: int64(question.ID),
	}, nil
}


func (c QuestionsSrvc) GetQuestions(ctx context.Context, questionsDto dto.QuestionSummeryRequest, userId uint) (dto.QuestionsSummeryResponse, error) {
	var (
		questions []*entity.Question
		currentUser *entity.User
		err       error
	)

	if questionsDto.QuestionValue == "" {
		questionsDto.QuestionValue = "all"
	}
	if questionsDto.SortValue == "" {
		questionsDto.SortValue = "publish_date"
	}
	if questionsDto.PageParam == 0 {
		questionsDto.PageParam = 1
	}

	questionsCount, err := c.QuestionsCount(ctx, questionsDto, userId)
	if err != nil {
		return dto.QuestionsSummeryResponse{Error: err}, err
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
		currentUser, err = r.Tables.Users.GetUserById(ctx,int64(userId))
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindQuestions)
	if err != nil {
		return dto.QuestionsSummeryResponse{Error: err}, err
	}
	// Create the data to pass to the template
	questionsData := make([]dto.QuestionSummery, len(questions))
	for i, question := range questions {
		
		questionsData[i] = dto.QuestionSummery{
			Title:         question.Title,
			PublishDate: question.PublishDate.Format("2006-01-02 15:04:05"),
			Deadline:    question.Deadline.Format("2006-01-02 15:04:05"),
			QuestionId: int64(question.ID),
			Status: question.Status,
			IsCurrentUserAdmin: currentUser.IsAdmin(),
		}
		
		if currentUser.IsAdmin() {
			
			var publisher *entity.User
			
			findPublisherQueryfunc:= func(r *repository.Repo) error {
				publisher,err = r.Users.GetUserById(ctx,int64(question.UserID))
				if err!= nil {
					fmt.Errorf("failed to find the publisher: %w",err)
				}
				return nil
			}

			err = c.db.Query(ctx, findPublisherQueryfunc)
			if err != nil {
				return dto.QuestionsSummeryResponse{Error: err}, err
			}

			questionsData[i].Publisher = publisher.Username
			questionsData[i].PublisherId = int64(publisher.ID)

		}

	}

	totalPages = questionsCount/10 + 1

	resp := dto.QuestionsSummeryResponse{
		Questions:  questionsData,
		TotalPages: totalPages,
		CurrentPage: (questionsDto.PageParam),
		SearchFilter: questionsDto.SearchFilter,
		QuestionFilter: questionsDto.QuestionValue,
		SortFilter:     questionsDto.SortValue,
		IsCurrentUserAdmin: currentUser.IsAdmin(),
		Error:          nil,
	}
	fmt.Println(resp.CurrentPage)

	return resp, nil
}

func (c QuestionsSrvc) GetQuestion(ctx context.Context, questionId uint) (*entity.Question, error) {
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
		return nil, err
	}
	return question, nil
}

func (c QuestionsSrvc) QuestionsCount(ctx context.Context, questionsDto dto.QuestionSummeryRequest, userId uint) (int, error) {

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
