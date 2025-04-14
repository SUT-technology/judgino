package repository

import (
	"context"
)

type QueryFunc = func(r *Repo) error

type Querier interface {
	Exec(query string, args ...interface{}) error
	Find(dest interface{}, query string, args ...interface{}) error
	First(dest interface{}, query string, args ...interface{}) error
}

type Pool interface {
	Query(ctx context.Context, f QueryFunc) error
	Close() error
}

type Tables struct {
	Users UserRepository
}

type Repo struct {
	Tables
	Querier
}
