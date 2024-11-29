package eb_user_service

import (
	"crmeb_go/internal/data/eb_user_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/model/eb_user_model"
	"crmeb_go/internal/server"
)

// GetRealNameImpl
type GetRealNameImpl interface {
	GetRealName(params service_data.GetRealNameParams) (data eb_user_data.GetRealNameData, err error)
}

// GetRealNameService
type GetRealNameService struct {
	svc *server.SvcContext
}

func NewGetRealNameService(svc *server.SvcContext) GetRealNameService {
	return GetRealNameService{svc: svc}
}

func (s GetRealNameService) GetRealName(params service_data.GetRealNameParams) (data eb_user_data.GetRealNameData, err error) {
	var userInfoModel eb_user_model.EbUserModel
	options := map[string]any{
		"select": "uid,real_name,account",
	}

	if err := s.svc.Repo.EbUserRepository.QueryOne(params.Ctx, "uid = ?", []any{params.UserId}, &userInfoModel, options); err != nil {
		return data, err
	}

	data.RealName = userInfoModel.RealName
	data.UserId = userInfoModel.UID

	return
}
