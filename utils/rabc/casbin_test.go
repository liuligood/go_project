package rabc

import (
	"crmeb_go/internal/model"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

const MySQLDSN = "root:dazhou520@tcp(127.0.0.1:3306)/crmeb?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return db
}

func Test(T *testing.T) {
	db := connectDB(MySQLDSN)

	adapterByDB, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		fmt.Printf("NewAdapterByDB http_err:【%v】", err)

		panic(err)
	}

	enforcer, err := casbin.NewEnforcer("utils/rabc/rbac_model.conf", adapterByDB)
	if err != nil {
		fmt.Printf("NewEnforcer http_err:【%v】", err)

		panic(err)
	}

	// 开启权限认证日志
	enforcer.EnableLog(true)

	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Printf("loadPolicy error")

		panic(err)
	}

	var roles []model.SystemRole
	if err := db.Find(&roles).Error; err != nil {
		fmt.Println("e")
	}

	var menus []model.SystemMenu
	if err := db.Find(&menus).Error; err != nil {
		fmt.Println("e2")
	}

	for _, role := range roles {
		// 查询角色关联的菜单
		var roleMenus []model.SystemRoleMenu
		if err := db.Where("rid = ?", role.ID).Find(&roleMenus).Error; err != nil {
			fmt.Printf("form error%v", err)
		}

		for _, rm := range roleMenus {
			lo.ForEach(menus, func(item model.SystemMenu, index int) {
				if item.ID == rm.MenuID {
					switch item.MenuType {
					case "M":
						_, err = enforcer.AddPolicy(role.RoleName, item.Component, "M")
						if err != nil {
							fmt.Printf("addPolicy error%v", err)
						}
					case "C":
						_, err = enforcer.AddPolicy(role.RoleName, item.Perms, "C")
						if err != nil {
							fmt.Printf("addPolicy error%v", err)
						}
					case "A":
						_, err = enforcer.AddPolicy(role.RoleName, item.Perms, "A")
						if err != nil {
							fmt.Printf("addPolicy error%v", err)
						}
					}

				}
			})
		}
	}

}
