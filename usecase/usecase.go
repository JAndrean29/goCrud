package usecase

import (
	"fmt"
	"goCrud/infrastructure/sqlitedb"
	"goCrud/model"
)

func CreateUserUsecase(user *model.User) (*model.User, error) {
	db := sqlitedb.InitiateSqliteConnection()
	query := `insert into user (name,age,gender) values (?,?,?)`
	result, err := db.Exec(query, user.Name, user.Age, user.Gender)
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
