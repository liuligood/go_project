package home_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/itime"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
	"time"
)

type HomeServiceImpl interface {
	IndexDate(params *request.BaseServiceParams) (data *response.HomeRateResp, err error)
}

type HomeService struct {
	svc *server.SvcContext
}

func NewHomeService(svc *server.SvcContext) *HomeService {
	return &HomeService{svc: svc}
}

func (h *HomeService) IndexDate(params *request.BaseServiceParams) (data *response.HomeRateResp, err error) {
	var resp response.HomeRateResp

	now := time.Now()
	// 今天0点时间
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayUnix := today.Unix()
	// 计算昨天0点的时间
	yesterdayZero := today.Add(-24 * time.Hour).Unix()

	// 计算昨天24点的时间
	yesterdayEnd := today.Add(-1 * time.Second).Unix()

	// 现在的时间
	nowUnix := time.Now().Unix()

	// 今日用户访问量
	nowPageViews, err := h.svc.Repo.UserVisitRecordRepository.FindRegisterNumByDate(params.Ctx, todayUnix, nowUnix)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(todayUnix)), zap.String("结束时间", itime.Format(nowUnix)), zap.Error(err))

		return &resp, err
	}

	// 昨天用户访问量
	yesterdayPageViews, err := h.svc.Repo.UserVisitRecordRepository.FindRegisterNumByDate(params.Ctx, yesterdayZero, yesterdayEnd)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(todayUnix)), zap.String("结束时间", itime.Format(nowUnix)), zap.Error(err))

		return &resp, err
	}

	resp.PageViews = nowPageViews
	resp.YesterdayPageViews = yesterdayPageViews

	return &resp, nil
}
