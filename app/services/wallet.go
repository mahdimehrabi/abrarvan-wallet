package services

import (
	"challange/app/repository"
)

type WalletService struct {
	walletRepository *repository.WalletRepository
}

func NewWalletService(walletRepository repository.WalletRepository) WalletService {
	return WalletService{walletRepository: &walletRepository}
}

func (ss WalletService) CreateUser(jsonStr []byte) error {
	return nil
}
