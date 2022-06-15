package tests

import (
	"challange/app/infrastracture"
	"challange/app/repository"
	"context"
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	ctx := context.TODO()
	l := infrastracture.NewLogger()
	rd := infrastracture.NewRedis()
	db := infrastracture.NewPgxDB(l)
	ur := repository.NewSegmentRepository(l, db, rd)
	rand := infrastracture.NewRandom()

	var usersCount int64
	db.QueryRow(ctx, "SELECT COUNT(*) FROM users;", []interface{}{}, &usersCount)
	id := rand.GenerateRandomStr(10)
	segment := rand.GenerateRandomStr(10)
	expired := time.Now().Add(time.Hour * time.Duration(rand.RandomNumber(700)))
	err := ur.Save(id, segment, expired)
	if err != nil {
		t.Errorf("error in saving user by repository:%s", err)
	}

	var newUserCount int64
	db.QueryRow(ctx, "SELECT COUNT(*) FROM users;", []interface{}{}, &newUserCount)
	if newUserCount != usersCount+1 {
		t.Fatalf("count problem in saving user by repository:%d not equal to %d", newUserCount, usersCount+1)
	}
}

func TestListUser(t *testing.T) {
	l := infrastracture.NewLogger()
	rd := infrastracture.NewRedis()
	db := infrastracture.NewPgxDB(l)
	ur := repository.NewSegmentRepository(l, db, rd)
	rand := infrastracture.NewRandom()

	id := rand.GenerateRandomStr(10)
	segment := rand.GenerateRandomStr(10)
	expired := time.Now().Add(time.Hour * time.Duration(rand.RandomNumber(700)))
	ur.Save(id, segment, expired)

	users, err := ur.List()

	if err != nil {
		t.Errorf("Error in listing users by repository:%s", err)
	}

	if len(users) < 1 {
		t.Fatalf("users count is:%d in users list", len(users))
	}
}
