package tests

import (
	"challange/app/infrastracture"
	"challange/app/models"
	"challange/app/repository"
	"challange/app/services"
	"challange/tests/mocks"
	"testing"
)

func TestCreateUser(t *testing.T) {
	logger := infrastracture.ArvanLogger{}
	db := mocks.NewDB()
	memoryDB := mocks.NewMemoryDB()
	walletRepository := repository.WalletRepository{
		Logger:   &logger,
		DB:       db,
		MemoryDB: memoryDB,
	}
	walletService := services.WalletService{
		WalletRepository: &walletRepository,
	}
	user := models.User{
		Mobile:         "09120401761",
		Credit:         4000,
		ReceivedCharge: true,
	}
	err := walletService.CreateUser(&user)
	if err != nil {
		t.Error("error in create user:", err)
	}
}
