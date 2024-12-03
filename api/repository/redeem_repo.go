package repository

import (
	"github.com/project-sistem-voucher/api/model"
	"gorm.io/gorm"
)

type RedeemRepository interface {
	SaveRedeem(redeem *model.Redeem) error
}

type redeemRepository struct {
	db *gorm.DB
}

func NewRedeemRepository(db *gorm.DB) RedeemRepository {
	return &redeemRepository{db: db}
}

func (r *redeemRepository) SaveRedeem(redeem *model.Redeem) error {
	return r.db.Create(redeem).Error
}
