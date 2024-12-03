package service

import (
	"fmt"
	"time"

	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/repository"
)

type RedeemService interface {
	RedeemVoucher(userID uint, kodeVoucher string, userPoints int) (model.Redeem, error)
}

type redeemService struct {
	repo        repository.RedeemRepository
	repoVoucher repository.VoucherRepository
}

func NewRedeemService(repo repository.RedeemRepository, repoVoucher repository.VoucherRepository) RedeemService {
	return &redeemService{
		repo:        repo,
		repoVoucher: repoVoucher,
	}
}

func (s *redeemService) RedeemVoucher(userID uint, kodeVoucher string, userPoints int) (model.Redeem, error) {
	var voucher model.Voucher
	err := s.repoVoucher.GetVoucherByKode(kodeVoucher, &voucher)
	if err != nil {
		return model.Redeem{}, fmt.Errorf("voucher with code '%s' not found", kodeVoucher)
	}

	// Validasi poin pengguna
	if userPoints < voucher.NilaiTukarPoin {
		return model.Redeem{}, fmt.Errorf("user does not have enough points to redeem voucher (required: %d, available: %d)", voucher.NilaiTukarPoin, userPoints)
	}

	// Validasi tanggal voucher
	if time.Now().Before(voucher.MulaiBerlaku) || time.Now().After(voucher.BerakhirBerlaku) {
		return model.Redeem{}, fmt.Errorf("voucher is not valid at this time")
	}

	// Validasi kuota voucher
	if voucher.Kuota <= 0 {
		return model.Redeem{}, fmt.Errorf("voucher has no remaining quota")
	}

	// Simpan data redeem
	redeem := model.Redeem{
		UserID:        userID,
		KodeVoucher:   kodeVoucher,
		TanggalRedeem: time.Now(),
	}

	err = s.repo.SaveRedeem(&redeem)
	if err != nil {
		return model.Redeem{}, fmt.Errorf("failed to save redeem record: %w", err)
	}

	return redeem, nil
}
