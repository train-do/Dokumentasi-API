package service

import (
	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/repository"
)

type ServiceApplicationVoucher interface {
	GetMyVoucherByCategory(userID int, voucherType string) ([]model.Voucher, error)
	ValidateVoucher(input model.InputApplyVoucher) (model.OutputApplyVoucher, error)
	CreateUseVoucher(use *model.Use) error
}

type serviceApplicationVoucher struct {
	repo repository.RepoApplicationVoucher
}

func NewServiceApplicationVoucher(repo repository.RepoApplicationVoucher) ServiceApplicationVoucher {
	return &serviceApplicationVoucher{repo: repo}
}

func (s *serviceApplicationVoucher) GetMyVoucherByCategory(userID int, voucherType string) ([]model.Voucher, error) {
	vouchers, err := s.repo.FindAll(userID, voucherType)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
func (s *serviceApplicationVoucher) ValidateVoucher(input model.InputApplyVoucher) (model.OutputApplyVoucher, error) {
	output, err := s.repo.FindValidVoucher(input)
	if err != nil {
		return model.OutputApplyVoucher{}, err
	}
	return output, nil
}
func (s *serviceApplicationVoucher) CreateUseVoucher(use *model.Use) error {
	err := s.repo.Insert(use)
	if err != nil {
		return err
	}
	return nil
}
