package cron

import (
	"crmeb_go/internal/server"
	timer "crmeb_go/utils/itimer"
	"fmt"
	"github.com/robfig/cron/v3"
)

func Timer(svc *server.SvcContext) {
	// 初始化定时任务
	timer.NewTimerTask()

	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 测试任务 @daily
		if _, err := timer.TimerTask.AddTaskByFunc("test", "*/10 * * * * *", func() {
			NewTest(svc).Exec() // 定时任务方法定在task文件包中
		}, "test", option...); err != nil {
			fmt.Println("add timer error:", err)
		}

	}()
}
