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

func (s *SystemGroupDataRepository) QueryValueByGid(ctx context.Context, gid int64) (data []*model.SystemGroupDatum, err error) {
	return s.Gen.WithContext(ctx).SystemGroupDatum.Where(s.Gen.SystemGroupDatum.Gid.Eq(gid),
		s.Gen.SystemGroupDatum.Status.Eq(define.SystemGroupDataOpenStatus)).
		Order(s.Gen.SystemGroupDatum.Sort.Asc(), s.Gen.SystemGroupDatum.ID.Asc()).Debug().
		Find()
}
