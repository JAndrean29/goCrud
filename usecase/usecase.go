package usecase

import (
	"errors"
	"goCrud/model"
	"goCrud/repository"
	"strings"
)

type CrudUsecase interface {
	CreateUser(user *model.User) (*model.User, error)
	GetAll() (*[]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int64) error
}

type crudUsecase struct {
	crudRepo repository.CrudRepository
}

func NewCrudUseCase(crudRepo repository.CrudRepository) CrudUsecase {
	return &crudUsecase{crudRepo: crudRepo}
}

// CreateUserUsecase implements CrudUsecase.
func (c *crudUsecase) CreateUser(user *model.User) (*model.User, error) {
	if user.Name == "" || user.Age <= 0 || (strings.ToLower(user.Gender) != "male" || strings.ToLower(user.Gender) != "female") {
		return nil, errors.New("invalid input detected")
	}

	return c.crudRepo.CreateUser(user)
}

// DeleteUserUsecase implements CrudUsecase.
func (c *crudUsecase) DeleteUser(id int64) error {
	if id <= 0 {
		return errors.New("invalid ID input")
	}

	return c.crudRepo.DeleteUser(id)
}

// GetAllUseCase implements CrudUsecase.
func (c *crudUsecase) GetAll() (*[]model.User, error) {
	return c.crudRepo.GetAll()
}

// UpdateUserUsecase implements CrudUsecase.
func (c *crudUsecase) UpdateUser(user *model.User) (*model.User, error) {
	if user.Name == "" || user.Age <= 0 || (strings.ToLower(user.Gender) != "male" || strings.ToLower(user.Gender) != "female") {
		return nil, errors.New("invalid input detected")
	}

	return c.crudRepo.UpdateUser(user)
}
