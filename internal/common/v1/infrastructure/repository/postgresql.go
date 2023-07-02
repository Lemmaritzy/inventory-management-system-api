package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"main/internal/common/v1/domain/ports"
)

type postgresqlRepo struct {
	ctx context.Context
	db  *pgxpool.Pool
}

func NewPostgresqlRepository(ctx context.Context, db *pgxpool.Pool) ports.IPostgresqlRepository {
	return &postgresqlRepo{
		ctx: ctx,
		db:  db,
	}
}
