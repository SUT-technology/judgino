package postgresql

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/entity"
	"gorm.io/gorm"
)

type questionsTable struct {
	db *gorm.DB
}

func newQuestionsTable(db *gorm.DB) questionsTable {
	return questionsTable{db: db}
}

func (c questionsTable) GetQuestionById(ctx context.Context, id uint) (*entity.Question, error) {
	var question entity.Question
	c.db.First(&question, id)
	return &question, nil
}

func (c questionsTable) GetQuestionByFilter(ctx context.Context, searchFilter string, questionFilter string, sortFilter string, pageParam int, userId uint) ([]*entity.Question, error) {
	var questions []*entity.Question
	var query *gorm.DB
	if questionFilter == "mine" && userId != 0 {
		query = c.db.Where("title ILIKE ?", "%"+searchFilter+"%").Where("user_id = ?", userId).Order(sortFilter).Offset(10*(pageParam-1) - 1).Limit(10)
	} else {
		query = c.db.Where("title ILIKE ?", "%"+searchFilter+"%").Order(sortFilter).Offset(10*(pageParam-1) - 1).Limit(10)
	}
	query.Find(&questions)
	fmt.Println(questions)
	return questions, nil
}

func (c questionsTable) GetQuestionsCount(ctx context.Context, searchFilter string, questionFilter string, userId uint) (int, error) {
	var count int64
	var query *gorm.DB
	if questionFilter == "mine" && userId != 0 {
		query = c.db.Where("title ILIKE ?", "%"+searchFilter+"%").Where("user_id = ?", userId)
	} else {
		query = c.db.Where("title ILIKE ?", "%"+searchFilter+"%")
	}
	query.Model(&entity.Question{}).Count(&count)
	return int(count), nil
}


func (c questionsTable) CreateQuestion(ctx context.Context,question *entity.Question)error {
	if  err:=c.db.Create(&question).Error; err!=nil {
		return err
	}
	return nil
}

func (c questionsTable) UpdateQuestion(ctx context.Context, questionId int64,updateData *entity.Question)error {
	var question entity.Question
	c.db.First(&question, questionId)

	if err := c.db.WithContext(ctx).Model(&question).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (c questionsTable) PublishQuestion(ctx context.Context, question *entity.Question,updateData *entity.Question) error {
	fmt.Printf("update data: %v",*updateData)
	if err := c.db.WithContext(ctx).Model(question).Updates(updateData).Error; err != nil {
		return err
	}

	return nil

}

