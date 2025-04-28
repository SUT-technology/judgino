package runnersrvc

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/repository"
)

type RunnerService struct {
	db repository.Pool
}

func NewRunnerSrvc(db repository.Pool) RunnerService {
	return RunnerService{
		db: db,
	}
}



func (c RunnerService) SendSubmissions(ctx context.Context) (dto.SubmissionRunResp, error) {
	var (
		submissions []*entity.Submission
		err         error
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
		var qt *entity.Question
		err = c.db.Query(ctx, func(r *repository.Repo) error {
			qt, err = r.Tables.Questions.GetQuestionById(ctx, submission.QuestionID)
			if err != nil {
				return fmt.Errorf("failed to get question by id: %w", err)
			}
			return nil
		})
		if err != nil {
			return dto.SubmissionRunResp{}, err
		}
		input, _ := ioutil.ReadFile(qt.InputURL)
		output, _ := ioutil.ReadFile(qt.OutputURL)
		code, _ := ioutil.ReadFile(submission.SubmitURL)
		inputString := string(input)
		outputString := string(output)
		codeString := string(code)
		submissionData[i] = dto.SubmissionRun{
			ID:             submission.ID,
			Code:           codeString,
			Input:          inputString,
			ExpectedOutput: outputString,
			TimeLimit:      int(qt.TimeLimit),
			MemoryLimit:    int(qt.MemoryLimit),
		}
	}
	return dto.SubmissionRunResp{
		Submissions: submissionData,
	}, nil

}

func (c RunnerService) SubmitResult(ctx context.Context, result dto.SubmissionResult) error {
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
