package izap

import (
	"crmeb_go/config"
	"crmeb_go/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewZap(z config.Zap) *zap.Logger {
	if ok, _ := utils.PathExists(z.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", z.Director)
		_ = os.Mkdir(z.Director, os.ModePerm)
	}
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

	return logger
}
