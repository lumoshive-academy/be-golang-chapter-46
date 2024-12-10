package service

import (
	"be-golang-chapter-46-implem/helper"
	"be-golang-chapter-46-implem/model"
	"be-golang-chapter-46-implem/repository"
	"errors"

	"go.uber.org/zap"
)

type UserServiceInterface interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type UserService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewUserService(repo repository.AllRepository, log *zap.Logger) UserServiceInterface {
	return &UserService{
		Repo: repo,
		Log:  log,
	}
}

// Login: Verifikasi kredensial user
func (userService *UserService) Login(email, password string) (*model.User, error) {
	userService.Log.Info("Attempting to log in user", zap.String("email", email))

	// Cari user berdasarkan email
	user, err := userService.Repo.UserRepo.Login(email, password)
	if err != nil {
		userService.Log.Error("Login failed", zap.Error(err))
		return nil, err
	}

	// Verifikasi password
	if !helper.CheckPassword(password, user.Password) {
		userService.Log.Warn("Invalid password", zap.String("email", email))
		return nil, errors.New("invalid email or password")
	}

	userService.Log.Info("User logged in successfully", zap.String("email", email))
	return user, nil
}

func (userService *UserService) Register(user *model.User) error {
	return userService.Repo.UserRepo.Create(user)
}
