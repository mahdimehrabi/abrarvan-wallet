package repository

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
)

type WalletRepository struct {
	logger   interfaces.Logger
	db       interfaces.DB
	memoryDb interfaces.MemoryDB
}

func NewWalletRepository(
	logger infrastracture.ArvanLogger,
	db infrastracture.PgxDB,
	memoryDB infrastracture.Redis) WalletRepository {
	return WalletRepository{
		db:       &db,
		logger:   &logger,
		memoryDb: &memoryDB,
	}
}
