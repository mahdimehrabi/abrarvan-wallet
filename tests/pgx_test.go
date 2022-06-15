package tests

import (
	"challange/app/infrastracture"
	"context"
	"testing"
	"time"
)

func TestPgx(t *testing.T) {
	l := infrastracture.NewLogger()
	db := infrastracture.NewPgxDB(l)
	ctx := context.Background()
	defer db.Conn.Close(ctx)
	rand := infrastracture.NewRandom()
	rand.RefreshSeed()
	id := rand.GenerateRandomStr(10)
	segment := rand.GenerateRandomStr(10)
	expired := time.Now().Add(time.Hour * time.Duration(rand.RandomNumber(700)))

	//----test exec method -----
	rowsAffected, err := db.Exec(
		ctx,
		"INSERT INTO USERS (ID,segment,expired_segment) values($1,$2,$3) ",
		[]interface{}{id, segment, expired},
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
	if values[len(values)-1][0] != id {
		t.Errorf(
			"pgx query values problem:%v is not equal to:%v",
			id, values[len(values)-1][0])
	}
	if values[len(values)-1][1] != segment {
		t.Errorf(
			"pgx query values problem:%v is not equal to:%v",
			segment, values[len(values)-1][1])
	}

	//----test QueryRow method -----
	var newID string
	var newSegment string
	var newExpiredSegment time.Time
	err = db.QueryRow(
		ctx,
		"SELECT * FROM users WHERE id=$1",
		[]interface{}{id},
		&newID, &newSegment, &newExpiredSegment)
	if err != nil {
		t.Errorf("Pgx query row problem: %s", err.Error())
	}
	if newID != id {
		t.Errorf(
			"pgx query row: %v is not equal to:%v",
			id, newID)
	}
	if newSegment != segment {
		t.Errorf(
			"pgx query row: %v is not equal to:%v",
			segment, newSegment)
	}
	if newExpiredSegment.Equal(expired) {
		t.Errorf(
			"pgx query row: %v is not equal to:%v",
			newExpiredSegment, expired)
	}
}
