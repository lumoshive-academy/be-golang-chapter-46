package repository

import (
	"be-golang-chapter-46-implem/model"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Create(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type UserRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepoInterface {
	return &UserRepository{
		DB:     db,
		Logger: log,
	}
}

// Login: Memverifikasi kredensial pengguna
func (userRepo *UserRepository) Login(email, password string) (*model.User, error) {
	var user model.User
	// Cari user berdasarkan email
	if err := userRepo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userRepo.Logger.Warn("User not found", zap.String("email", email))
			return nil, errors.New("invalid email or password")
		}
		userRepo.Logger.Error("Database error during login", zap.Error(err))
		return nil, err
	}

	userRepo.Logger.Info("User logged in successfully", zap.String("email", user.Email))
	return &user, nil
}

func (userRepo *UserRepository) Create(user *model.User) error {
	// Gunakan transaksi untuk memastikan integritas data
	err := userRepo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			userRepo.Logger.Error("Failed to create user", zap.Error(err))
			return err
		}
		return nil
	})

	// Kembalikan error jika transaksi gagal
	if err != nil {
		return err
	}

	userRepo.Logger.Info("User created successfully", zap.String("email", user.Email))
	return nil
}
