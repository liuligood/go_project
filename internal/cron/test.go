package cron

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Test struct {
	ctx    context.Context
	svcCtx *server.SvcContext
}

func NewTest(svcCtx *server.SvcContext) *Test {
	ctx := context.Background()

	return &Test{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Test) Exec() {
	fmt.Println(time.Now())
	fmt.Println("测试")
	_, err := l.svcCtx.Gen.User.WithContext(l.ctx).Where(l.svcCtx.Gen.User.UID.Eq(int64(1))).First()
	if err != nil {
		fmt.Println(err)
	}

	err = l.svcCtx.Gen.User.WithContext(l.ctx).Create(&model.User{UID: 1, Nickname: "123", Phone: "12345678901"})
	if err != nil {
		izap.Log.Error("create user error", zap.Error(err))
		fmt.Println(err)
	}

	var user model.User
	err = l.svcCtx.Repo.UserRepository.QueryOne(l.ctx, "uid = ?", []interface{}{1}, &user)
	if err != nil {
		fmt.Println("errorm", err)
	}

}
