package repository

import (
	"time"

	"github.com/project-sistem-voucher/api/model"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	CreateVoucher(voucher *model.Voucher) error
	FindByKodeVoucher(kode string) (*model.Voucher, error)
	DeleteVoucherByID(voucherID uint) error
	FindByID(voucherID uint) (*model.Voucher, error)
	UpdateVoucher(voucherID uint, updatedVoucher *model.Voucher) error
	GetVouchers(params map[string]string) ([]model.Voucher, error)
	GetVouchersForRedeem(userPoints int, vouchers *[]model.Voucher) error
	GetVoucherByKode(kode string, voucher *model.Voucher) error
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db: db}
}

func (r *voucherRepository) CreateVoucher(voucher *model.Voucher) error {
	return r.db.Create(voucher).Error
}

func (r *voucherRepository) FindByKodeVoucher(kode string) (*model.Voucher, error) {
	var voucher model.Voucher
	err := r.db.Where("kode_voucher = ?", kode).First(&voucher).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &voucher, nil
}

func (r *voucherRepository) DeleteVoucherByID(voucherID uint) error {
	return r.db.Delete(&model.Voucher{}, voucherID).Error
}

func (r *voucherRepository) FindByID(voucherID uint) (*model.Voucher, error) {
	var voucher model.Voucher
	err := r.db.First(&voucher, voucherID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &voucher, nil
}

func (r *voucherRepository) UpdateVoucher(voucherID uint, updatedVoucher *model.Voucher) error {
	result := r.db.Model(&model.Voucher{}).Where("id = ?", voucherID).Updates(updatedVoucher)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *voucherRepository) GetVouchers(params map[string]string) ([]model.Voucher, error) {
	var vouchers []model.Voucher
	query := r.db

	if tipeVoucher, ok := params["tipe_voucher"]; ok && tipeVoucher != "" {
		query = query.Where("tipe_voucher = ?", tipeVoucher)
	}

	if status, ok := params["status"]; ok && status != "" {
		now := time.Now()
		if status == "aktif" {
			query = query.Where("mulai_berlaku <= ? AND berakhir_berlaku >= ?", now, now)
		} else if status == "non-aktif" {
			query = query.Where("berakhir_berlaku < ? OR mulai_berlaku > ?", now, now)
		}
	}

	if area, ok := params["area"]; ok && area != "" {
		query = query.Where("area_berlaku LIKE ?", "%"+area+"%")
	}

	if metodePembayaran, ok := params["metode_pembayaran"]; ok && metodePembayaran != "" {
		query = query.Where("metode_pembayaran = ?", metodePembayaran)
	}

	err := query.Find(&vouchers).Error
	if err != nil {
		return nil, err
	}

	return vouchers, nil
}

func (r *voucherRepository) GetVouchersForRedeem(userPoints int, vouchers *[]model.Voucher) error {
	return r.db.Where("nilai_tukar_poin <= ?", userPoints).Find(vouchers).Error
}

func (r *voucherRepository) GetVoucherByKode(kode string, voucher *model.Voucher) error {
	return r.db.Where("kode_voucher = ?", kode).First(voucher).Error
}
