package rabc

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

//var Enforcer *casbin.Enforcer

func NewEnforcer(db *gorm.DB) *casbin.Enforcer {
	adapterByDB, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		fmt.Printf("NewAdapterByDB errorm:【%v】", err)

		panic(err)
	}

	enforcer, err := casbin.NewEnforcer("utils/rabc/rbac_model.conf", adapterByDB)
	if err != nil {
		fmt.Printf("NewEnforcer errorm:【%v】", err)

		panic(err)
	}

	// 开启权限认证日志
	enforcer.EnableLog(true)

	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Printf("loadPolicy error")

		panic(err)
	}

	return enforcer
}
