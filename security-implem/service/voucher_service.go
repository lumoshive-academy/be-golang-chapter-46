package service

import (
	"be-golang-chapter-46-implem/model"
	"be-golang-chapter-46-implem/repository"

	"go.uber.org/zap"
)

type VoucherServiceInterface interface {
	Create(voucher *model.Voucher) error
	GetAll() (*[]model.Voucher, error)
}

type VoucherService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewVoucherService(repo repository.AllRepository, log *zap.Logger) VoucherServiceInterface {
	return &VoucherService{
		Repo: repo,
		Log:  log,
	}
}

func (voucherService *VoucherService) Create(voucher *model.Voucher) error {
	return nil
}

func (voucherService *VoucherService) GetAll() (*[]model.Voucher, error) {
	return voucherService.Repo.VoucherRepo.GetAll()
}
