package manager

import "github.com/project-sistem-voucher/api/service"

type ServiceManager interface {
	VoucherService() service.VoucherService
	RedeemService() service.RedeemService
	ServiceApplicationVoucher() service.ServiceApplicationVoucher
	ServiceHistoryVoucher() service.ServiceHistoryVoucher
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (m *serviceManager) VoucherService() service.VoucherService {
	return service.NewVoucherService(m.repoManager.VoucherRepo())
}

func (m *serviceManager) RedeemService() service.RedeemService {
	return service.NewRedeemService(m.repoManager.RedeemRepo(), m.repoManager.VoucherRepo())
}

func (m *serviceManager) ServiceApplicationVoucher() service.ServiceApplicationVoucher {
	return service.NewServiceApplicationVoucher(m.repoManager.RepoApplicationVoucher())
}

func (m *serviceManager) ServiceHistoryVoucher() service.ServiceHistoryVoucher {
	return service.NewServiceHistoryVoucher(m.repoManager.RepoHistoryVoucher())
}
