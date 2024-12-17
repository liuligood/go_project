package service_data

import (
	"crmeb_go/internal/common/session_context"
	"github.com/gin-gonic/gin"
)

type BaseServiceParams struct {
	*session_context.SessionContext
}

func (s *BaseServiceParams) SetSessionContext(c *gin.Context) {
	s.SessionContext = session_context.GetSessionContext(c)
}
