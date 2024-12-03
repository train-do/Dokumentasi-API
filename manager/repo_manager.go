package manager

import "github.com/project-sistem-voucher/api/repository"

type RepoManager interface {
	VoucherRepo() repository.VoucherRepository
	RedeemRepo() repository.RedeemRepository
	RepoApplicationVoucher() repository.RepoApplicationVoucher
	RepoHistoryVoucher() repository.RepoHistoryVoucher
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) VoucherRepo() repository.VoucherRepository {
	return repository.NewVoucherRepository(m.infraManager.Conn())
}

func (m *repoManager) RedeemRepo() repository.RedeemRepository {
	return repository.NewRedeemRepository(m.infraManager.Conn())
}

func (m *repoManager) RepoApplicationVoucher() repository.RepoApplicationVoucher {
	return repository.NewRepoApplicationVoucher(m.infraManager.Conn())
}

func (m *repoManager) RepoHistoryVoucher() repository.RepoHistoryVoucher {
	return repository.NewRepoHistoryVoucher(m.infraManager.Conn())
}
