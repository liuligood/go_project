package itime

import (
	"crmeb_go/define"
	"time"
)

func Format(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(define.SystemTimeFormat)
}
