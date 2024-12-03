package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/service"
	"gorm.io/gorm"
)

type VoucherHandler struct {
	Service service.VoucherService
}

func NewVoucherHandler(service service.VoucherService) *VoucherHandler {
	return &VoucherHandler{Service: service}
}

// Get voucher all endpoint
// @Summary Get Voucher By Category
// @Description Get Voucher By Category.
// @Tags apply-voucher
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Voucher
// @Failure 404 {object} model.Voucher
// @Failure 500 {object} model.Voucher
// @Router  /api/v1/apply-voucher/use [get]
func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
	var input model.Voucher
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk membuat voucher
	voucher, err := h.Service.CreateVoucher(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Voucher created successfully", "data": voucher})
}

func (h *VoucherHandler) DeleteVoucher(c *gin.Context) {
	idParam := c.Param("id")
	voucherID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voucher_id tidak valid"})
		return
	}

	err = h.Service.DeleteVoucherByID(uint(voucherID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voucher berhasil dihapus"})
}

func (h *VoucherHandler) UpdateVoucher(c *gin.Context) {
	idParam := c.Param("id")
	voucherID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voucher_id tidak valid"})
		return
	}

	var updatedVoucher model.Voucher
	if err := c.ShouldBindJSON(&updatedVoucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Service.UpdateVoucher(uint(voucherID), &updatedVoucher)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "voucher tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voucher berhasil diupdate"})
}

func (h *VoucherHandler) GetVouchers(c *gin.Context) {
	params := map[string]string{
		"tipe_voucher":      c.Query("tipe_voucher"),
		"status":            c.Query("status"),
		"area":              c.Query("area"),
		"metode_pembayaran": c.Query("metode_pembayaran"),
	}

	vouchers, err := h.Service.GetVouchers(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vouchers})
}

func (h *VoucherHandler) GetVouchersForRedeem(c *gin.Context) {
	userIDStr := c.DefaultQuery("user_id", "")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	userPoints := getUserPoints(userID)

	vouchers, err := h.Service.GetVouchersForRedeem(userPoints)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch vouchers"})
		return
	}

	c.JSON(http.StatusOK, vouchers)
}

func getUserPoints(userID int) int {
	return 100
}
