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
