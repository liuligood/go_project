package user_visit_repository

import (
	"context"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

// UserVisitRecordRepository UserVisitRepository 用户访问记录表
type UserVisitRecordRepository struct {
	*base_repository.DBRepository
}

// NewUserVisitRecordRepository NewUserVisitRepository 用户访问记录表
func NewUserVisitRecordRepository(db *gorm.DB, gen *gen.Query) *UserVisitRecordRepository {
	return &UserVisitRecordRepository{base_repository.NewRepository(db, gen)}
}

func (u *UserVisitRecordRepository) FindPageViewsByDate(ctx context.Context, start int64, end int64) (data int64, err error) {
	return u.Gen.WithContext(ctx).UserVisitRecord.Select(u.Gen.UserVisitRecord.UID).Where(u.Gen.UserVisitRecord.Date.Between(start, end)).Count()
}
