package storage

import (
	"context"
	"english_bot/config"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPsqlDB(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.DBName)

	return sqlx.Connect("pgx", connectionUrl)
}

func InitConnectionPoolPsqlDB(c *config.Config) (*pgxpool.Pool, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s account=%s password=%s dbname=%s",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.DBName)

	return pgxpool.New(context.Background(), connectionUrl)
}
