// This is the factory/provider/supplier, they rummages the data then gives it to the chef
package repository

import (
	"fmt"

	"goCrud/model"

	"github.com/jmoiron/sqlx"
)

type CrudRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetAll() (*[]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int64) error
}

type crudRepository struct {
	db *sqlx.DB
}

func NewCrudRepository(db *sqlx.DB) CrudRepository {
	return &crudRepository{db: db}
}

func (r *crudRepository) CreateUser(user *model.User) (*model.User, error) {
	result, err := r.db.Exec(`insert into user (name,age,gender) values (?,?,?)`, user.Name, user.Age, user.Gender)
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed 2: %w", err)
	}

	user.ID = id
	return user, nil
}

func (r *crudRepository) GetAll() (*[]model.User, error) {
	var users []model.User
	err := r.db.Select(&users, `select * from user`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users data: %w", err)
	}

	return &users, nil
}

func (r *crudRepository) UpdateUser(user *model.User) (*model.User, error) {
	result, err := r.db.Exec(`update user set name=?, age=?, gender=? where id=?`, user.Name, user.Age, user.Gender, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed update 1: %w", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed update 2: %w", err)
	}

	if row == 0 {
		return nil, fmt.Errorf("selected user ID does not exists: %d", user.ID)
	}

	return user, nil
}

func (r *crudRepository) DeleteUser(id int64) error {
	result, err := r.db.Exec(`delete from user where id=?`, id)
	if err != nil {
		return fmt.Errorf("failed delete 1: %w", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed delete 2: %w", err)
	}

	if row == 0 {
		return fmt.Errorf("selected ID does note exists: %d", id)
	}

	return nil
}
