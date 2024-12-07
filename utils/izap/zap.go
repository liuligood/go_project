package izap

import (
	"crmeb_go/config"
	"crmeb_go/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *zap.Logger

func NewZap(z config.Zap) {
	// 判断是否有Director文件夹 没有就创建
	if ok, _ := utils.PathExists(z.Director); !ok {
		fmt.Printf("create %v directory\n", z.Director)
		_ = os.Mkdir(z.Director, os.ModePerm)
	}

	// 据配置文件中的日志级别设置，创建多个core 实例
	levels := z.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)

	for i := 0; i < length; i++ {
		core := NewZapCore(&z, levels[i])
		cores = append(cores, core)
	}

	logger := zap.New(zapcore.NewTee(cores...))
	if z.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	Log = logger
}
