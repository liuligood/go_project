package admin_service

import (
	"crmeb_go/define"
	service_data "crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/system_menu_service"
	"crmeb_go/utils/izap"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"strings"
	"time"
)

type LoginUserInfoServiceImpl interface {
	Login(params service_data.LoginParams) (resp response.LoginResp, err error)
}

type LoginUserInfoService struct {
	svc *server.SvcContext
}

func NewLoginUserInfoService(svc *server.SvcContext) *LoginUserInfoService {
	return &LoginUserInfoService{svc: svc}
}

func (l *LoginUserInfoService) LoginUserInfo(params service_data.LoginUserInfoParams) (resp response.LoginUserInfoResp, err error) {
	systemAdmin, err := l.svc.Repo.SystemAdminRepository.QueryAdminById(params.BaseServiceParams.Ctx, params.BaseServiceParams.LoginUserInfo.UserId)
	if err != nil {
		izap.Log.Error("SystemAdminRepository.QueryAdminById", zap.Int64("userId", params.BaseServiceParams.LoginUserInfo.UserId), zap.Error(err))

		return
	}

	roleList := strings.Split(systemAdmin.Roles, ",")
	permList := make([]string, 0, len(roleList))

	if lo.Contains(roleList, define.SystemAdminRoles) {
		permList = append(permList, define.SystemAdminAllPerm)
	} else {
		permissionList, err := system_menu_service.NewPermissionByUserIdService(l.svc).GetPermissionByUserId(params.BaseServiceParams)
		if err != nil {
			izap.Log.Error("PermissionByUserIdService.GetPermissionByUserId", zap.Int64("userId", params.BaseServiceParams.LoginUserInfo.UserId), zap.Error(err))

			return resp, err
		}

		permList = lo.Map(permissionList, func(item response.Permission, index int) string {
			return item.Path
		})
	}

	resp = response.LoginUserInfoResp{
		ID:              systemAdmin.ID,
		Account:         systemAdmin.Account,
		RealName:        systemAdmin.RealName,
		Roles:           systemAdmin.Roles,
		AddTime:         time.Unix(systemAdmin.CreatedAt, 0).Format(define.SystemTimeFormat),
		LoginCount:      systemAdmin.LoginCount,
		Level:           systemAdmin.Level,
		Status:          systemAdmin.Status == define.SystemAdminStatusValid,
		Phone:           systemAdmin.Phone,
		IsSms:           systemAdmin.IsSms == define.SystemAdminIsSmsValid,
		LastIP:          systemAdmin.LastIP,
		PermissionsList: permList,
	}
	return
}
