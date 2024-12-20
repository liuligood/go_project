package system_group_data_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"encoding/json"
	"go.uber.org/zap"
)

type GetValueListImpl interface {
	GetValueList(params request.GetGetValueListParams) (data response.GetGetValueListResult, err error)
}

type GetValueListService struct {
	svc *server.SvcContext
}

func NewGetValueListService(svc *server.SvcContext) *GetValueListService {
	return &GetValueListService{svc: svc}
}

func (g *GetValueListService) GetValueList(params request.GetGetValueListParams) (data response.GetGetValueListResult, err error) {
	SystemGroupDataList, err := g.svc.Repo.SystemGroupDataRepository.QueryValueByGid(params.BaseServiceParams.Ctx, params.Gid)
	if err != nil {
		izap.Log.Error("SystemGroupDataRepository.QueryValueByGid [http_err]:%v", zap.Error(err))

		return
	}

	if len(SystemGroupDataList) == 0 {
		return
	}

	valueDataList := make([]response.ValueData, 0, len(SystemGroupDataList))
	for _, v := range SystemGroupDataList {
		var valueData response.ValueData
		err = json.Unmarshal([]byte(v.Value), &valueData)
		if err != nil {
			izap.Log.Error("json.Unmarshal [http_err]:%v", zap.Error(err))

			return
		}

		valueDataList = append(valueDataList, valueData)
	}

	data = response.GetGetValueListResult{Data: map[int64][]response.ValueData{params.Gid: valueDataList}}
	return
}
