package model

import "time"

type OutputApplyVoucher struct {
	IsValid       bool    `json:"isValid"`
	NominalDiskon float64 `json:"nominalDiscount"`
	Message       string  `json:"message,omitempty"`
}
type InputApplyVoucher struct {
	UserID             uint    `binding:"required"`
	KodeVoucher        string  `binding:"required"`
	NominalTransaction float64 `binding:"required"`
	MethodPayment      string  `binding:"required"`
	Area               string  `binding:"required"`
}
type Use struct {
	ID                 uint      `json:"id,omitempty"`
	UserID             uint      `json:"userId,omitempty"`
	VoucherCode        string    `json:"voucherCode,omitempty"`
	NominalTransaction float64   `json:"nominalTransaction,omitempty"`
	DiscountValue      float64   `json:"discountValue,omitempty"`
	CreatedAt          time.Time `json:"useByDate,omitempty"`
	// Voucher            Voucher   `json:"omitempty" gorm:"foreignKey:VoucherCode;references:KodeVoucher"`
}
