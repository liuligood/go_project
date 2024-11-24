package server

import (
	"context"
	"crmeb_go/config"
)

type SvcContext struct {
	Ctx  context.Context
	Conf config.Conf
}
