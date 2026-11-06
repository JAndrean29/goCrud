package usecase

import (
	"context"
	"goCrud/repository"
)

type userRepo interface {
	CreateUser(ctx context.Context, req repository.CreateUserParams) (out repository.User, err error)
	DeleteUser(ctx context.Context, userID int64) (err error)
	UpdateUser(ctx context.Context, req repository.UpdateUserParams) (out repository.User, err error)
	GetAll(ctx context.Context) ([]repository.User, error)
}
