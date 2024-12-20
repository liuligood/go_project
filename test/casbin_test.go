package test

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

const MySQLDSN = "root:dazhou520@tcp(127.0.0.1:3306)/crmeb?charset=utf8mb4&parseTime=True"

func Test(t *testing.T) {
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	adapterByDB, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		fmt.Printf("NewAdapterByDB http_err:【%v】", err)

		panic(err)
	}
	enforcer, err := casbin.NewEnforcer("D:\\go\\crmeb_go\\utils\\rabc\\rbac_model.conf", adapterByDB, true)
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

	// 添加策略
	enforcer.AddPolicy("dazhou", "data1", "read")
	// 判断策略是否存在
	policy, err := enforcer.HasPolicy("dazhou", "data1", "read")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("是否存在", policy)

	enforcer.AddPolicy("dazhou2", "data1", "read")

	// 添加角色
	enforcer.AddRoleForUser("dazhou", "admin")
	enforcer.AddRoleForUser("dazhou", "admin2")
	// 删除角色
	enforcer.DeleteRole("admin2")
	// 获取所有角色
	roles, err := enforcer.GetAllRoles()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("roles", roles)

	// 获取全部policy
	getPolicy, err := enforcer.GetPolicy()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("getPolicy", getPolicy)

	enforcer.UpdatePolicy([]string{"dazhou2", "data1", "read"}, []string{"dazhou2", "data1", "write"})
	// 删除策略
	enforcer.RemovePolicy("dazhou2", "data1", "write")

	policy2, err := enforcer.HasPolicy("dazhou2", "data1", "write")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("是否存在,policy2", policy2)

	ok, err := enforcer.Enforce("dazhou", "data1", "read")
	if err != nil {
		fmt.Println(err)
	}

	if ok {
		fmt.Println("判断成功")
	} else {
		fmt.Println("判断失败")
	}

}
