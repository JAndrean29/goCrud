package usecase

import (
	"context"
	"errors"
	"goCrud/config"
	"goCrud/repository"
	"strings"
)

type CrudUsecase interface {
	CreateUser(ctx context.Context, user repository.CreateUserParams) (repository.User, error)
	GetAll(ctx context.Context) ([]repository.User, error)
	UpdateUser(ctx context.Context, user repository.UpdateUserParams) (repository.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type crudUsecase struct {
	config   *config.Cfg
	crudRepo *repository.Queries
}

// CreateUser implements CrudUsecase.
func (c crudUsecase) CreateUser(ctx context.Context, user repository.CreateUserParams) (repository.User, error) {
	if user.Name.String == "" || user.Age.Int32 <= 0 || (strings.ToLower(user.Gender.String) != "male" && strings.ToLower(user.Gender.String) != "female") {
		return repository.User{}, errors.New("invalid input detected")
	}

	return c.crudRepo.CreateUser(ctx, user)
}

// DeleteUser implements CrudUsecase.
func (c crudUsecase) DeleteUser(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid ID input")
	}

	return c.crudRepo.DeleteUser(ctx, id)
}

// GetAll implements CrudUsecase.
func (c crudUsecase) GetAll(ctx context.Context) ([]repository.User, error) {
	return c.crudRepo.GetAll(ctx)
}

// UpdateUser implements CrudUsecase.
func (c crudUsecase) UpdateUser(ctx context.Context, user repository.UpdateUserParams) (repository.User, error) {
	if user.Name.String == "" || user.Age.Int32 <= 0 || (strings.ToLower(user.Gender.String) != "male" && strings.ToLower(user.Gender.String) != "female") {
		return repository.User{}, errors.New("invalid input detected")
	}

	return c.crudRepo.UpdateUser(ctx, user)
}

func NewCrudUseCase(config *config.Cfg, crudRepo *repository.Queries) CrudUsecase {
	return &crudUsecase{
		config:   config,
		crudRepo: crudRepo,
	}
}

/* OLD colde before integration
// CreateUserUsecase implements CrudUsecase.
func (c *crudUsecase) CreateUser(user *model.User) (*model.User, error) {
	if user.Name == "" || user.Age <= 0 || (strings.ToLower(user.Gender) != "male" && strings.ToLower(user.Gender) != "female") {
		return nil, errors.New("invalid input detected")
	}

	return c.CreateUser(user)
}

// DeleteUserUsecase implements CrudUsecase.
func (c *crudUsecase) DeleteUser(id int32) error {
	if id <= 0 {
		return errors.New("invalid ID input")
	}

	return c.DeleteUser(id)
}

// GetAllUseCase implements CrudUsecase.
func (c *crudUsecase) GetAll() (*[]model.User, error) {
	return c.GetAll()
}

// UpdateUserUsecase implements CrudUsecase.
func (c *crudUsecase) UpdateUser(user *model.User) error {
	if user.Name == "" || user.Age <= 0 || (strings.ToLower(user.Gender) != "male" && strings.ToLower(user.Gender) != "female") {
		return errors.New("invalid input detected")
	}

	return c.UpdateUser(user)
}
*/
