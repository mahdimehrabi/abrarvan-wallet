package infrastracture

import (
	"challange/app/interfaces"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Transaction struct {
	Tx pgx.Tx
}

func (t Transaction) Rollback(ctx context.Context) error {
	return t.Tx.Rollback(ctx)
}

func (t Transaction) Exec(ctx context.Context, sql string, arguments ...interface{}) (int64, error) {
	cm, err := t.Tx.Exec(ctx, sql, arguments...)
	return cm.RowsAffected(), err
}

func (t Transaction) Commit(ctx context.Context) error {
	return t.Tx.Commit(ctx)
}

type PgxDB struct {
	logger interfaces.Logger
	Conn   *pgx.Conn
}

func NewPgxDB(logger ArvanLogger) PgxDB {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DBUsername"), os.Getenv("DBPassword"),
		os.Getenv("DBHost"), os.Getenv("DBPort"),
		os.Getenv("DBName"),
	)
	conn, err := pgx.Connect(context.TODO(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return PgxDB{
		logger: &logger,
		Conn:   conn,
	}
}

//Exec => execute sql command and return rows affected count and err
func (db *PgxDB) Exec(
	ctx context.Context,
	query string,
	parameters []interface{}) (int64, error) {
	cmdTag, err := db.Conn.Exec(ctx, query, parameters...)
	return cmdTag.RowsAffected(), err
}

//Query => for get multiple rows
func (db *PgxDB) Query(
	ctx context.Context,
	query string,
	parameters []interface{}) (slc [][]interface{}, err error) {
	rows, err := db.Conn.Query(ctx, query, parameters...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			//Just use error instead of writing warning method to save development time
			db.logger.Error(fmt.Sprintf("warrning:%s", err))
			continue
		}
		slc = append(slc, values)
	}
	return slc, nil
}

//QueryRow => for get signle row
func (db *PgxDB) QueryRow(
	ctx context.Context,
	query string,
	parameters []interface{},
	scans ...interface{}) error {
	return db.Conn.QueryRow(ctx, query, parameters...).Scan(scans...)
}

func (db *PgxDB) Begin(ctx context.Context) (interfaces.Transaction, error) {
	t, err := db.Conn.Begin(ctx)
	tx := Transaction{Tx: t}
	return tx, err
}
