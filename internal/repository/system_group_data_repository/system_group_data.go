package system_group_data_repository

import (
	"context"
	"crmeb_go/define"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SystemGroupDataRepository struct {
	*base_repository.DBRepository
}

func NewSystemGroupDataRepository(db *gorm.DB, gen *gen.Query) *SystemGroupDataRepository {
	return &SystemGroupDataRepository{base_repository.NewRepository(db, gen)}
}

func (s *SystemGroupDataRepository) QueryValueByGid(ctx context.Context, gid int64) (data []*model.SystemGroupData, err error) {
	return s.Gen.WithContext(ctx).SystemGroupData.Where(s.Gen.SystemGroupData.Gid.Eq(gid),
		s.Gen.SystemGroupData.Status.Eq(define.SystemGroupDataOpenStatus)).
		Order(s.Gen.SystemGroupData.Sort.Asc(), s.Gen.SystemGroupData.ID.Asc()).Debug().
		Find()
}
