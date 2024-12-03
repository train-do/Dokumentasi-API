package model

import "time"

type Voucher struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	NamaVoucher      string    `json:"nama_voucher" binding:"required"`
	KodeVoucher      string    `json:"kode_voucher" gorm:"type:varchar(100);unique;not null;"`
	TipeVoucher      string    `json:"tipe_voucher" binding:"required,oneof=e-commerce redeem-poin"`
	Deskripsi        string    `json:"deskripsi" binding:"required"`
	JenisVoucher     string    `json:"jenis_voucher" binding:"required,oneof=gratis-ongkir diskon"`
	Ketentuan        string    `json:"ketentuan" binding:"required"`
	MinimanBelanja   float64   `json:"minimal_belanja,omitempty"`
	MetodePembayaran string    `json:"metode_pembayaran" binding:"required"`
	PersentaseDiskon *float64  `json:"persentase_diskon,omitempty"`
	NominalDiskon    *float64  `json:"nominal_diskon,omitempty"`
	MulaiBerlaku     time.Time `json:"mulai_berlaku" binding:"required"`
	BerakhirBerlaku  time.Time `json:"berakhir_berlaku" binding:"required"`
	AreaBerlaku      string    `json:"area_berlaku" binding:"required"`
	Point            int       `json:"point,omitempty"`
	Kuota            int       `json:"kuota,omitempty"`
	NilaiTukarPoin   int       `json:"nilai_tukar_poin,omitempty"`
	Redeem           *Redeem   `json:"reedem,omitempty" gorm:"foreignKey:KodeVoucher;references:KodeVoucher"`
	Use              *Use      `json:"use,omitempty" gorm:"foreignKey:VoucherCode;references:KodeVoucher"`
}
