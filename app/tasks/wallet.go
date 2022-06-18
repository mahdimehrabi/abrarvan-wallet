package tasks

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/models"
	"context"
	"github.com/hibiken/asynq"
	"strings"
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
	rows, err := et.db.Query(ctx, "SELECT * FROM codes WHERE consumer_count>0", []interface{}{})
	if err != nil {
		return err
	}
	for _, row := range rows {
		code := row[0].(string)
		codeJosn, err := et.memoryDB.Get("code_" + code)
		if err != nil {
			et.logger.Error(err.Error())
			continue
		}
		codeModel := models.Code{}
		err = codeModel.FromJSON(strings.NewReader(codeJosn))
		if err != nil {
			et.logger.Error(err.Error())
			continue
		}
		_, err = et.db.Exec(ctx,
			"UPDATE codes SET consumer_count=$1 WHERE code=$2",
			[]interface{}{codeModel.ConsumerCount, code})
		if err != nil {
			et.logger.Error(err.Error())
		}
	}
	return nil
}
