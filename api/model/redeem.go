package model

import "time"

type Redeem struct {
	RedeemID      uint      `gorm:"primaryKey" json:"redeem_id,omitempty"`
	UserID        uint      `json:"user_id,omitempty" binding:"required"`
	KodeVoucher   string    `json:"kode_voucher,omitempty" binding:"required"`
	TanggalRedeem time.Time `json:"tanggal_redeem,omitempty"`
}
