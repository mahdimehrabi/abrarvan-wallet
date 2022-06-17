package repository

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/models"
	"context"
)

type WalletRepository struct {
	Logger   interfaces.Logger
	DB       interfaces.DB
	MemoryDB interfaces.MemoryDB
}

func (r WalletRepository) CreateUser(user *models.User) error {
	parameters := []interface{}{
		user.Mobile, user.Credit, user.ReceivedCharge,
	}
	_, err := r.DB.Exec(context.TODO(),
		"INSERT INTO USERS(mobile,credit,received_charge) values($1,$2,$3)",
		parameters,
	)
	return err
}

func NewWalletRepository(
	logger infrastracture.ArvanLogger,
	db infrastracture.PgxDB,
	memoryDB infrastracture.Redis) WalletRepository {
	return WalletRepository{
		DB:       &db,
		Logger:   &logger,
		MemoryDB: &memoryDB,
	}
}
