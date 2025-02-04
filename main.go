package main

import (
	"goCrud/infrastructure/sqlitedb"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	sqlitedb.InitiateSqliteConnection()

	r.POST("/users/create", CreateUserHandler)

	r.Run()
}
