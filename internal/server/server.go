package server

import (
	"context"
	"crmeb_go/config"
	"go.uber.org/zap"
)

type SvcContext struct {
	Ctx    context.Context
	Conf   config.Conf
	Logger *zap.Logger
}
