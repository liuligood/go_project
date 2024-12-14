package system_group_data_service

import (
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"encoding/json"
	"go.uber.org/zap"
)

type GetValueListImpl interface {
	GetValueList(params service_data.GetGetValueListParams) (data service_data.GetGetValueListResult, err error)
}

type GetValueListService struct {
	svc *server.SvcContext
}

func NewGetValueListService(svc *server.SvcContext) *GetValueListService {
	return &GetValueListService{svc: svc}
}

func (g *GetValueListService) GetValueList(params service_data.GetGetValueListParams) (data service_data.GetGetValueListResult, err error) {
	systemGroupDataList, err := g.svc.Repo.SystemGroupDataRepository.QueryValueByGid(params.BaseServiceParams.Ctx, params.Gid)
	if err != nil {
		izap.Log.Error("SystemGroupDataRepository.QueryValueByGid [err]:%v", zap.Error(err))

		return
	}

	if len(systemGroupDataList) == 0 {
		return
	}

	valueDataList := make([]service_data.ValueData, 0, len(systemGroupDataList))
	for _, v := range systemGroupDataList {
		var valueData service_data.ValueData
		err = json.Unmarshal([]byte(v.Value), &valueData)
		if err != nil {
			izap.Log.Error("json.Unmarshal [err]:%v", zap.Error(err))

			return
		}

		valueDataList = append(valueDataList, valueData)
	}

	data = service_data.GetGetValueListResult{Data: map[int64][]service_data.ValueData{params.Gid: valueDataList}}
	return
}
