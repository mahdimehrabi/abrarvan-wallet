package tasks

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"context"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeWalletCodesUpdate = "update:WalletCode"
)

type WalletTask struct {
	logger   interfaces.Logger
	memoryDB interfaces.MemoryDB
	db       interfaces.DB
}

func NewWalletTask(
	logger infrastracture.ArvanLogger,
	redis infrastracture.Redis,
	db infrastracture.PgxDB) WalletTask {
	return WalletTask{
		logger:   &logger,
		memoryDB: &redis,
		db:       &db,
	}
}

func (et *WalletTask) NewUpdateCodesWalletTask() (*asynq.Task, error) {
	return asynq.NewTask(
		TypeWalletCodesUpdate,
		[]byte{},
		asynq.Timeout(80*time.Second),
		asynq.MaxRetry(2)), nil
}

//this method get count of Wallets users and Store count of Wallet users in memory db
func (et WalletTask) HandleWalletCodesUpdateTask(ctx context.Context, t *asynq.Task) error {
	return nil
}
