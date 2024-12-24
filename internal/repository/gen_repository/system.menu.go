package gen_repository

import (
	"crmeb_go/internal/model"
	"crmeb_go/internal/model/model_data"
)

type SystemMenuQuerier interface {
	//	SELECT m.* FROM system_menu as m
	//  right join system_role_menu as  rm on rm.menu_id = m.id
	//  right join system_role as  r on rm.rid = r.id
	//  right join system_admin as  a on FIND_IN_SET(r.id, a.roles)
	// 		{{where}}
	// 			{{if condition.UserID !=0}}
	// 				a.id = @condition.UserID AND
	// 			{{end}}
	// 			m.deleted_at = 0 AND r.status =1
	// 		{{end}}
	// GROUP BY m.id
	FindPermissionByUserId(condition model_data.UserIdCondition) ([]*model.SystemMenu, error)

	//	SELECT m.* FROM system_menu as m
	//  right join system_role_menu as  rm on rm.menu_id = m.id
	//  right join system_role as  r on rm.rid = r.id
	//  right join system_admin as  a on FIND_IN_SET(r.id, a.roles)
	// 		{{where}}
	// 			{{if condition.UserID !=0}}
	// 				a.id = @condition.UserID AND
	// 			{{end}}
	// 			m.deleted_at = 0 AND r.status =1 and m.menu_type != 'A' and m.is_show =1
	// 		{{end}}
	// GROUP BY m.id
	QueryMenusByUserId(condition model_data.UserIdCondition) ([]*model.SystemMenu, error)
}
