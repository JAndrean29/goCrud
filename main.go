// This is the customer, asking for a ready to eat meal
package main

import (
	"fmt"
	"goCrud/config"
	"goCrud/di"
	"goCrud/handler/router"
	"goCrud/infrastructure/postgresql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//load config
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error loading config: %v", err)
	}
	fmt.Printf("loaded the following config: %s\n", cfg.Database)

	if cfg.Database.User == " " || cfg.Database.Database == " " {
		log.Fatalf("missing config: %+v\n", cfg.Database)
	}

	//init DB
	db, err := postgresql.NewPool(cfg)
	if err != nil {
		log.Fatalf("error loading db: %v", err)
	}

	log.Println("successfully connected to postgresql db!")

	//Build the service
	crudService := di.BuildCrudService(cfg, db)

	startServer(crudService)
}

func startServer(service *di.CrudService) {
	r := router.Setup(service)

	r.Run()
}

/*
useful commands:
docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
*/
