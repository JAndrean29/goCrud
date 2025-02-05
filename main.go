// This is the customer, asking for a ready to eat meal
package main

import (
	"goCrud/di"
	"goCrud/handler/router"
	"goCrud/infrastructure/sqlitedb"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//init DB
	db := sqlitedb.InitiateSqliteConnection()

	//Build the service
	crudService := di.BuildCrudService(db)

	startServer(crudService)
}

func startServer(service *di.CrudService) {
	r := router.Setup(service)

	r.Run()
}
