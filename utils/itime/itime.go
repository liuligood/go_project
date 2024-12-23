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
