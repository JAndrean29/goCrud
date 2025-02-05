package main

import (
	"goCrud/di"
	"goCrud/infrastructure/sqlitedb"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//init DB
	db := sqlitedb.InitiateSqliteConnection()

	//Build the service
	crudService := di.BuildCrudService(db)

	//set the Handler
	crudHandler := crudService.CrudHandler

	//declare GIN for routing
	r := gin.Default()

	//routes
	r.GET("/users", crudHandler.GetAll)
	r.POST("/users/create", crudHandler.CreateUser)
	r.POST("/users/edit", crudHandler.UpdateUser)
	r.DELETE("/usres/delete/:id", crudHandler.DeleteUser)

	//RUN GIN SERVER
	r.Run()
}
