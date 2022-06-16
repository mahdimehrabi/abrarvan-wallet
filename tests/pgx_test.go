package tests

import (
	"challange/app/infrastracture"
	"context"
	"github.com/jackc/pgtype"
	"testing"
)

func TestPgx(t *testing.T) {
	l := infrastracture.NewLogger()
	db := infrastracture.NewPgxDB(l)
	ctx := context.Background()
	defer db.Conn.Close(ctx)
	rand := infrastracture.NewRandom()
	rand.RefreshSeed()
	mobile := rand.GenerateRandomStr(10)
	credit := rand.RandomNumber(1000000000)

	//----test exec method -----
	rowsAffected, err := db.Exec(
		ctx,
		"INSERT INTO USERS (mobile,credit) values($1,$2) ",
		[]interface{}{mobile, credit},
	)
	if err != nil {
		t.Errorf("Pgx exec: %s", err.Error())
	}
	if rowsAffected < 1 {
		t.Errorf("Pgx %d rows effected", rowsAffected)
	}

	//----test query method -----
	values, err := db.Query(ctx, "SELECT * FROM users", []interface{}{})
	if err != nil {
		t.Errorf("Pgx query problem: %s", err.Error())
	}
	if len(values) < 1 {
		t.Errorf("Pgx query len is %d", len(values))

	}
	if values[len(values)-1][0] != mobile {
		t.Errorf(
			"pgx query values problem:%v is not equal to:%v",
			mobile, values[len(values)-1][0])
	}
	pgCredit := values[len(values)-1][1].(pgtype.Numeric)
	if int(pgCredit.Int.Int64()) != credit {
		t.Errorf(
			"pgx query values problem:%v is not equal to:%v",
			credit, values[len(values)-1][1].(int))
	}

	//----test QueryRow method -----
	var newMobile string
	var newCredit int
	var recievedCharge bool
	err = db.QueryRow(
		ctx,
		"SELECT * FROM users WHERE mobile=$1",
		[]interface{}{mobile},
		&newMobile, &newCredit, &recievedCharge)
	if err != nil {
		t.Errorf("Pgx query row problem: %s", err.Error())
	}
	if newMobile != mobile {
		t.Errorf(
			"pgx query row: %v is not equal to:%v",
			mobile, newMobile)
	}
	if newCredit != credit {
		t.Errorf(
			"pgx query row: %v is not equal to:%v",
			credit, newCredit)
	}

}
