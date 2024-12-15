package system_menu_service

import (
	"crmeb_go/internal/data/admin"
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type PermissionByUserIdServiceImpl interface {
	GetPermissionByUserId(params service_data.BaseServiceParams) (resp []admin.Permission, err error)
}

type PermissionByUserIdService struct {
	svc *server.SvcContext
}

func NewPermissionByUserIdService(svc *server.SvcContext) *PermissionByUserIdService {
	return &PermissionByUserIdService{svc: svc}
}

func (p *PermissionByUserIdService) GetPermissionByUserId(params service_data.BaseServiceParams) (resp []admin.Permission, err error) {
	menus, err := p.svc.Repo.SystemMenuRepository.QueryMenuByUserId(params.Ctx, params.LoginUserInfo.UserId)
	if err != nil {
		izap.Log.Error("SystemMenuRepository.QueryMenuByUserId", zap.Int64("userId", params.LoginUserInfo.UserId), zap.Error(err))

		return
	}

	permissionList := make([]admin.Permission, 0, len(menus))
	lo.ForEach(menus, func(item *model.SystemMenu, index int) {
		var permission admin.Permission
		permission.Path = item.Perms
		permission.Id = item.ID
		permission.Pid = item.Pid
		permission.Sort = item.Sort
		permission.Name = item.Name

		permissionList = append(permissionList, permission)
	})

	resp = permissionList
	return
}
