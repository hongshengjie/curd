package model

import "database/sql"

// IndexColumn IndexColumn
type IndexColumn struct {
	SeqNo           int    // seq_no
	Cid             int    // cid
	IndexColumnName string // column_name
	*Column
}

// MyIndexColumns MyIndexColumns
func MyIndexColumns(db *sql.DB, schema string, table string, index string) ([]*IndexColumn, error) {
	var err error

	const sqlstr = `SELECT ` +
		`seq_in_index AS seq_no, ` +
		`column_name ` +
		`FROM information_schema.statistics ` +
		`WHERE index_schema = ? AND table_name = ? AND index_name = ? ` +
		`ORDER BY seq_in_index`

	q, err := db.Query(sqlstr, schema, table, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		err = q.Scan(&ic.SeqNo, &ic.IndexColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}
