package handler

import (
	"goCrud/model"
	"goCrud/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var req model.User

type CrudHandler struct {
	crudUsecase usecase.CrudUsecase
}

func NewCrudHandler(crudUsecase usecase.CrudUsecase) CrudHandler {
	return CrudHandler{crudUsecase: crudUsecase}
}

func (h *CrudHandler) GetAll(c *gin.Context) {
	users, err := h.crudUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *CrudHandler) CreateUser(c *gin.Context) {
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.crudUsecase.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
}

func (h *CrudHandler) UpdateUser(c *gin.Context) {
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.crudUsecase.UpdateUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
}

func (h *CrudHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		panic(err)
	}

	err = h.crudUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user delete success!"})
}
