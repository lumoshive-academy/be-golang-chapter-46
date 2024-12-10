package handler

import (
	"be-golang-chapter-46-implem/database"
	"be-golang-chapter-46-implem/infra/jwt"
	"be-golang-chapter-46-implem/service"

	"go.uber.org/zap"
)

type AllHandler struct {
	VoucherHandler VoucherHadler
	Auth           AuthHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, rdb database.Cacher, jwt jwt.JWT) AllHandler {
	return AllHandler{
		VoucherHandler: NewVoucherHandler(service, log),
		Auth:           NewAuthHandler(service, log, rdb, jwt),
	}
}
