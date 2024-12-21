package itime

import (
	"time"
)

func Format(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(time.DateTime)
}
