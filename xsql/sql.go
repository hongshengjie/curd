package xsql

import (
	"context"
	"database/sql"
)

var ErrNoRows = sql.ErrNoRows

type DBTX interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// Int64 count or select only one int field
func Int64(ctx context.Context, qu Querier, dt DBTX) (int64, error) {
	sqlstr, args := qu.Query()
	q, err := dt.QueryContext(ctx, sqlstr, args...)
	if err != nil {
		return 0, err
	}
	defer q.Close()
	return ScanInt64(q)
}

// Int64 count or select only one int field
func Int64s(ctx context.Context, qu Querier, dt DBTX) ([]int64, error) {
	sqlstr, args := qu.Query()
	q, err := dt.QueryContext(ctx, sqlstr, args...)
	if err != nil {
		return nil, err
	}
	defer q.Close()
	var arr []int64
	if err = ScanSlice(q, &arr); err != nil {
		return nil, err
	}
	return arr, nil
}

// Query FindRaw many record by raw sql
func String(ctx context.Context, qu Querier, dt DBTX) (string, error) {
	sqlstr, args := qu.Query()
	q, err := dt.QueryContext(ctx, sqlstr, args...)
	if err != nil {
		return "", err
	}
	defer q.Close()
	return ScanString(q)
}

// Query FindRaw many record by raw sql
func Strings(ctx context.Context, qu Querier, dt DBTX) ([]string, error) {
	sqlstr, args := qu.Query()
	q, err := dt.QueryContext(ctx, sqlstr, args...)
	if err != nil {
		return nil, err
	}
	defer q.Close()
	var arr []string
	if err = ScanSlice(q, &arr); err != nil {
		return nil, err
	}
	return arr, nil
}
