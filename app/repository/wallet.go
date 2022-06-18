package repository

import (
	"bytes"
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/models"
	"context"
	"errors"
	"strings"
	"time"
)

type WalletRepository struct {
	Logger   interfaces.Logger
	DB       interfaces.DB
	MemoryDB interfaces.MemoryDB
}

func (r WalletRepository) CreateUser(user *models.User) error {
	parameters := []interface{}{
		user.Mobile, 0, false,
	}
	_, err := r.DB.Exec(context.TODO(),
		"INSERT INTO USERS(mobile,credit,received_charge) values($1,$2,$3) ON CONFLICT DO NOTHING",
		parameters,
	)
	return err
}

func (r WalletRepository) TryDecreaseConsumerCount(code string) error {

	//check code availability
	codeModel, err := r.checkCreateCodeMemoryDB(code)
	if err != nil || codeModel.ConsumerCount < 1 {
		return errors.New("this code is not available")
	}

	//create db transaction for update user
	ctx := context.Background()
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	//update users credit
	err = tx.Exec(ctx, "UPDATE users SET credit=credit+$1 ", codeModel.Credit)
	if err != nil {
		return err
	}

	//decrease consumer count
	consumerCount, err := r.MemoryDB.DecreaseConsumerCount(code)
	if err != nil {
		return err
	}
	//commit user credit to db
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	parameters := []interface{}{
		code,
	}
	if consumerCount == 0 {
		//if consumer count equal to zero update consumer count in postgres too
		rowsAffected, err := r.DB.Exec(
			context.TODO(),
			"UPDATE codes (consumer_count) values() WHERE code=$1",
			parameters,
		)
		if err != nil || rowsAffected < 1 {
			return errors.New("Failed to update codes")
		}
	}
	//consumerCount==-1 mean consumer_count is 0 and already updated in postgres
	if consumerCount == -1 {
		return errors.New("code expired")
	}

	return nil
}

func (r WalletRepository) checkCreateCodeMemoryDB(code string) (codeModel models.Code, err error) {
	code, err = r.MemoryDB.Get("code_" + code)
	if err == nil && code != "" {
		err = codeModel.FromJSON(strings.NewReader(code))
		if err != nil {
			return
		}
		return
	}

	var credit float64
	var consumerCount int64
	parameters := []interface{}{code}

	//check key exist postgres
	err = r.DB.QueryRow(context.TODO(),
		"SELECT * FROM codes where code=$1",
		parameters, code, credit, consumerCount)
	if err != nil {
		return
	}
	if code == "" {
		err = errors.New("Code not found!")
		return
	}
	codeModel = models.Code{
		Code:          code,
		Credit:        credit,
		ConsumerCount: int(consumerCount),
	}

	var buff bytes.Buffer
	err = codeModel.ToJSON(&buff)
	if err != nil {
		return
	}
	err = r.MemoryDB.Set("code_"+code, buff.String(), 24*time.Hour)
	if err != nil {
		return
	}

	return
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
