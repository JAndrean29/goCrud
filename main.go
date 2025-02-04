package main

import (
	"goCrud/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	r.POST("/users/create", handler.CreateUserHandler)

	r.Run()
}
