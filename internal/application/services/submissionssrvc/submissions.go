package submissionssrvc

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type SubmissionService struct {
	db repository.Pool
}

func NewSubmissionSrvc(db repository.Pool) SubmissionService {
	return SubmissionService{
		db: db,
	}
}


func (c SubmissionService) GetSubmissions(ctx context.Context, submissionDto dto.SubmissionRequest) (dto.SubmissionsResponse, error) {
	var (
		submissions []*entity.Submission
		err  error
	)
	if submissionDto.SubmissionValue == "all" && !submissionDto.IsAdmin {
		//Todo handle erro
		return dto.SubmissionsResponse{}, fmt.Errorf("user is not admin and submission value is all")
	}

	submissionsCount, err := c.SubmissionsCount(ctx, submissionDto)
	if err != nil {
		return dto.SubmissionsResponse{}, err
	}
	totalPages := submissionsCount / 10 + 1

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
		submissions, err = r.Tables.Submissions.GetSubmissionsByFilter(ctx, submissionDto.UserId, uint(submissionDto.QuestionId), submissionDto.SubmissionValue, submissionDto.FinalValue == "final", submissionDto.PageParam)
		if err != nil {
			return fmt.Errorf("failed to get submissions: %w", err)
		}
		return nil
	}

	err = c.db.Query(ctx, queryFuncFindSubmissions)
	if err != nil {
		// Todo fix error
		return dto.SubmissionsResponse{}, err
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
			user, err := r.Tables.Users.GetUserById(ctx, submission.UserID)
			if err != nil {
				return fmt.Errorf("failed to get user by id: %w", err)
			}
			un = user.FirstName
			return nil
		})
		if err != nil {
			return dto.SubmissionsResponse{}, err
		}
		fmt.Println(submission.UserID)

		var typ string
		if submission.IsFinal {
			typ = "final"
		} else {
			typ = "normal"
		}
		submissionsData[i] = dto.Submission{
			QuestionTitle: qt,
			UserName:     un,
			Status:       submission.Status,
			Date:         submission.SubmitTime.Format("2006-01-02 15:04:05"),
			Type:         typ,
		}
	}

	submissionsCount, err = c.SubmissionsCount(ctx, submissionDto)
	if err != nil {
		return dto.SubmissionsResponse{}, err
	}
	totalPages = submissionsCount / 10 + 1

	resp := dto.SubmissionsResponse{
		Submissions: submissionsData,
		TotalPages:  totalPages,
		QuestionId:  submissionDto.QuestionId,
		CurrentPage: int(submissionDto.PageParam),
		SubmissionFilter: submissionDto.SubmissionValue,
		FinalFilter: submissionDto.FinalValue,
	}
	return resp, nil
}

func (c SubmissionService) SubmissionsCount(ctx context.Context, submissionDto dto.SubmissionRequest) (int, error) {
	var (
		submissionsCount int
		err              error
	)

	queryFuncFindSubmissionsCount := func(r *repository.Repo) error {
		submissionsCount, err = r.Tables.Submissions.GetSubmissionsCount(ctx, submissionDto.UserId, uint(submissionDto.QuestionId), submissionDto.SubmissionValue, submissionDto.FinalValue == "final")
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
