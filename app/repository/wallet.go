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

func (r WalletRepository) TryDecreaseConsumerCount(code string, mobile string) error {

	//check code availability
	codeModel, err := r.GetCreateCodeMemoryDB(code)
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
	rowsAffected, err := tx.Exec(ctx,
		"UPDATE users SET credit=credit+$1,received_charge=TRUE WHERE mobile=$2 AND received_charge=FALSE",
		codeModel.Credit, mobile)
	if rowsAffected != 1 || err != nil {
		return errors.New("can't find proper user with this mobile")
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

func (r WalletRepository) GetCreateCodeMemoryDB(code string) (codeModel models.Code, err error) {
	codeJson, err := r.MemoryDB.Get("code_" + code)
	if err == nil && code != "" {
		err = codeModel.FromJSON(strings.NewReader(codeJson))
		if err != nil {
			return
		}
		return
	}

	codeModel, err = r.getCodeFromDB(code)
	if err != nil {
		return
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

func (r WalletRepository) getCodeFromDB(code string) (models.Code, error) {
	var credit float64
	var consumerCount int64
	parameters := []interface{}{code}
	err := r.DB.QueryRow(context.TODO(),
		"SELECT * FROM codes where code=$1",
		parameters, &code, &credit, &consumerCount)
	if err != nil {
		return models.Code{}, err
	}
	if code == "" {
		err = errors.New("Code not found!")
		return models.Code{}, err
	}
	codeModel := models.Code{
		Code:          code,
		Credit:        credit,
		ConsumerCount: int(consumerCount),
	}
	return codeModel, nil
}

func (r WalletRepository) UserBalance(user *models.User) error {
	err := r.DB.QueryRow(context.TODO(), "SELECT * FROM users WHERE mobile=$1",
		[]interface{}{user.Mobile},
		&user.Mobile, &user.Credit, &user.ReceivedCharge)
	if err != nil {
		return err
	}
	return nil
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
