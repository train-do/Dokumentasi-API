package service_test

import (
	"testing"
	"time"

	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVoucherRepository struct {
	mock.Mock
}

func (*MockVoucherRepository) GetVoucherByKode(kode string, voucher *model.Voucher) error {
	return nil
}

func (m *MockVoucherRepository) CreateVoucher(voucher *model.Voucher) error {
	args := m.Called(voucher)
	return args.Error(0)
}

func (m *MockVoucherRepository) FindByKodeVoucher(kode string) (*model.Voucher, error) {
	args := m.Called(kode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Voucher), args.Error(1)
}

func (m *MockVoucherRepository) FindByID(id uint) (*model.Voucher, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Voucher), args.Error(1)
}

func (m *MockVoucherRepository) DeleteVoucherByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockVoucherRepository) UpdateVoucher(id uint, voucher *model.Voucher) error {
	args := m.Called(id, voucher)
	return args.Error(0)
}

func (m *MockVoucherRepository) GetVouchers(params map[string]string) ([]model.Voucher, error) {
	args := m.Called(params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Voucher), args.Error(1)
}

func (m *MockVoucherRepository) GetVouchersForRedeem(userPoints int, vouchers *[]model.Voucher) error {
	args := m.Called(userPoints, vouchers)
	return args.Error(0)
}

func TestCreateVoucher(t *testing.T) {
	mockRepo := new(MockVoucherRepository)
	service := service.NewVoucherService(mockRepo)

	tests := []struct {
		name          string
		input         model.Voucher
		mockBehavior  func()
		expectedError string
	}{
		{
			name: "success create voucher",
			input: model.Voucher{
				KodeVoucher:     "DISKON50",
				NamaVoucher:     "Diskon 50%",
				MulaiBerlaku:    time.Now(),
				BerakhirBerlaku: time.Now().AddDate(0, 0, 7),
			},
			mockBehavior: func() {
				mockRepo.On("FindByKodeVoucher", "DISKON50").Return(nil, nil)
				mockRepo.On("CreateVoucher", mock.Anything).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "error duplicate voucher code",
			input: model.Voucher{
				KodeVoucher:     "DISKON50",
				NamaVoucher:     "Diskon 50%",
				MulaiBerlaku:    time.Now(),
				BerakhirBerlaku: time.Now().AddDate(0, 0, 7),
			},
			mockBehavior: func() {
				mockRepo.On("FindByKodeVoucher", "DISKON50").Return(&model.Voucher{}, nil)
			},
			expectedError: "kode voucher sudah digunakan",
		},
		{
			name: "error invalid date",
			input: model.Voucher{
				KodeVoucher:     "DISKON50",
				NamaVoucher:     "Diskon 50%",
				MulaiBerlaku:    time.Now(),
				BerakhirBerlaku: time.Now().AddDate(0, -1, 0),
			},
			mockBehavior:  func() {},
			expectedError: "tanggal kadaluarsa tidak boleh sebelum tanggal mulai",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			tt.mockBehavior()

			voucher, err := service.CreateVoucher(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, voucher)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteVoucherByID(t *testing.T) {
	mockRepo := new(MockVoucherRepository)
	service := service.NewVoucherService(mockRepo)

	tests := []struct {
		name          string
		voucherID     uint
		mockBehavior  func()
		expectedError string
	}{
		{
			name:      "success delete voucher",
			voucherID: 1,
			mockBehavior: func() {
				mockRepo.On("FindByID", uint(1)).Return(&model.Voucher{}, nil)
				mockRepo.On("DeleteVoucherByID", uint(1)).Return(nil)
			},
			expectedError: "",
		},
		{
			name:      "voucher not found",
			voucherID: 2,
			mockBehavior: func() {
				mockRepo.On("FindByID", uint(2)).Return(nil, nil)
			},
			expectedError: "voucher tidak ditemukan",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			tt.mockBehavior()

			err := service.DeleteVoucherByID(tt.voucherID)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
