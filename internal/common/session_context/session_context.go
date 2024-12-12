package session_context

import (
	"context"
	"crmeb_go/define"
	"crmeb_go/internal/data/common/session"
	"github.com/gin-gonic/gin"
)

// 会话上下文
type SessionContext struct {
	GinCtx *gin.Context
	Ctx    context.Context
	Svc    *session.SvcContext
	//LoginUserInfo *session_data.LoginUserInfo
	//EnjoyMeta     *enjoy_meta_context.EnjoyMeta
}

func New(ginCtx *gin.Context, svc *session.SvcContext) *SessionContext {
	return &SessionContext{
		GinCtx: ginCtx,
		Ctx:    toContext(ginCtx),
		Svc:    svc,
	}
}

func toContext(ginCtx *gin.Context) context.Context {
	requestCtx := ginCtx.Request.Context()
	newCtx := context.Background() // 请求取消会导致后续数据库操作事务回滚，暂不使用request.context
	if ginCtx.Keys == nil {
		return newCtx
	}

	// 遍历 gin.Context.Keys，将每个键值对添加到新的 context.Context 中
	for key, value := range ginCtx.Keys {
		if newCtx.Value(key) == nil {
			newCtx = context.WithValue(newCtx, key, value)
		}
	}

	// host
	if host, ok := requestCtx.Value("request_host").(string); ok {
		newCtx = context.WithValue(newCtx, "request_host", host)
	}

	//// trace
	//span := trace.SpanFromContext(requestCtx)
	//if span != nil {
	//	newCtx = trace.ContextWithSpan(newCtx, span)
	//}

	return newCtx

}

// 获取登录用户信息
//func GetLoginUserInfo(c *gin.Context) *session_data.LoginUserInfo {
//	return GetSessionContext(c).LoginUserInfo
//}

// 获取MerchId
//func GetUserMerchId(c *gin.Context) string {
//	return GetSessionContext(c).LoginUserInfo.CodeMerchInfo.MerchId
//}

// 获取会话上下文
func GetSessionContext(c *gin.Context) *SessionContext {
	sessionContextAny, _ := c.Get(define.SystemSessionContext)
	// 注入 request_host
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "request_host", c.Request.Host))
	sessionContext := sessionContextAny.(*SessionContext)
	sessionContext.Ctx = toContext(c)

	return sessionContext
}
