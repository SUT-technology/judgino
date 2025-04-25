package submissionssrvc

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/repository"
)

type SubmissionService struct {
	db repository.Pool
}

func NewSubmissionSrvc(db repository.Pool) SubmissionService {
	return SubmissionService{
		db: db,
	}
}

func (c SubmissionService) GetSubmissions(ctx context.Context, submissionDto dto.SubmissionRequest, userId uint, isAdmin bool, questionId int) (dto.SubmissionsResponse, error) {
	var (
		submissions []*entity.Submission
		err         error
	)

	if submissionDto.SubmissionValue == "" {
		submissionDto.SubmissionValue = "all"
	}
	if submissionDto.FinalValue == "" {
		submissionDto.FinalValue = "all"
	}
	if submissionDto.PageParam == 0 {
		submissionDto.PageParam = 1
	}

	if submissionDto.SubmissionValue == "all" && !isAdmin {
		//Todo handle erro
		return dto.SubmissionsResponse{Error: err}, fmt.Errorf("user is not admin and submission value is all")
	}

	submissionsCount, err := c.SubmissionsCount(ctx, submissionDto, userId, questionId)
	if err != nil {
		return dto.SubmissionsResponse{Error: err}, err
	}
	totalPages := submissionsCount/10 + 1

	if submissionDto.PageAction == "next" && submissionDto.PageParam < uint(totalPages) {
		submissionDto.PageParam++
	}
	if submissionDto.PageAction == "prev" && submissionDto.PageParam > 1 {
		submissionDto.PageParam--
	}

	if submissionDto.PageParam > uint(totalPages) {
		submissionDto.PageParam = uint(totalPages)
	}

	queryFuncFindSubmissions := func(r *repository.Repo) error {
		submissions, err = r.Tables.Submissions.GetSubmissionsByFilter(ctx, userId, uint(questionId), submissionDto.SubmissionValue, submissionDto.FinalValue == "final", submissionDto.PageParam)
		if err != nil {
			return fmt.Errorf("failed to get submissions: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindSubmissions)
	if err != nil {
		// Todo fix error
		return dto.SubmissionsResponse{Error: err}, err
	}

	// Create the data to pass to the template
	submissionsData := make([]dto.Submission, len(submissions))
	for i, submission := range submissions {
		var qt string
		err := c.db.Query(ctx, func(r *repository.Repo) error {
			question, err := r.Tables.Questions.GetQuestionById(ctx, submission.QuestionID)
			if err != nil {
				return fmt.Errorf("failed to get question by id: %w", err)
			}
			qt = question.Title
			return nil
		})
		if err != nil {
			return dto.SubmissionsResponse{}, err
		}
		var un string
		err = c.db.Query(ctx, func(r *repository.Repo) error {
			user, err := r.Tables.Users.GetUserById(ctx, int64(submission.UserID))
			if err != nil {
				return fmt.Errorf("failed to get user by id: %w", err)
			}
			un = user.FirstName
			return nil
		})
		if err != nil {
			return dto.SubmissionsResponse{}, err
		}

		var typ string
		if submission.IsFinal {
			typ = "final"
		} else {
			typ = "normal"
		}
		submissionsData[i] = dto.Submission{
			QuestionTitle: qt,
			UserName:      un,
			Status:        submission.Status,
			Date:          submission.SubmitTime.Format("2006-01-02 15:04:05"),
			Type:          typ,
		}
	}

	if err != nil {
		return dto.SubmissionsResponse{Error: err}, err
	}
	totalPages = submissionsCount/10 + 1

	resp := dto.SubmissionsResponse{
		Submissions:      submissionsData,
		TotalPages:       totalPages,
		QuestionId:       questionId,
		CurrentPage:      int(submissionDto.PageParam),
		SubmissionFilter: submissionDto.SubmissionValue,
		FinalFilter:      submissionDto.FinalValue,
		Error:            nil,
	}
	return resp, nil
}

func (c SubmissionService) SubmissionsCount(ctx context.Context, submissionDto dto.SubmissionRequest, userId uint, questionId int) (int, error) {
	var (
		submissionsCount int
		err              error
	)

	queryFuncFindSubmissionsCount := func(r *repository.Repo) error {
		submissionsCount, err = r.Tables.Submissions.GetSubmissionsCount(ctx, userId, uint(questionId), submissionDto.SubmissionValue, submissionDto.FinalValue == "final")
		if err != nil {
			return fmt.Errorf("failed to get submissions count: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindSubmissionsCount)
	if err != nil {
		return 0, err
	}

	return submissionsCount, nil
}

func (c SubmissionService) SubmitQuestion(ctx context.Context, submitDto dto.SubmitRequest, userId int64, questionId int) error {

	submission := entity.Submission{
		SubmitURL:  submitDto.SubmitUrl,
		IsFinal:    false,
		QuestionID: uint(questionId),
		UserID:     uint(userId),
		Status:     2,
		SubmitTime: time.Now(),
	}
	queryFuncFindUser := func(r *repository.Repo) error {
		err := r.Tables.Submissions.CreateSubmission(ctx, submission)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}
		return nil
	}
	err := c.db.Query(ctx, queryFuncFindUser)
	if err != nil {
		return err
	}

	return nil
}

func (c SubmissionService) SendSubmissions(ctx context.Context) (dto.SubmissionRunResp, error) {
	var (
		submissions []*entity.Submission
		err error
	)
	queryFuncFindSubmissions := func(r *repository.Repo) error {
		submissions, err = r.Tables.Submissions.GetUnjudgedSubmissions(ctx)
		if err != nil {
			return fmt.Errorf("failed to get submissions: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindSubmissions)
	if err != nil {
		// Todo fix error
		return dto.SubmissionRunResp{}, err
	}
	var submissionData = make([]dto.SubmissionRun, len(submissions))
	for i, submission := range submissions {
		input, _ := ioutil.ReadFile(submission.Question.InputURL)
		output, _ := ioutil.ReadFile(submission.Question.OutputURL)
		code, _ := ioutil.ReadFile(submission.SubmitURL)
		inputString := string(input)
		outputString := string(output)
		codeString := string(code)
		submissionData[i] = dto.SubmissionRun{
			ID: submission.ID,
			Code: codeString,
			Input: inputString,
			ExpectedOutput: outputString,
			TimeLimit: int(submission.Question.TimeLimit),
			MemoryLimit: int(submission.Question.MemoryLimit),
		}
	}
	return dto.SubmissionRunResp{
		Submissions: submissionData,
	}, nil

}

func (c SubmissionService) SubmitResult(ctx context.Context, result dto.SubmissionResult) error {
	queryFuncFindSubmissions := func(r *repository.Repo) error {
		submission, err := r.Tables.Submissions.GetSubmissionById(ctx, result.ID)
		if err != nil {
			return fmt.Errorf("failed to get submission: %w", err)
		}
		submission.Status = int64(result.Status)
		fmt.Println(submission)
		err = r.Tables.Submissions.UpdateSubmission(ctx, submission)
		if err != nil {
			return fmt.Errorf("failed to update submission: %w", err)
		}
		return nil
	}

	err := c.db.Query(ctx, queryFuncFindSubmissions)
	if err != nil {
		return err
	}

	return nil
}