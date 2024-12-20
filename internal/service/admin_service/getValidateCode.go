package admin_service

import (
	"context"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type GetValidateCodeGetRealNameImpl interface {
	GetValidateCode(context context.Context) (data *response.ValidateCodeResp, err error)
}

type GetValidateCodeService struct {
	svc *server.SvcContext
}

func NewGetValidateCodeService(svc *server.SvcContext) *GetValidateCodeService {
	return &GetValidateCodeService{svc: svc}
}

func (a *GetValidateCodeService) GetValidateCode() (data *response.ValidateCodeResp, err error) {
	err, code, id := a.svc.CaptchaClient.GetCaptcha()
	if err != nil {
		izap.Log.Error("生成验证码错误err:", zap.Error(err))

		return
	}

	return &response.ValidateCodeResp{Code: code, Key: id}, err
}
