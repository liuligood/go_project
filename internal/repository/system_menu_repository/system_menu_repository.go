package system_menu_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/model/model_data"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SystemMenuRepository struct {
	*base_repository.DBRepository
}

func NewSystemMenuRepository(db *gorm.DB, gen *gen.Query) *SystemMenuRepository {
	return &SystemMenuRepository{base_repository.NewRepository(db, gen)}
}

func (s *SystemMenuRepository) FindPermissionByUserId(ctx context.Context, userId int64) (data []*model.SystemMenu, err error) {
	return s.Gen.WithContext(ctx).SystemMenu.Debug().FindPermissionByUserId(model_data.UserIdCondition{UserID: userId})
}

func (s *SystemMenuRepository) FindAllMenu(ctx context.Context) (data []*model.SystemMenu, err error) {
	return s.Gen.WithContext(ctx).SystemMenu.Where(s.Gen.SystemMenu.MenuType.Neq("M")).Find()
}

func (s *SystemMenuRepository) FindAllMenuByA(ctx context.Context) (data []*model.SystemMenu, err error) {
	return s.Gen.WithContext(ctx).SystemMenu.Where(s.Gen.SystemMenu.MenuType.Neq("A"),
		s.Gen.SystemMenu.IsShow.Eq(1)).Find()
}

func (s *SystemMenuRepository) FindMenusByUserId(ctx context.Context, userId int64) (data []*model.SystemMenu, err error) {
	return s.Gen.WithContext(ctx).SystemMenu.Debug().QueryMenusByUserId(model_data.UserIdCondition{UserID: userId})
}
