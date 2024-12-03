package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/service"
)

type HandlerHistoryVoucher struct {
	Service service.ServiceHistoryVoucher
}

func NewHandlerHistoryVoucher(service service.ServiceHistoryVoucher) *HandlerHistoryVoucher {
	return &HandlerHistoryVoucher{Service: service}
}
func (h *HandlerHistoryVoucher) GetReedemVoucherByUserId(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	vouchers, err := h.Service.GetHistoryReedemByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Get History Reedem By User Id Sucess", "data": vouchers})
}
func (h *HandlerHistoryVoucher) GetUseVoucherByUserId(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	vouchers, err := h.Service.GetHistoryUseByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Get History Use By User Id Sucess", "data": vouchers})
}
func (h *HandlerHistoryVoucher) GetAllUseByVoucherCode(c *gin.Context) {
	voucherCode := c.Param("kode_voucher")
	vouchers, err := h.Service.GetHistoryByVoucherCode(voucherCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Get History By Code Voucher Sucess", "data": vouchers})
}
