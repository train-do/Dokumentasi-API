package repository

import (
	"fmt"
	"strings"

	"github.com/project-sistem-voucher/api/model"
	"gorm.io/gorm"
)

type RepoApplicationVoucher interface {
	FindAll(userID int, voucherType string) ([]model.Voucher, error)
	FindValidVoucher(input model.InputApplyVoucher) (model.OutputApplyVoucher, error)
	Insert(use *model.Use) error
}

type repoApplicationVoucher struct {
	DB *gorm.DB
}

func NewRepoApplicationVoucher(db *gorm.DB) RepoApplicationVoucher {
	return &repoApplicationVoucher{DB: db}
}

func (r *repoApplicationVoucher) FindAll(userID int, voucherType string) ([]model.Voucher, error) {
	var vouchers []model.Voucher
	// if err := r.DB.Preload("Reedem", "user_id = ?", userID).Find(&vouchers, "tipe_voucher ilike ?", voucherType).Error; err != nil {
	// 	fmt.Printf("failed to fetch vouchers: %v RepoApplicationVoucher\n", err)
	// 	return []model.Voucher{}, err
	// }
	query := r.DB.
		// Select("*").
		Joins("JOIN redeems ON redeems.kode_voucher = vouchers.kode_voucher").
		Where("redeems.user_id = ? AND vouchers.tipe_voucher ilike ?", userID, voucherType).
		Preload("Redeem")
	err := query.Find(&vouchers).Error
	if err != nil {
		fmt.Printf("failed to fetch vouchers: %v RepoApplicationVoucher\n", err)
		return []model.Voucher{}, err
	}
	return vouchers, nil
}
func (r *repoApplicationVoucher) FindValidVoucher(input model.InputApplyVoucher) (model.OutputApplyVoucher, error) {
	output := model.OutputApplyVoucher{
		IsValid:       false,
		NominalDiskon: 0,
		Message:       "",
	}
	var voucher model.Voucher
	if err := r.DB.Find(&voucher, "kode_voucher ilike ?", input.KodeVoucher).Error; err != nil {
		fmt.Println("ERROR FIND VOUCHER VALIDATE :", err)
		output.Message = "Voucher Code invalid"
		return output, err
	}
	fmt.Println(voucher.MinimanBelanja, input.NominalTransaction)
	if voucher.MinimanBelanja > input.NominalTransaction {
		output.Message = "Minimal Transaction not reached"
		return output, fmt.Errorf(" Minimal Transaction not reached")
	}
	fmt.Println(voucher.AreaBerlaku, input.Area)
	if !strings.EqualFold(voucher.AreaBerlaku, input.Area) {
		output.Message = "Area is invalid"
		return output, fmt.Errorf(" Area is invalid")
	}
	fmt.Println(voucher.MetodePembayaran, input.MethodPayment)
	if !strings.EqualFold(voucher.MetodePembayaran, input.MethodPayment) {
		output.Message = "Method Payment is invalid"
		return output, fmt.Errorf(" Method Payment is invalid")
	}
	fmt.Println(voucher.Kuota)
	if voucher.Kuota <= 0 {
		output.Message = "Quota has run out"
		return output, fmt.Errorf(" Quota has run out")
	}
	output.IsValid = true
	if voucher.PersentaseDiskon != nil {
		output.NominalDiskon = input.NominalTransaction * *voucher.PersentaseDiskon
	}
	if voucher.NominalDiskon != nil {
		output.NominalDiskon = *voucher.NominalDiskon
	}
	output.Message = "Code Voucher Valid"
	return output, nil
}
func (r *repoApplicationVoucher) Insert(use *model.Use) error {
	var voucher model.Voucher
	if err := r.DB.Find(&voucher, "kode_voucher ilike ?", use.VoucherCode).Error; err != nil {
		fmt.Println("ERROR FIND VOUCHER INSERT :", err)
		return err
	}
	fmt.Printf("%+v\n", voucher)
	if voucher.PersentaseDiskon != nil {
		use.DiscountValue = use.NominalTransaction * *voucher.PersentaseDiskon
	}
	insertUse := r.DB.Create(&use)
	if insertUse.Error != nil {
		fmt.Println("Error inserting data:", insertUse.Error)
	} else {
		fmt.Println("Data inserted successfully, User ID:", use.ID)
	}
	updateKuota := r.DB.Model(&model.Voucher{}).Where("kode_voucher = ?", use.VoucherCode).Update("kuota", gorm.Expr("kuota - ?", 1))
	if updateKuota.Error != nil {
		fmt.Println("Error updating data:", updateKuota.Error)
	} else {
		fmt.Println("Update successful. Rows affected:", updateKuota.RowsAffected)
	}
	return nil
}
