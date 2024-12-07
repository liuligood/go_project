package middleware

import (
	"crmeb_go/internal/common/session_context"
	"crmeb_go/internal/data/common_data/session_context_data"
	"crmeb_go/internal/data/ctx_key_data"
	"crmeb_go/internal/server"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(svc *server.SvcContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 存放登录信息到上下文
		sessionContext := session_context.New(c, &session_context_data.SvcContext{
			Repo:            svc.Repo,
			Conf:            svc.Conf,
			Gorm:            svc.Gorm,
			RedisClient:     svc.RedisClient,
			RedisClientList: svc.RedisClientList,
		})

		c.Set(ctx_key_data.SessionContext, sessionContext)

		c.Next()
	}
}
