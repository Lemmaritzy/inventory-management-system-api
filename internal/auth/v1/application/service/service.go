package service

import (
	"context"
	"main/config"
	"main/internal/auth/v1/domain/ports"
	"main/pkg/logger"
)

type service struct {
	ctx    context.Context
	cfg    *config.Config
	pgRepo ports.IPostgresqlRepository
	logger logger.Logger
}

func NewService(ctx context.Context, cfg *config.Config, pgRepo ports.IPostgresqlRepository, logger logger.Logger) ports.IService {
	return &service{
		ctx:    ctx,
		cfg:    cfg,
		pgRepo: pgRepo,
		logger: logger,
	}
}
