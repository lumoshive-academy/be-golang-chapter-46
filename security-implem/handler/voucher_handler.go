package handler

import (
	"be-golang-chapter-46-implem/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoucherHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewVoucherHandler(service service.AllService, log *zap.Logger) VoucherHadler {
	return VoucherHadler{
		Service: service,
		Log:     log,
	}
}

// @summary API created shipping
// @description API create shipping
// @tags shipping
// @Produce json
// @param token header string true "Token to be passed after login"
// @param ID-KEY header string true "ID Key to be passed after login"
// @param shipping body model.Shipping true "shipping data"
// @Success 200 {object} helper.Response "success"
// @Success 401 {object} helper.Response "Unauthorized"
// @Success 500 {object} helper.Response "server error"
// @Router /shipping [POST]
func (voucherHadler *VoucherHadler) Create(c *gin.Context) {
}

// @summary API Get all shipping
// @description get all shipping
// @tags shipping
// @Produce json
// @Success 200 {object} helper.Response{data=[]model.Shipping} "success"
// @Success 401 {object} helper.Response "Unauthorized"
// @Success 500 {object} helper.Response "server error"
// @Security ApiKeyAuthToken
// @Security ApiKeyAuthIDKey
// @Router /shipping [get]
func (voucherHadler *VoucherHadler) GetAll(c *gin.Context) {

}
