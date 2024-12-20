package admin_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
	"strings"
)

type GetMenusImpl interface {
	GetMenus(params request.GetMenusParams) (resp []*response.GetMenusResp, err error)
}

type GetMenusService struct {
	svc *server.SvcContext
}

func NewGetMenusService(svc *server.SvcContext) *GetMenusService {
	return &GetMenusService{svc: svc}
}

func (g *GetMenusService) GetMenus(params request.GetMenusParams) (resp []*response.GetMenusResp, err error) {
	var systemMenus []*model.SystemMenu

	if strings.Contains(params.BaseServiceParams.LoginUserInfo.Roles, "1") {
		// 超管获取全部
		systemMenus, err = g.svc.Repo.SystemMenuRepository.QueryAllMenuByA(params.BaseServiceParams.Ctx)
		if err != nil {
			izap.Log.Error("查询菜单错误:", zap.Error(err))

			return
		}
	} else {
		// 非超管
		systemMenus, err = g.svc.Repo.SystemMenuRepository.QueryMenusByUserId(params.BaseServiceParams.Ctx, params.BaseServiceParams.LoginUserInfo.UserId)
		if err != nil {
			izap.Log.Error("根据用户id查询菜单错误:", zap.Any("userId", params.BaseServiceParams.LoginUserInfo.UserId), zap.Error(err))

			return
		}

	}

	return g.buildMenuTree(systemMenus), nil
}

// ToTree 转换为树形结构
func (g *GetMenusService) buildMenuTree(systemMenus []*model.SystemMenu) []*response.GetMenusResp {
	menuMap := make(map[int64]*response.GetMenusResp)
	var menuTree []*response.GetMenusResp

	// 第一次遍历，创建所有的 GetMenusResp 对象并存入 map
	for _, menu := range systemMenus {
		getMenusResp := &response.GetMenusResp{
			ID:        int(menu.ID),
			Pid:       int(menu.Pid),
			Name:      menu.Name,
			Icon:      menu.Icon,
			Perms:     menu.Perms,
			Component: menu.Component,
			MenuType:  menu.MenuType,
			Sort:      int(menu.Sort),
			ChildList: nil, // 初始时子列表为空
		}
		menuMap[menu.ID] = getMenusResp
	}

	// 第二次遍历，建立父子关系
	for _, menu := range systemMenus {
		if menu.Pid == 0 { // 或者其他表示顶级菜单的条件
			menuTree = append(menuTree, menuMap[menu.ID])
		} else if parentMenu, exists := menuMap[menu.Pid]; exists {
			if parentMenu.ChildList == nil {
				parentMenu.ChildList = &response.RoleTrees{}
			}
			*parentMenu.ChildList = append(*parentMenu.ChildList, menuMap[menu.ID])
		}
	}

	return menuTree
}
