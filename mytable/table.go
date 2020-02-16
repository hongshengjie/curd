package mytable

import (
	"database/sql"

	"github.com/knq/snaker"
)

// Table Table
type Table struct {
	TableName   string    // table name
	GoTableName string    // go struct name
	Fields      []*Column // columns
	Indexes     []*Index  // indexes
	PrimaryKey  *Column   // priomary_key column
	ImportTime  bool      // is need import time
}

// NewTable NewTable
func NewTable(db *sql.DB, schema, table string) *Table {
	mytable := &Table{
		TableName:   table,
		GoTableName: snaker.SnakeToCamelIdentifier(table),
	}
	columns, err := MyTableColumns(db, schema, table)
	if err != nil {
		panic(err)
	}
	if len(columns) <= 0 {
		panic("schema or table not exist")
	}

	indexes, err := MyTableIndexes(db, schema, table)
	if err != nil {
		panic(err)
	}

	for _, v := range indexes {
		for _, fieldName := range v.IndexFields {
			for _, c := range columns {
				if c.ColumnName == fieldName {
					v.IndexColumns = append(v.IndexColumns, c)
					break
				}
			}

		}

	}
	mytable.Indexes = indexes
	mytable.Fields = columns
	for _, v := range columns {
		if v.IsPrimaryKey {
			mytable.PrimaryKey = v
			break
		}
	}
	if mytable.PrimaryKey == nil {
		panic("table do not have a primary key")
	}
	for _, v := range columns {
		if v.GoColumnType == "time.Time" {
			mytable.ImportTime = true
			break
		}
	}
	return mytable

}
