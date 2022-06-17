package mocks

import "context"

type DB struct {
	MockExecFn func(context.Context,
		string,
		[]interface{}) (int64, error)

	MockQueryRowFn func(context.Context,
		string,
		[]interface{},
		...interface{}) error

	MockQueryFn func(context.Context,
		string,
		[]interface{},
	) ([][]interface{}, error)
}

func (db *DB) Exec(ctx context.Context,
	query string,
	parameters []interface{}) (int64, error) {
	return db.MockExecFn(ctx, query, parameters)
}

func (db *DB) QueryRow(ctx context.Context,
	query string,
	parameters []interface{},
	scans ...interface{}) error {
	return db.MockQueryRowFn(ctx, query, parameters, scans)
}

func (db *DB) Query(ctx context.Context,
	query string,
	parameters []interface{}) ([][]interface{}, error) {
	return db.MockQueryFn(ctx, query, parameters)
}

func NewDB() *DB {
	return &DB{
		MockExecFn: func(ctx context.Context, s string, i []interface{}) (int64, error) {
			return 1, nil
		},
		MockQueryFn: func(ctx context.Context, s string, i []interface{}) ([][]interface{}, error) {
			return [][]interface{}{}, nil
		},
		MockQueryRowFn: func(ctx context.Context, s string, i []interface{}, i2 ...interface{}) error {
			return nil
		},
	}
}
