// This is the server or waitress, they know the menu, and know what to serve, without needing to know the how and what is used
package di

import (
	"goCrud/handler"
	"goCrud/repository"
	"goCrud/usecase"

	"github.com/jmoiron/sqlx"
)

type CrudService struct {
	DB             *sqlx.DB
	CrudHandler    handler.CrudHandler
	CrudUsecase    usecase.CrudUsecase
	CrudRepository repository.Queries
}

func BuildCrudService(db *sqlx.DB) *CrudService {
	repo := repository.New(db)

	crudUsecase := usecase.NewCrudUseCase(repo)

	crudHandler := handler.NewCrudHandler(crudUsecase)

	return &CrudService{
		DB:             db,
		CrudRepository: repo,
		CrudUsecase:    crudUsecase,
		CrudHandler:    crudHandler,
	}
}
