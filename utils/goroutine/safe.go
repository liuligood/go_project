package goroutine

import (
	"crmeb_go/define"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
	"time"
)

func SafeGo(handler func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				izap.Log.Error("safe go handler panic", zap.String("timestamp", time.Now().Format(define.SystemTimeFormat)),
					zap.Any("recovered", e), zap.Stack("stack"))
			}
		}()
		handler()
	}()
}
