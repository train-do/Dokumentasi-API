package service

import (
	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/repository"
)

type ServiceHistoryVoucher interface {
	GetHistoryReedemByUserId(userID int) ([]model.Voucher, error)
	GetHistoryUseByUserId(userID int) ([]model.Voucher, error)
	GetHistoryByVoucherCode(voucher_code string) ([]model.Voucher, error)
}

type serviceHistoryVoucher struct {
	repo repository.RepoHistoryVoucher
}

func NewServiceHistoryVoucher(repo repository.RepoHistoryVoucher) ServiceHistoryVoucher {
	return &serviceHistoryVoucher{repo: repo}
}

func (s *serviceHistoryVoucher) GetHistoryReedemByUserId(userID int) ([]model.Voucher, error) {
	vouchers, err := s.repo.FindHistoryReedemByUserId(userID)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
func (s *serviceHistoryVoucher) GetHistoryUseByUserId(userID int) ([]model.Voucher, error) {
	vouchers, err := s.repo.FindHistoryUseByUserId(userID)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
func (s *serviceHistoryVoucher) GetHistoryByVoucherCode(voucher_code string) ([]model.Voucher, error) {
	vouchers, err := s.repo.FindHistoryByVoucherCode(voucher_code)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
