package system_store_staff_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"github.com/samber/lo"
)

type SystemStoreStaffListServiceImpl interface {
	SystemStoreStaffList(params *request.SystemStaffListParams) (data *response.SystemStaffListResp, err error)
}

type SystemStoreStaffListService struct {
	svc *server.SvcContext
}

func NewSystemStoreStaffListService(svc *server.SvcContext) *SystemStoreStaffListService {
	return &SystemStoreStaffListService{svc: svc}
}

func (s *SystemStoreStaffListService) SystemStoreStaffList(params *request.SystemStaffListParams) (*response.SystemStaffListResp, error) {
	var err error
	var data response.SystemStaffListResp

	systemStoreStaff, err := s.svc.Repo.SystemStoreStaffRepository.FindListByStoreId(params.BaseServiceParams.Ctx, params.StoreId)
	if err != nil {
		return &data, err
	}

	if len(systemStoreStaff) == 0 {
		return &data, nil
	}

	storeMap := make(map[int64]*model.SystemStore, len(systemStoreStaff))
	userMap := make(map[int64]*model.User, len(systemStoreStaff))

	count, err := s.svc.Repo.SystemStoreStaffRepository.FindListByStoreIdCount(params.BaseServiceParams.Ctx, params.StoreId)
	if err != nil {
		return &data, err
	}

	userIdList := lo.Map(systemStoreStaff, func(item *model.SystemStoreStaff, index int) int64 {
		return item.UID
	})
	storeIdList := lo.Map(systemStoreStaff, func(item *model.SystemStoreStaff, index int) int64 {
		return item.StoreID
	})

	if len(userIdList) > 0 {
		userList, err := s.svc.Repo.UserRepository.FindListById(params.BaseServiceParams.Ctx, userIdList)
		if err != nil {
			return &data, err
		}

		for _, user := range userList {
			userMap[user.UID] = user
		}
	}

	if len(storeIdList) > 0 {
		storeList, err := s.svc.Repo.SystemStoreRepository.FindListById(params.BaseServiceParams.Ctx, storeIdList)
		if err != nil {
			return &data, err
		}

		for _, store := range storeList {
			storeMap[store.ID] = store
		}
	}

	for _, v := range systemStoreStaff {
		var systemStaffListData response.SystemStaffListData
		systemStaffListData.Marshal(v)

		if mStore, ok := storeMap[v.StoreID]; ok {
			var store response.SystemStore
			store.Marshal(mStore)
			systemStaffListData.SystemStore = store
		}

		if mUser, ok := userMap[v.UID]; ok {
			var user response.User
			user.Marshal(mUser)
			systemStaffListData.User = user
		}

		data.List = append(data.List, systemStaffListData)
	}

	data.Count = count

	return &data, nil
}
