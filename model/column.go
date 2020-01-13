package model

import (
	"database/sql"
	"github.com/knq/snaker"
)

// Column Column
type Column struct {
	FieldOrdinal int            // field_ordinal
	ColumnName   string         // column_name
	DataType     string         // data_type
	NotNull      bool           // not_null
	DefaultValue sql.NullString // default_value
	IsPrimaryKey bool           // is_primary_key
	GoName       string         // go field name
	GoType       string         // go field type
}

// MyTableColumns MyTableColumns
func MyTableColumns(db *sql.DB, schema string, table string) ([]*Column, error) {
	var err error

	const sqlstr = `SELECT ` +
		`ordinal_position AS field_ordinal, ` +
		`column_name, ` +
		`IF(data_type = 'enum', column_name, column_type) AS data_type, ` +
		`IF(is_nullable = 'YES', false, true) AS not_null, ` +
		`column_default AS default_value, ` +
		`IF(column_key = 'PRI', true, false) AS is_primary_key ` +
		`FROM information_schema.columns ` +
		`WHERE table_schema = ? AND table_name = ? ` +
		`ORDER BY ordinal_position`

	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*Column{}
	for q.Next() {
		c := Column{}
		err = q.Scan(&c.FieldOrdinal, &c.ColumnName, &c.DataType, &c.NotNull, &c.DefaultValue, &c.IsPrimaryKey)
		if err != nil {
			return nil, err
		}
		c.GoName = snaker.SnakeToCamel(c.ColumnName)
		c.GoType = MysqlTypeToGoType(c.DataType)
		res = append(res, &c)
	}

	return res, nil
}
