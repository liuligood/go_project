package repository

import (
	"crmeb_go/internal/model/user_model"
	"crmeb_go/internal/repository/user_repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	UserRepository *user_repository.UserRepository
}

func Register(db *gorm.DB, log *zap.Logger) *Container {
	return &Container{
		UserRepository: user_repository.NewUserRepository(user_model.NewUserModel(db), log),
	}
}
