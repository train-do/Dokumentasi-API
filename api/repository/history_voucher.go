package repository

import (
	"fmt"

	"github.com/project-sistem-voucher/api/model"
	"gorm.io/gorm"
)

type RepoHistoryVoucher interface {
	FindHistoryReedemByUserId(userID int) ([]model.Voucher, error)
	FindHistoryUseByUserId(userID int) ([]model.Voucher, error)
	FindHistoryByVoucherCode(voucher_code string) ([]model.Voucher, error)
}

type repoHistoryVoucher struct {
	DB *gorm.DB
}

func NewRepoHistoryVoucher(db *gorm.DB) RepoHistoryVoucher {
	return &repoHistoryVoucher{DB: db}
}
func (r *repoHistoryVoucher) FindHistoryReedemByUserId(userID int) ([]model.Voucher, error) {
	fmt.Println("MASUK HISTORY REEDEM")
	var vouchers []model.Voucher
	query := r.DB.
		// Select("*").
		Joins("JOIN redeems ON redeems.kode_voucher = vouchers.kode_voucher").
		Where("redeems.user_id = ?", userID).
		Preload("Redeem")
	err := query.Find(&vouchers).Error
	if err != nil {
		fmt.Printf("failed to fetch vouchers: %v RepoApplicationVoucher\n", err)
		return []model.Voucher{}, err
	}
	return vouchers, nil
}
func (r *repoHistoryVoucher) FindHistoryUseByUserId(userID int) ([]model.Voucher, error) {
	fmt.Println("MASUK HISTORY USE")
	var vouchers []model.Voucher
	query := r.DB.
		// Select("*").
		Joins("JOIN uses ON uses.voucher_code = vouchers.kode_voucher").
		Where("uses.user_id = ?", userID).
		Preload("Use")
	err := query.Find(&vouchers).Error
	if err != nil {
		fmt.Printf("failed to fetch vouchers: %v RepoApplicationVoucher\n", err)
		return []model.Voucher{}, err
	}
	return vouchers, nil
}
func (r *repoHistoryVoucher) FindHistoryByVoucherCode(voucher_code string) ([]model.Voucher, error) {
	fmt.Println("MASUK HISTORY ALL")
	var vouchers []model.Voucher
	query := r.DB.
		// Select("*").
		Joins("JOIN uses ON uses.voucher_code = vouchers.kode_voucher").
		Where("uses.voucher_code ilike ?", voucher_code).
		Preload("Use")
	err := query.Find(&vouchers).Error
	if err != nil {
		fmt.Printf("failed to fetch vouchers: %v RepoApplicationVoucher\n", err)
		return []model.Voucher{}, err
	}
	return vouchers, nil
}
