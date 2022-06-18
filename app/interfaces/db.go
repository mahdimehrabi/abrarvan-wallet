package interfaces

import (
	"context"
)

type Model interface {
	SliceToModel(data []interface{}) error
}

type DB interface {
	//execute sql command and return rows affected count and err
	Exec(
		ctx context.Context,
		query string,
		parameters []interface{}) (int64, error)

	//get signle row
	QueryRow(
		ctx context.Context,
		query string,
		parameters []interface{},
		scans ...interface{}) error

	//get multiple rows
	Query(ctx context.Context,
		query string,
		parameters []interface{},
	) (slc [][]interface{}, err error)
	//get a transaction
	Begin(ctx context.Context) (Transaction, error)
}

type Transaction interface {
	Rollback(ctx context.Context) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (int64, error)
	Commit(ctx context.Context) error
}
