package itime

import (
	"crmeb_go/define"
	"time"
)

func Format(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(time.DateTime)
}

func CalculateDateRange(data string) (int64, int64) {
	var startTime, endTime int64
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 今天0点
	nowTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	// 今天24点
	nowEndTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, now.Location()).Unix()

	switch data {
	case define.AdminSearchDateDay:
		startTime = nowTime
		endTime = nowEndTime
	case define.AdminSearchDateYesterday:
		date := today.Add(-24 * time.Hour)
		startTime = date.Unix()
		endTime = date.Add(-1 * time.Second).Unix()
	case define.AdminSearchDateLatelySeven:
		date := today.Add(-168 * time.Hour)
		startTime = date.Unix()
		endTime = today.Add(-1 * time.Second).Unix()
	case define.AdminSearchDateWeek:
		startTime = today.AddDate(0, 0, -int(today.Weekday())).Unix()
		endTime = now.Unix()
	case define.AdminSearchDatePreWeek:
		date := today.AddDate(0, 0, -int(today.Weekday())+1)
		startTime = date.Unix()
		endTime = date.AddDate(0, 0, 6).Unix()
	case define.AdminSearchDateLatelyThirty:
		date := today.AddDate(0, 0, -30)
		startTime = date.Unix()
		endTime = now.Unix()
	case define.AdminSearchDateMonth:
		date := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		startTime = date.Unix()
		endTime = now.Unix()
	case define.AdminSearchDatePreMonth:
		date := today.AddDate(0, -1, -now.Day()+1)
		startTime = date.Unix()
		endDate := date.AddDate(0, 1, -1)
		endTime = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 59, endDate.Location()).Unix()
	case define.AdminSearchDateYear:
	case define.AdminSearchDatePreYear:
	default:
		// 今天0点
		startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
		// 今天24点
		endTime = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, now.Location()).Unix()
	}

	return startTime, endTime
}

func getLastWeekStartEnd() (startTime, endTime time.Time) {
	now := time.Now().In(time.Local) // 使用本地时区的时间

	// 找到本周一
	today := now.Weekday()
	offsetToMonday := int(today) - 1 // 如果今天是周一，则offset为0；如果是周日，offset为6
	if offsetToMonday < 0 {
		offsetToMonday = 6 // 如果今天是周日，那么需要向前移动6天到达上一个星期的周一
	}

	// 从当前时间减去偏移量得到本周一的时间，并设置为当天的00:00:00
	startOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -offsetToMonday)

	// 上周的开始时间是本周一往前7天的午夜
	startTime = startOfWeek.AddDate(0, 0, -7)

	// 上周的结束时间是上周日的最后一秒
	endTime = startTime.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	return
}
