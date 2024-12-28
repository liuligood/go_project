package system_admin_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"strings"

	oldGen "gorm.io/gen"
	"gorm.io/gorm"
)

type SystemAdminRepository struct {
	*base_repository.DBRepository
}

func NewSystemAdminRepository(db *gorm.DB, gen *gen.Query) *SystemAdminRepository {
	return &SystemAdminRepository{base_repository.NewRepository(db, gen)}
}

func (r *SystemAdminRepository) FindAdminByAccount(ctx context.Context, account string) (data *model.SystemAdmin, err error) {
	return r.Gen.SystemAdmin.WithContext(ctx).Where(r.Gen.SystemAdmin.Account.Eq(account)).First()
}

func (r *SystemAdminRepository) FindAdminById(ctx context.Context, userId int64) (data *model.SystemAdmin, err error) {
	return r.Gen.SystemAdmin.WithContext(ctx).Where(r.Gen.SystemAdmin.ID.Eq(userId)).First()
}

func (r *SystemAdminRepository) UpdateSystemAdminIpCount(ctx context.Context, id int64, requestHost string) (resultInfo oldGen.ResultInfo, err error) {
	// 去掉端口号
	var clientIP string
	if strings.Contains(requestHost, ":") {
		ipParts := strings.Split(requestHost, ":")
		clientIP = ipParts[0]
	}

	return r.Gen.SystemAdmin.WithContext(ctx).Where(r.Gen.SystemAdmin.ID.Eq(id)).
		UpdateSimple(r.Gen.SystemAdmin.LoginCount.Add(1), r.Gen.SystemAdmin.LastIP.Value(clientIP))

}
