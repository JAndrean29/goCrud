package postgresql

import (
	"context"
	"fmt"
	"goCrud/config"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Cfg) (*pgx.Conn, error) {
	dsn := constructUrl(&cfg.Database)

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("connection failed, error: %v", err)
	}

	return db, nil
}

func NewPool(cfg *config.Cfg) (*pgxpool.Pool, error) {
	pgConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.Database, cfg.Database.User, cfg.Database.Password, cfg.Database.SSLMODE))
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func constructUrl(c *config.DbConfig) string {
	return fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s sslmode=%s", c.Host, c.Port, c.Database, c.User, c.Password, c.SSLMODE)
}
