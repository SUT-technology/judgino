package postgresql

import (
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"gorm.io/gorm"
)

type GormQuerier struct {
	DB *gorm.DB
}

func (g *GormQuerier) Exec(query string, args ...interface{}) error {
	return g.DB.Exec(query, args...).Error
}

func (g *GormQuerier) Find(dest interface{}, query string, args ...interface{}) error {
	return g.DB.Raw(query, args...).Scan(dest).Error
}

func (g *GormQuerier) First(dest interface{}, query string, args ...interface{}) error {
	return g.DB.Raw(query, args...).First(dest).Error
}

func New(db *gorm.DB) repository.Tables {
	return repository.Tables{
		Users: newUsersTable(db),
		Questions: newQuestionsTable(db),
		Submissions: newSubmissionsTable(db),
	}
}
