package admin_service

import (
	"context"
	"crmeb_go/internal/data/admin_data"
	"crmeb_go/internal/server"
	"crmeb_go/utils/ibase64_captcha"
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

func (a GetValidateCodeService) GetValidateCode() (admin_data.ValidateCodeData, error) {
	var data admin_data.ValidateCodeData
	captchaClient := ibase64_captcha.NewCaptchaClient(a.svc.RedisClient, a.svc.Ctx)

	err, code, id := captchaClient.GetCaptcha()
	if err != nil {
		izap.Log.Error("生成验证码错误err:", zap.Error(err))

		return data, err
	}

	data.Code = code
	data.Key = id

	return data, err
}
