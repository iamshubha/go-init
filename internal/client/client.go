package client

import (
	"golang.org/x/exp/slog"
)

type ClientImpl interface {
}
type client struct {
	logger *slog.Logger
}

func NewClient(logger *slog.Logger) *client {
	return &client{
		logger: logger,
	}
}
