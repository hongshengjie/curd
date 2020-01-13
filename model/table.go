package model

import (
	"database/sql"

	"github.com/knq/snaker"
)

// Table Table
type Table struct {
	Name       string    // table name
	GoName     string    // go struct name
	Fields     []*Column // columns
	Indexes    []*Index  // indexes
	PrimaryKey *Column   // priomary_key column
}

// NewTable NewTable
func NewTable(db *sql.DB, schema, table string) *Table {
	mytable := &Table{
		Name:   table,
		GoName: snaker.SnakeToCamel(table),
	}
	columns, err := MyTableColumns(db, schema, table)
	if err != nil {
		panic(err)
	}

	indexes, err := MyTableIndexes(db, schema, table)
	if err != nil {
		panic(err)
	}

	for _, v := range indexes {
		indexColumns, err := MyIndexColumns(db, schema, table, v.IndexName)
		if err != nil {
			panic(err)
		}
		var indexGoName string
		for _, d := range indexColumns {
			for _, c := range columns {
				if c.ColumnName == d.IndexColumnName {
					d.Column = c
					break
				}
			}
			indexGoName = indexGoName + d.GoName
		}
		v.IndexColumns = indexColumns
		v.IndexGoName = indexGoName

	}
	mytable.Indexes = indexes
	mytable.Fields = columns
	for _, v := range columns {
		if v.IsPrimaryKey {
			mytable.PrimaryKey = v
			break
		}
	}
	return mytable

}
