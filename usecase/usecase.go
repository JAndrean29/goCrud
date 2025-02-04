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

func ShowAll() (*[]model.User, error) {
	var users []model.User
	err := db.Select(&users, `select * from user`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users data: %w", err)
	}

	return &users, nil
}
