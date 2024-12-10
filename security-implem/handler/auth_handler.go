package handler

import (
	"be-golang-chapter-46-implem/database"
	"be-golang-chapter-46-implem/helper"
	"be-golang-chapter-46-implem/infra/jwt"
	"be-golang-chapter-46-implem/model"
	"be-golang-chapter-46-implem/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Cacher  database.Cacher
	Jwt     jwt.JWT
}

func NewAuthHandler(service service.AllService, log *zap.Logger, rdb database.Cacher, jwt jwt.JWT) AuthHandler {
	return AuthHandler{
		Service: service,
		Log:     log,
		Cacher:  rdb,
		Jwt:     jwt,
	}
}

// @summary Login user
// @description Login user
// @tags auth
// @Accept json
// @Produce json
// @param login body model.LoginRequest true "user data"
// @Success 201 {object} helper.Response "Login success"
// @Success 401 {object} helper.Response "Unauthorized"
// @Success 500 {object} helper.Response "server error"
// @Router /auth/login [post]
func (auth *AuthHandler) Login(c *gin.Context) {
	var loginUser model.LoginUser

	// Bind JSON body ke struct loginUser
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		auth.Log.Error("Invalid request body", zap.Error(err))
		helper.BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	// Ambil IP address dari request
	ip := c.ClientIP()
	if ip == "" {
		ip = "unknown" // Default jika IP tidak ditemukan
		auth.Log.Warn("Failed to retrieve client IP")
	}

	// Validasi login melalui service
	user, err := auth.Service.UserService.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		auth.Log.Warn("Login failed", zap.Error(err))
		helper.BadResponse(c, "invalid email or password", http.StatusUnauthorized)
		return
	}

	// Buat token JWT
	token, err := auth.Jwt.CreateToken(user.Email, ip, user.ID)
	if err != nil {
		auth.Log.Error("Failed to create JWT token", zap.Error(err))
		helper.BadResponse(c, "failed to create token", http.StatusInternalServerError)
		return
	}

	// Buat response data
	data := gin.H{
		"user":  user,
		"token": token,
	}

	auth.Log.Info("User logged in successfully", zap.String("email", user.Email))
	helper.SuccessResponseWithData(c, "success", http.StatusOK, data)
}

func (auth *AuthHandler) Register(c *gin.Context) {
	var user model.User

	// Bind JSON body ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		helper.BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	user.Password = helper.HashPassword(user.Password)

	// Panggil service untuk registrasi user
	if err := auth.Service.UserService.Register(&user); err != nil {
		helper.BadResponse(c, "failed to register user", http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	helper.SuccessResponse(c, "success register", http.StatusCreated)

}

func (auth *AuthHandler) AddIp(c *gin.Context) {
	var request struct {
		IPs []string `json:"ips"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.BadResponse(c, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Check if IP list is not empty
	if len(request.IPs) == 0 {
		helper.BadResponse(c, "IP list cannot be empty", http.StatusBadRequest)
		return
	}

	// Add each IP to the whitelist
	err := auth.Cacher.SAdd("whitelist", request.IPs...)
	if err != nil {
		helper.BadResponse(c, "Failed to add IPs to whitelist", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(c, fmt.Sprintf("IPs %v added to whitelist", request.IPs), http.StatusOK)
}
