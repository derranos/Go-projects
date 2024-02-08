package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	pool *pgxpool.Pool
}

func New(connStr string) (*PGRepo, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &PGRepo{pool: pool}, nil
}

const ConnStr = "postgres://postgres:psql@localhost:5432/test"
