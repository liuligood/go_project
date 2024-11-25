package user_repository

import (
	"crmeb_go/internal/model/user_model"
	"crmeb_go/internal/repository"
	"go.uber.org/zap"
)

type UserRepository struct {
	*repository.DBRepository
}

func NewUserRepository(model *user_model.UserModel, log *zap.Logger) *UserRepository {
	return &UserRepository{DBRepository: repository.NewDBRepository(model, log)}
}
