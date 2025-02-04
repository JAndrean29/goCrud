package usecase

import (
	"fmt"
	"goCrud/infrastructure/sqlitedb"
	"goCrud/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserUsecase(user *model.User) (*model.User, error) {
	query := `insert into user (name,age,gender) values (?,?,?)`
	result, err := sqlitedb.Exec(query, user.Name, user.Age, user.Gender)
	if err != nil {
		return nil, fmt.Errorf("failed:", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed 2:", err)
	}

	user.ID = id
	return user, nil
}

func createUserHandler(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(req)

	createdUser, err := CreateUserUsecase(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": createdUser})
}
