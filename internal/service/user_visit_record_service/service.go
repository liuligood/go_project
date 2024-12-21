package user_visit_record_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/itime"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type UserVisitRecordServiceImpl interface {
	FindPageViewsByDate(params *request.DateParams) (data *response.DateResp, err error)
}

type UserVisitRecordService struct {
	svc *server.SvcContext
}

func NewUserVisitRecordService(svc *server.SvcContext) *UserVisitRecordService {
	return &UserVisitRecordService{svc: svc}
}

func (u *UserVisitRecordService) FindPageViewsByDate(params *request.DateParams) (*response.DateResp, error) {
	var data response.DateResp
	// 今日用户访问量
	nowPageViews, err := u.svc.Repo.UserVisitRecordRepository.FindPageViewsByDate(params.BaseServiceParams.Ctx, params.Start, params.End)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.Start)), zap.String("结束时间", itime.Format(params.End)), zap.Error(err))

		return &data, err
	}

	// 昨天用户访问量
	yesterdayPageViews, err := u.svc.Repo.UserVisitRecordRepository.FindPageViewsByDate(params.BaseServiceParams.Ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.YesterdayStart)), zap.String("结束时间", itime.Format(params.YesterdayEnd)), zap.Error(err))

		return &data, err
	}

	data.NowData = nowPageViews
	data.YesterdayData = yesterdayPageViews
	return &data, nil
}
