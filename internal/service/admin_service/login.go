package admin_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/ijwt"
	"crmeb_go/utils/imd5"
	"crmeb_go/utils/izap"
	"errors"
	"go.uber.org/zap"
)

type LoginServiceImpl interface {
	Login(params *request.LoginParams) (resp *response.LoginResp, err error)
}

type LoginService struct {
	svc *server.SvcContext
}

func NewLoginService(svc *server.SvcContext) *LoginService {
	return &LoginService{svc: svc}
}

func (l *LoginService) Login(params *request.LoginParams) (resp *response.LoginResp, err error) {
	// 校验验证码
	verify := l.svc.CaptchaClient.Verify(params.Key, params.Code)
	if !verify {
		return resp, errors.New("验证码错误")
	}

	// 校验用户名和密码
	systemAdmin, err := l.svc.Repo.SystemAdminRepository.FindAdminByAccount(params.Ctx, params.Account)
	if err != nil {
		izap.Log.Error("用户名或密码错误")

		return resp, errors.New("用户名或密码错误")
	}

	// 校验密码
	check := imd5.CheckPasswordHash(params.Pwd, systemAdmin.Pwd)
	if !check {
		return resp, errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := ijwt.GenerateToken(systemAdmin.ID, l.svc.Conf.JWT.AccessExpire, l.svc.Conf.JWT.AccessSecret, systemAdmin.Roles)
	if err != nil {
		return resp, err
	}

	// 更新最后登录信息
	l.svc.Repo.SystemAdminRepository.UpdateSystemAdminIpCount(params.Ctx, systemAdmin.ID, params.Ctx.Value(define.SystemRequestHost).(string))
	if err != nil {
		izap.Log.Error("更新最后登录信息失败", zap.Error(err))

		return resp, err
	}

	return &response.LoginResp{Id: systemAdmin.ID, Account: systemAdmin.Account, Token: token, RealName: systemAdmin.RealName, IsSms: false}, nil
}
