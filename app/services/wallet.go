package services

import (
	"challange/app/models"
	"challange/app/repository"
)

type WalletService struct {
	walletRepository *repository.WalletRepository
}

func NewWalletService(walletRepository repository.WalletRepository) WalletService {
	return WalletService{walletRepository: &walletRepository}
}

func (s WalletService) CreateUser(user *models.User) error {
	return s.walletRepository.CreateUser(user)
}
