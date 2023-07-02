package http

import (
	"context"
	"main/config"
	"main/internal/auth/v1/domain/ports"
	"main/pkg/logger"
)

type handler struct {
	ctx     context.Context
	cfg     *config.Config
	service ports.IService
	logger  logger.Logger
}

func NewHttpHandler(ctx context.Context, cfg *config.Config, service ports.IService, logger logger.Logger) ports.IHandlers {
	return &handler{
		ctx:     ctx,
		cfg:     cfg,
		service: service,
		logger:  logger,
	}
}
