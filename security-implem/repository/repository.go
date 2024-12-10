package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AllRepository struct {
	VoucherRepo VoucherRepoInterface
	UserRepo    UserRepoInterface
}

func NewAllRepository(db *gorm.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		VoucherRepo: NewVoucherRepository(db, log),
		UserRepo:    NewUserRepository(db, log),
	}
}
