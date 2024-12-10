package service

import (
	"be-golang-chapter-46-implem/repository"

	"go.uber.org/zap"
)

type AllService struct {
	VoucherService VoucherServiceInterface
	UserService    UserServiceInterface
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		VoucherService: NewVoucherService(repo, log),
		UserService:    NewUserService(repo, log),
	}
}
