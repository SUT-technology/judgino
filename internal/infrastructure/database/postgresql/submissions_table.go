package postgresql

import (
	"context"
	"fmt"
	"time"

	// "fmt"

	"github.com/SUT-technology/judgino/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

		if finalFilter {
			query = c.db.Where("question_id = ?", questionId).Where("user_id = ?", userId).Where("is_final = ?", true).Order("submit_time").Offset(10*(int(pageParam)-1) - 1).Limit(10)
		} else {
			query = c.db.Where("question_id = ?", questionId).Where("user_id = ?", userId).Order("submit_time").Offset(10*(int(pageParam)-1) - 1).Limit(10)
		}
	} else {
		if finalFilter {
			query = c.db.Where("question_id = ?", questionId).Where("is_final = ?", true).Order("submit_time").Offset(10*(int(pageParam)-1) - 1).Limit(10)
		} else {
			query = c.db.Where("question_id = ?", questionId).Order("submit_time").Offset(10*(int(pageParam)-1) - 1).Limit(10)
		}
	}
	query.Find(&submissions)
	return submissions, nil
}
func (c submissionsTable) GetSubmissionsCount(ctx context.Context, userId uint, questionId uint, submissionFilter string, finalFilter bool) (int, error) {
	var count int64
	var query *gorm.DB
	if submissionFilter == "mine" && userId != 0 {
		query = c.db.Model(&entity.Submission{}).Where("question_id = ?", questionId).Where("user_id = ?", userId).Where("is_final = ?", finalFilter)
	} else {
		query = c.db.Model(&entity.Submission{}).Where("question_id = ?", questionId).Where("is_final = ?", finalFilter)
	}
	query.Count(&count)
	return int(count), nil
}

func (c submissionsTable) CreateSubmission(ctx context.Context, submission entity.Submission) error {
	if err := c.db.Create(&submission).Error; err != nil {
		return err
	}
	return nil
}

func (c submissionsTable) UpdateSubmission(ctx context.Context, submission *entity.Submission) error {
	if err := c.db.Save(&submission).Error; err != nil {
		return err
	}
	return nil
}

func (c submissionsTable) GetSubmissionsForRunner(ctx context.Context, limit int) ([]*entity.Submission, error) {
	var submissions []*entity.Submission

	now := time.Now()

	query := c.db.
		Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
		Where(
			c.db.Where("status = ?", 1).
				Or("status = 2 AND runner_started_at < ?", now.Add(-1*time.Minute)),
		).
		Order("id ASC").
		Limit(limit)

	if err := query.Find(&submissions).Error; err != nil {
		return nil, fmt.Errorf("querying submissions: %w", err)
	}

	if len(submissions) == 0 {
		return nil, nil
	}

	var toFailIDs []uint
	var toRunningIDs []uint

	for _, submission := range submissions {
		if submission.TryCount >= 3 {
			toFailIDs = append(toFailIDs, submission.ID)
		} else {
			toRunningIDs = append(toRunningIDs, submission.ID)
		}
	}

	if len(toFailIDs) > 0 {
		if err := c.db.Model(&entity.Submission{}).
			Where("id IN ?", toFailIDs).
			Updates(map[string]interface{}{
				"status": 9,
			}).Error; err != nil {
			return nil, fmt.Errorf("bulk update fail submissions: %w", err)
		}
	}

	if len(toRunningIDs) > 0 {
		if err := c.db.Model(&entity.Submission{}).
			Where("id IN ?", toRunningIDs).
			Updates(map[string]interface{}{
				"status":            2,
				"runner_started_at": now,
				"try_count":         gorm.Expr("try_count + 1"),
			}).Error; err != nil {
			return nil, fmt.Errorf("bulk update running submissions: %w", err)
		}
	}

	return submissions, nil
}
