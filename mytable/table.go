package mytable

import (
	"database/sql"
	"sort"
	"strings"

	"github.com/hongshengjie/crud/snaker"
)

// Table Table
type Table struct {
	TableName       string    // table name
	GoTableName     string    // go struct name
	PackageName     string    // package name
	Fields          []*Column // columns
	Indexes         []*Index  // indexes
	IndexesCol      []*Column // indexes column 用于生成where.go
	WhereImportTime bool      // is need import time where.go
	PrimaryKey      *Column   // priomary_key column
	ImportTime      bool      // is need import time

}

// NewTable NewTable
func NewTable(db *sql.DB, schema, table string) *Table {
	gotableName := snaker.SnakeToCamelIdentifier(table)
	mytable := &Table{
		TableName:   table,
		GoTableName: gotableName,
		PackageName: strings.ToLower(gotableName),
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
	indexcols := make(map[string]*Column)
	for _, v := range indexes {
		for _, fieldName := range v.IndexFields {
			for _, c := range columns {
				if c.ColumnName == fieldName {
					v.IndexColumns = append(v.IndexColumns, c)
					indexcols[c.ColumnName] = c
					break
				}
			}

		}
	}
	indexcolarr := make([]*Column, 0, len(indexcols))
	for _, v := range indexcols {
		if v.GoColumnType == "time.Time" {
			mytable.WhereImportTime = true
		}
		indexcolarr = append(indexcolarr, v)
	}
	sort.Slice(indexcolarr, func(i, j int) bool {
		return indexcolarr[i].OrdinalPosition < indexcolarr[j].OrdinalPosition
	})
	mytable.IndexesCol = indexcolarr

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
