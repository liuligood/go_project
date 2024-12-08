package service_data

import (
	"crmeb_go/internal/common/session_context"
	"github.com/gin-gonic/gin"
)

type BaseServiceParams struct {
	*session_context.SessionContext
}

func (s *BaseServiceParams) SetSessionContext(c *gin.Context) {
	sessionContext := session_context.GetSessionContext(c)
	//sessionContext.EnjoyMeta = enjoy_meta_context.NewEnjoyMetaContext(sessionContext.Svc, sessionContext.Ctx)
	s.SessionContext = sessionContext
}
