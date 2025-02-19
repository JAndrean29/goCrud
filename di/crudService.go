// This is the server or waitress, they know the menu, and know what to serve, without needing to know the how and what is used
package di

import (
	"goCrud/config"
	"goCrud/handler"
	"goCrud/repository"
	"goCrud/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CrudService struct {
	Cfg            *config.Cfg
	DB             *pgxpool.Pool
	CrudHandler    handler.CrudHandler
	CrudUsecase    usecase.CrudUsecase
	CrudRepository repository.Queries
}

func BuildCrudService(c *config.Cfg, db *pgxpool.Pool) *CrudService {
	repo := repository.New(db)

	crudUsecase := usecase.NewCrudUseCase(c, repo)

	crudHandler := handler.NewCrudHandler(crudUsecase)

	return &CrudService{
		DB:             db,
		CrudRepository: *repo,
		CrudUsecase:    crudUsecase,
		CrudHandler:    crudHandler,
	}
}
