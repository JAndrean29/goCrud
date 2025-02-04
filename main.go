package main

import (
	"goCrud/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetUserHandler)
	r.POST("/users/create", handler.CreateUserHandler)
	r.POST("/users/edit", handler.UpdateUserHandler)

	r.Run()
}
