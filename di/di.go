package di

import (
	"goCrud/handler"

	"github.com/jmoiron/sqlx"
)

type CrudService struct {
	DB 	*sqlx.DB
	CrudHandler handler.CrudHandler
}

func BuildCrudService
