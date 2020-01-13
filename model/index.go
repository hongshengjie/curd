package model

import (
	"database/sql"
)

// Index represents an index.
type Index struct {
	IndexName    string         // index_name
	IndexGoName  string         // indexGoName
	IsUnique     bool           // is_unique
	IsPrimary    bool           // is_primary
	SeqNo        int            // seq_no
	Origin       string         // origin
	IsPartial    bool           // is_partial
	IndexColumns []*IndexColumn //indexes columns
}

// MyTableIndexes runs a custom query, returning results as Index.
func MyTableIndexes(db *sql.DB, schema string, table string) ([]*Index, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`DISTINCT index_name, ` +
		`NOT non_unique AS is_unique ` +
		`FROM information_schema.statistics ` +
		`WHERE index_name <> 'PRIMARY' AND index_schema = ? AND table_name = ?`

	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Index{}
	for q.Next() {
		i := Index{}

		// scan
		err = q.Scan(&i.IndexName, &i.IsUnique)
		if err != nil {
			return nil, err
		}
		if i.IndexName == "ix_mtime" {
			continue
		}

		res = append(res, &i)
	}

	return res, nil
}
