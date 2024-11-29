package middleware

import (
	"crmeb_go/internal/common/session_context"
	"crmeb_go/internal/data/ctx_key_data"
	"crmeb_go/internal/server"
	"github.com/gin-gonic/gin"
)

func ApiAuthVisitorMiddleWare(svc *server.SvcContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数校验
		//var param validation.AuthMiddlewareParams
		//var err error
		//if err = c.ShouldBindHeader(&param); err != nil {
		//	c.JSON(http.StatusOK, xhttp.Error(err))
		//	c.Abort()
		//	return
		//}
		//// token校验
		//thirdUserId, err := svc.Tool.TokenTool.GetTokenUserId(param.Token)
		//if err != nil {
		//	c.JSON(http.StatusOK, xhttp.Error(err, http_code_data.StatusUnSilentAuthorized))
		//	c.Abort()
		//	return
		//}
		// 将 Body 内容写回，以便后续的绑定操作可以读取

		// 存放登录信息到会话上下文
		sessionContext := session_context.GetSessionContext(c)
		c.Set(ctx_key_data.SessionContext, sessionContext)
		c.Next()
	}
}
