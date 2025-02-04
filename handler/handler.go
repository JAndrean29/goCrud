package handler

import (
	"goCrud/model"
	"goCrud/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var req model.User

func CreateUserHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := usecase.CreateUserUsecase(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": createdUser})
}

func GetUserHandler(c *gin.Context) {
	users, err := usecase.GetAllUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUserHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	affectedUser, err := usecase.UpdateUserUsecase(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": affectedUser})
}

func DeleteUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		panic(err)
	}

	err = usecase.DeleteUserUsecase(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user delete success!"})
}
