package repository

import (
	"be-golang-chapter-46-implem/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VoucherRepoInterface interface {
	Create(customer *model.Voucher) error
	GetAll() (*[]model.Voucher, error)
	GetByID(id int) (*model.Voucher, error)
}

type VoucherRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewVoucherRepository(db *gorm.DB, log *zap.Logger) VoucherRepoInterface {
	return &VoucherRepository{
		DB:     db,
		Logger: log,
	}
}

func (voucherRepo *VoucherRepository) Create(shipping *model.Voucher) error {
	return nil
}

func (voucherRepo *VoucherRepository) GetAll() (*[]model.Voucher, error) {
	var vouchers []model.Voucher

	return &vouchers, nil
}

func (voucherRepo *VoucherRepository) GetByID(id int) (*model.Voucher, error) {
	var voucher model.Voucher

	return &voucher, nil
}
