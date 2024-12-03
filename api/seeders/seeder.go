package seeders

import (
	"log"
	"time"

	"github.com/project-sistem-voucher/api/model"
	"gorm.io/gorm"
)

func SeedVouchers(db *gorm.DB) {
	persentasiDisc := 0.5
	vouchers := []model.Voucher{
		{
			NamaVoucher:      "Gratis Ongkir Jawa",
			KodeVoucher:      "ONGKIRJAWA",
			TipeVoucher:      "e-commerce",
			JenisVoucher:     "Gratis Ongkir",
			Ketentuan:        "Minimal pembelian Rp100.000 dengan metode pembayaran transfer bank.",
			MinimanBelanja:   50000.00,
			MetodePembayaran: "COD",
			Kuota:            5,
			AreaBerlaku:      "Jawa",
			MulaiBerlaku:     time.Now(),
			BerakhirBerlaku:  time.Now().AddDate(0, 1, 0),
		},
		{
			NamaVoucher:      "Diskon 50%",
			KodeVoucher:      "DISKON50",
			TipeVoucher:      "redeem poin",
			JenisVoucher:     "Diskon",
			PersentaseDiskon: &persentasiDisc,
			Ketentuan:        "Diskon 50% untuk pembelian minimal Rp200.000.",
			AreaBerlaku:      "Nasional",
			Point:            100,
			MulaiBerlaku:     time.Now(),
			BerakhirBerlaku:  time.Now().AddDate(0, 2, 0),
		},
		{
			NamaVoucher:      "Voucher Diskon",
			KodeVoucher:      "VOUCHER123",
			TipeVoucher:      "diskon",
			Deskripsi:        "Diskon 10%",
			JenisVoucher:     "produk",
			Ketentuan:        "Minimal belanja Rp100.000",
			MinimanBelanja:   100000,
			MetodePembayaran: "kartu kredit",
			PersentaseDiskon: &persentasiDisc,
			MulaiBerlaku:     time.Now(),
			BerakhirBerlaku:  time.Now().AddDate(0, 0, 30),
			AreaBerlaku:      "Indonesia",
			Point:            50,
			Kuota:            100,
			NilaiTukarPoin:   20,
		},
	}

	for _, voucher := range vouchers {
		err := db.Create(&voucher).Error
		if err != nil {
			log.Printf("Gagal membuat voucher %s: %v\n", voucher.KodeVoucher, err)
		} else {
			log.Printf("Berhasil menambahkan voucher: %s\n", voucher.KodeVoucher)
		}
	}
}

func SeedRedeem(db *gorm.DB) error {

	redeems := []model.Redeem{
		{
			UserID:        1,
			KodeVoucher:   "DISKON50",
			TanggalRedeem: time.Now(),
		},
		{
			UserID:        2,
			KodeVoucher:   "VOUCHER123",
			TanggalRedeem: time.Now(),
		},
	}

	for _, redeem := range redeems {
		db.Create(&redeem)
	}

	return nil
}
