package manager

import (
	_ "github.com/lib/pq"
	"github.com/project-sistem-voucher/config"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	cfg *config.Config
}

func (i *infraManager) Conn() *gorm.DB {
	return config.DB
}

func NewInfraManager(configParam *config.Config) InfraManager {
	infra := &infraManager{
		cfg: configParam,
	}

	return infra
}
