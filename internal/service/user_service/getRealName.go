package user_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type GetRealNameImpl interface {
	GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error)
}

type GetRealNameService struct {
	svc *server.SvcContext
}

func NewGetRealNameService(svc *server.SvcContext) GetRealNameService {
	return GetRealNameService{svc: svc}
}

func (s GetRealNameService) GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error) {
	userInfoModel, err := s.svc.Repo.UserRepository.FindRealName(params.Ctx, params.UserId)
	if err != nil {
		izap.Log.Error("EbUserRepository.QueryOne [errorm]:%v", zap.Error(err))

		return
	}

	data.RealName = userInfoModel.RealName
	data.UserId = uint64(userInfoModel.UID)

	return
}
