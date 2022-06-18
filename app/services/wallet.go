package services

import (
	"challange/app/models"
	"challange/app/repository"
)

type WalletService struct {
	WalletRepository *repository.WalletRepository
}

func NewWalletService(walletRepository repository.WalletRepository) WalletService {
	return WalletService{WalletRepository: &walletRepository}
}

func (s WalletService) CreateUser(user *models.User) error {
	return s.WalletRepository.CreateUser(user)
}

func (s WalletService) TryDecreaseConsumerCount(code string, mobile string) error {
	return s.WalletRepository.TryDecreaseConsumerCount(code, mobile)
}

func (s WalletService) ReportCode(code string) (models.Code, error) {
	return s.WalletRepository.GetCreateCodeMemoryDB(code)
}

func (s WalletService) UserBalance(user *models.User) error {
	return s.WalletRepository.UserBalance(user)
}
