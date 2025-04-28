package pool

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/SUT-technology/judgino/internal/infrastructure/database/postgresql"
	"github.com/SUT-technology/judgino/internal/interface/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Pool wraps the GORM database connection
type Pool struct {
	db *gorm.DB
}

// New opens a connection to a PostgreSQL database using GORM
func New(cfg config.DB) (*Pool, error) {
	// PostgreSQL DSN (update placeholders with your actual database details)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open(%s): %w", dsn, err)
	}

	// err = db.AutoMigrate(&entity.User{}, &entity.Question{}, &entity.Submission{})
	// if err != nil {
	// 	log.Fatalf("Error migrating schema: %v", err)
	// }
	// fmt.Println("Connected to PostgreSQL using GORM!")
	return &Pool{db: db}, nil
}

// Query starts a transaction and executes the given function
func (p *Pool) Query(ctx context.Context, f repository.QueryFunc) error {
	// Begin a transaction
	tx := p.db.Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	r := &repository.Repo{
		Querier: &postgresql.GormQuerier{DB: tx},
		Tables:  postgresql.New(tx),
	}

	// Execute the function with the transaction
	err := f(r)
	if err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			return fmt.Errorf("rollback: %w query: %w", rollbackErr, err)
		}
		return fmt.Errorf("query: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

// Close is a placeholder for GORM cleanup (not strictly needed)
func (p *Pool) Close() error {
	fmt.Println("Closing database connection.")
	return nil
}
