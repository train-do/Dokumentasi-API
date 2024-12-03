package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/service"
)

type RedeemHandler struct {
	service service.RedeemService
}

func NewRedeemHandler(service service.RedeemService) *RedeemHandler {
	return &RedeemHandler{service: service}
}

func (h *RedeemHandler) RedeemVoucher(c *gin.Context) {
	var input struct {
		UserID      uint   `json:"user_id" binding:"required"`
		KodeVoucher string `json:"kode_voucher" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userPoints := 200

	redeem, err := h.service.RedeemVoucher(input.UserID, input.KodeVoucher, userPoints)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"redeem_id":      redeem.RedeemID,
		"user_id":        redeem.UserID,
		"kode_voucher":   redeem.KodeVoucher,
		"tanggal_redeem": redeem.TanggalRedeem.Format("2006-01-02 15:04:05"),
	})
}
