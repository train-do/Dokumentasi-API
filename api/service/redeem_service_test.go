package service

import (
	"errors"
	"testing"
	"time"

	"github.com/project-sistem-voucher/api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedeemRepository struct {
	mock.Mock
}

func (m *MockRedeemRepository) SaveRedeem(redeem *model.Redeem) error {
	args := m.Called(redeem)
	return args.Error(0)
}

type MockVoucherRepository struct {
	mock.Mock
}

func (*MockVoucherRepository) CreateVoucher(voucher *model.Voucher) error {
	return nil
}

func (*MockVoucherRepository) DeleteVoucherByID(voucherID uint) error {
	return nil
}

func (*MockVoucherRepository) FindByID(voucherID uint) (*model.Voucher, error) {
	return nil, nil
}

func (*MockVoucherRepository) FindByKodeVoucher(kode string) (*model.Voucher, error) {
	return nil, nil
}

func (*MockVoucherRepository) GetVouchers(params map[string]string) ([]model.Voucher, error) {
	return nil, nil
}

func (*MockVoucherRepository) GetVouchersForRedeem(userPoints int, vouchers *[]model.Voucher) error {
	return nil
}

func (*MockVoucherRepository) UpdateVoucher(voucherID uint, updatedVoucher *model.Voucher) error {
	return nil
}

func (m *MockVoucherRepository) GetVoucherByKode(kodeVoucher string, voucher *model.Voucher) error {
	args := m.Called(kodeVoucher)
	if args.Get(0) != nil {
		*voucher = args.Get(0).(model.Voucher)
	}
	return args.Error(1)
}

func TestRedeemVoucher(t *testing.T) {
	mockRedeemRepo := new(MockRedeemRepository)
	mockVoucherRepo := new(MockVoucherRepository)
	service := NewRedeemService(mockRedeemRepo, mockVoucherRepo)

	now := time.Now()
	tests := []struct {
		name          string
		userID        uint
		kodeVoucher   string
		userPoints    int
		mockBehavior  func()
		expectedError string
	}{
		{
			name:        "success redeem voucher",
			userID:      1,
			kodeVoucher: "DISC10",
			userPoints:  100,
			mockBehavior: func() {
				mockVoucherRepo.On("GetVoucherByKode", "DISC10").Return(model.Voucher{
					KodeVoucher:     "DISC10",
					NilaiTukarPoin:  50,
					MulaiBerlaku:    now.Add(-1 * time.Hour),
					BerakhirBerlaku: now.Add(1 * time.Hour),
					Kuota:           10,
				}, nil)
				mockRedeemRepo.On("SaveRedeem", mock.Anything).Return(nil)
			},
			expectedError: "",
		},
		{
			name:        "voucher not found",
			userID:      1,
			kodeVoucher: "UNKNOWN",
			userPoints:  100,
			mockBehavior: func() {
				mockVoucherRepo.On("GetVoucherByKode", "UNKNOWN").Return(model.Voucher{}, errors.New("voucher not found"))
			},
			expectedError: "voucher with code 'UNKNOWN' not found",
		},
		{
			name:        "not enough points",
			userID:      1,
			kodeVoucher: "DISC10",
			userPoints:  30,
			mockBehavior: func() {
				mockVoucherRepo.On("GetVoucherByKode", "DISC10").Return(model.Voucher{
					KodeVoucher:     "DISC10",
					NilaiTukarPoin:  50,
					MulaiBerlaku:    now.Add(-1 * time.Hour),
					BerakhirBerlaku: now.Add(1 * time.Hour),
					Kuota:           10,
				}, nil)
			},
			expectedError: "user does not have enough points to redeem voucher (required: 50, available: 30)",
		},
		{
			name:        "voucher expired",
			userID:      1,
			kodeVoucher: "DISC10",
			userPoints:  100,
			mockBehavior: func() {
				mockVoucherRepo.On("GetVoucherByKode", "DISC10").Return(model.Voucher{
					KodeVoucher:     "DISC10",
					NilaiTukarPoin:  50,
					MulaiBerlaku:    now.Add(-2 * time.Hour),
					BerakhirBerlaku: now.Add(-1 * time.Hour),
					Kuota:           10,
				}, nil)
			},
			expectedError: "voucher is not valid at this time",
		},
		{
			name:        "voucher quota exceeded",
			userID:      1,
			kodeVoucher: "DISC10",
			userPoints:  100,
			mockBehavior: func() {
				mockVoucherRepo.On("GetVoucherByKode", "DISC10").Return(model.Voucher{
					KodeVoucher:     "DISC10",
					NilaiTukarPoin:  50,
					MulaiBerlaku:    now.Add(-1 * time.Hour),
					BerakhirBerlaku: now.Add(1 * time.Hour),
					Kuota:           0,
				}, nil)
			},
			expectedError: "voucher has no remaining quota",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRedeemRepo.ExpectedCalls = nil
			mockVoucherRepo.ExpectedCalls = nil

			tt.mockBehavior()

			redeem, err := service.RedeemVoucher(tt.userID, tt.kodeVoucher, tt.userPoints)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
				assert.Empty(t, redeem)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, redeem)
				assert.Equal(t, tt.kodeVoucher, redeem.KodeVoucher)
				assert.Equal(t, tt.userID, redeem.UserID)
			}

			mockRedeemRepo.AssertExpectations(t)
			mockVoucherRepo.AssertExpectations(t)
		})
	}
}
