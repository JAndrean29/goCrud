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
	CrudRepository repository.CrudRepository
}

func BuildCrudService(db *sqlx.DB) *CrudService {
	repo := repository.NewCrudRepository(db)

	crudUsecase := usecase.NewCrudUseCase(repo)

	crudHandler := handler.NewCrudHandler(crudUsecase)

	return &CrudService{
		DB:             db,
		CrudRepository: repo,
		CrudUsecase:    crudUsecase,
		CrudHandler:    crudHandler,
	}
}
