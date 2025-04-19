package postgresql

import (
	"context"
	// "fmt"

	"github.com/SUT-technology/judgino/internal/domain/entity"
	"gorm.io/gorm"
)

type submissionsTable struct {
	db *gorm.DB
}

func newSubmissionsTable(db *gorm.DB) submissionsTable {
	return submissionsTable{db: db}
}

func (c submissionsTable) GetSubmissionById(ctx context.Context, id uint) (*entity.Submission, error) {
	var submission entity.Submission
	c.db.First(&submission, id)
	return &submission, nil
}

func (c submissionsTable) GetSubmissionsByFilter(ctx context.Context, userId uint, questionId uint, submissionFilter string, finalFilter bool, pageParam uint) ([]*entity.Submission, error) {
	var submissions []*entity.Submission
	var query *gorm.DB
	if submissionFilter == "mine" && userId != 0 {
		query = c.db.Where("question_id = ?", questionId).Where("user_id = ?", userId).Where("is_final = ?", finalFilter).Order("submit_time").Offset(10 * (int(pageParam) - 1) - 1).Limit(10)
	} else {
		query = c.db.Where("question_id = ?", questionId).Where("is_final = ?", finalFilter).Order("submit_time").Offset(10 * (int(pageParam) - 1) - 1).Limit(10)
	}
	query.Find(&submissions)
	return submissions, nil
}