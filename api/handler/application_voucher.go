package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/service"
)

type HandlerApplicationVoucher struct {
	Service service.ServiceApplicationVoucher
}

func NewHandlerApplicationVoucher(service service.ServiceApplicationVoucher) *HandlerApplicationVoucher {
	return &HandlerApplicationVoucher{Service: service}
}

// GetMyVoucherByCategory godoc
// @Summary Get Voucher By Category
// @Description Get Voucher By Category.
// @Tags apply-voucher
// @Produce  json
// @Success 200 {object} map[string]interface{} "OK"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router  /api/v1/apply-voucher/{userID}/{voucherType} [get]
func (h *HandlerApplicationVoucher) GetMyVoucherByCategory(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	voucherType := c.Param("voucherType")
	fmt.Println(userID, voucherType, "________")
	vouchers, err := h.Service.GetMyVoucherByCategory(userID, voucherType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get My Voucher By Category Sucess", "data": vouchers})
}
func (h *HandlerApplicationVoucher) ValidateVoucher(c *gin.Context) {
	var input model.InputApplyVoucher
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	output, err := h.Service.ValidateVoucher(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Validate Voucher Sucess", "data": output})
}
func (h *HandlerApplicationVoucher) CreateUseVoucher(c *gin.Context) {
	var use model.Use
	if err := c.ShouldBindJSON(&use); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.CreateUseVoucher(&use)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Create Use Succes", "data": use})
}
