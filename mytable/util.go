package mytable

import (
	"strings"
)

// MysqlToGoFieldType MysqlToGoFieldType
func MysqlToGoFieldType(dt, ct string) string {
	var unsigned bool
	if strings.Contains(ct, "unsigned") {
		unsigned = true
	}
	var typ string
	switch dt {
	case "bit":
		typ = "[]uint8"
	case "bool", "boolean":
		typ = "bool"
	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext", "json":
		typ = "string"
	case "tinyint":
		typ = "int8"
		if unsigned {
			typ = "uint8"
		}
	case "smallint":
		typ = "int16"
		if unsigned {
			typ = "uint16"
		}
	case "mediumint", "int", "integer":
		typ = "int32"
		if unsigned {
			typ = "uint32"
		}
	case "bigint":
		typ = "int64"
		if unsigned {
			typ = "uint64"
		}
	case "float":
		typ = "float32"
	case "decimal", "double":
		typ = "float64"
	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		typ = "[]byte"
	case "timestamp", "datetime", "date":
		typ = "time.Time"
	case "time", "year", "enum", "set":
		typ = "string"
	default:
		typ = "UNKNOWN"
	}
	return typ
}

//SQLTool SQLTool
func SQLTool(t *Table, omit bool, flag string) string {
	var ns []string
	for _, v := range t.Fields {
		if omit {
			if v.IsPrimaryKey || v.IsDefaultCurrentTimestamp {
				continue
			}
		}
		switch flag {
		case "field":
			ns = append(ns, v.ColumnName)
		case "?":
			ns = append(ns, "?")
		case "gofield":
			ns = append(ns, "&a."+v.GoColumnName)
		case "set":
			ns = append(ns, v.ColumnName+" = ? ")
		default:
			ns = append(ns, flag)
		}

	}
	return strings.Join(ns, ",")
}

// SQLUpdate sql update
func SQLUpdate(t *Table, omit bool) string {
	var ns []string
	var prime *Column
	for _, v := range t.Fields {
		if omit {
			if v.IsDefaultCurrentTimestamp {
				continue
			}
		}
		if v.IsPrimaryKey {
			prime = v
			continue
		}
		ns = append(ns, "a."+v.GoColumnName)
	}
	ns = append(ns, "a."+prime.GoColumnName)

	return strings.Join(ns, ",")
}

// SQLIndexParamList SQLIndexParamList
func SQLIndexParamList(index *Index, needType bool) string {
	var ns []string
	for _, v := range index.IndexColumns {
		s := v.ColumnName
		if needType {
			s = s + " " + v.GoColumnType
		}
		ns = append(ns, s)
	}
	return strings.Join(ns, ",")
}

// SQLIndexQuery SQLIndexQuery
func SQLIndexQuery(index *Index) string {
	var ns []string
	for _, v := range index.IndexColumns {
		s := v.ColumnName + " = ? "

		ns = append(ns, s)
	}
	return strings.Join(ns, "AND ")
}
