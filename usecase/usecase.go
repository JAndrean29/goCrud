package usecase

import (
	"fmt"
	"goCrud/infrastructure/sqlitedb"
	"goCrud/model"
)

var db = sqlitedb.InitiateSqliteConnection()

func CreateUserUsecase(user *model.User) (*model.User, error) {
	result, err := db.Exec(`insert into user (name,age,gender) values (?,?,?)`, user.Name, user.Age, user.Gender)
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

func GetAllUseCase() (*[]model.User, error) {
	var users []model.User
	err := db.Select(&users, `select * from user`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users data: %w", err)
	}

	return &users, nil
}

func UpdateUserUsecase(user *model.User) (*model.User, error) {
	result, err := db.Exec(`update user set name=?, age=?, gender=? where id=?`, user.Name, user.Age, user.Gender, user.ID)
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

func DeleteUserUsecase(id int64) error {
	result, err := db.Exec(`delete from user where id=?`, id)
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
