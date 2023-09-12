package service

import (
	"github.com/iamshubha/go-init/internal/client"
	"github.com/iamshubha/go-init/internal/dao"
	"golang.org/x/exp/slog"
)

type ServiceImpl interface {
}

type service struct {
	client client.ClientImpl
	db     dao.DBImpl
	logger *slog.Logger
}

func NewService(client client.ClientImpl, db dao.DBImpl, logger *slog.Logger) *service {
	return &service{
		client: client,
		db:     db,
		logger: logger,
	}
}
