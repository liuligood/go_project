package admin_service

import (
	"context"
	"crmeb_go/internal/data/admin_data"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type GetValidateCodeGetRealNameImpl interface {
	GetValidateCode(context context.Context) (data admin_data.ValidateCodeData, err error)
}

type GetValidateCodeService struct {
	svc *server.SvcContext
}

func NewGetValidateCodeService(svc *server.SvcContext) *GetValidateCodeService {
	return &GetValidateCodeService{svc: svc}
}

func (a *GetValidateCodeService) GetValidateCode() (data admin_data.ValidateCodeData, err error) {
	err, code, id := a.svc.CaptchaClient.GetCaptcha()
	if err != nil {
		izap.Log.Error("生成验证码错误err:", zap.Error(err))

		return
	}

	return admin_data.ValidateCodeData{Code: code, Key: id}, err
}
