package model

import (
	"regexp"
	"strconv"
	"strings"
)

// MysqlTypeToGoType MysqlTypeToGoType
func MysqlTypeToGoType(dt string) string {
	if strings.HasSuffix(dt, " unsigned") {
		dt = dt[:len(dt)-len(" unsigned")]
	}
	dt, _, _ = ParsePrecision(dt)

	var typ string

	switch dt {
	case "bit":
		typ = "uint64"
	case "bool", "boolean":
		typ = "bool"
	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext":
		typ = "string"
	case "tinyint":
		typ = "int8"
	case "smallint":
		typ = "int16"
	case "mediumint":
		typ = "int32"
	case "bigint", "int", "integer":
		typ = "int64"
	case "float":
		typ = "float32"
	case "decimal", "double":
		typ = "float64"
	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		typ = "[]byte"
	case "timestamp", "datetime", "date":
		typ = "time.Time"
	case "time":
		// time is not supported by the MySQL driver. Can parse the string to time.Time in the user code.
		typ = "string"
	default:
	}
	return typ
}

// PrecScaleRE is the regexp that matches "(precision[,scale])" definitions in a
// database.
var PrecScaleRE = regexp.MustCompile(`\(([0-9]+)(\s*,[0-9]+)?\)$`)

// ParsePrecision extracts (precision[,scale]) strings from a data type and
// returns the data type without the string.
func ParsePrecision(dt string) (string, int, int) {
	var err error

	precision := -1
	scale := -1

	m := PrecScaleRE.FindStringSubmatchIndex(dt)
	if m != nil {
		// extract precision
		precision, err = strconv.Atoi(dt[m[2]:m[3]])
		if err != nil {
			panic("could not convert precision")
		}

		// extract scale
		if m[4] != -1 {
			scale, err = strconv.Atoi(dt[m[4]+1 : m[5]])
			if err != nil {
				panic("could not convert scale")
			}
		}

		// change dt
		dt = dt[:m[0]] + dt[m[1]:]
	}

	return dt, precision, scale
}

//SQLInsertFields SQLInsertFields
func SQLInsertFields(t *Table, all bool) string {
	var ns []string
	for _, v := range t.Fields {
		if !all {
			if v.IsPrimaryKey || v.ColumnName == "ctime" || v.ColumnName == "mtime" {
				continue
			}
		}
		ns = append(ns, v.ColumnName)
	}
	return strings.Join(ns, ",")
}

// SQLInsertValue SQLInsertValue
func SQLInsertValue(t *Table, all bool) string {
	var ns []string
	for _, v := range t.Fields {
		if !all {
			if v.IsPrimaryKey || v.ColumnName == "ctime" || v.ColumnName == "mtime" {
				continue
			}
		}
		ns = append(ns, "?")
	}
	return strings.Join(ns, ",")
}

//SQLInsertGoValue SQLInsertGoValue
func SQLInsertGoValue(t *Table, all bool) string {
	var ns []string
	for _, v := range t.Fields {
		if !all {
			if v.IsPrimaryKey || v.ColumnName == "ctime" || v.ColumnName == "mtime" {
				continue
			}
		}

		ns = append(ns, "&a."+v.GoName)
	}
	return strings.Join(ns, ",")
}

// SQLUpdateSet SQLUpdateSet
func SQLUpdateSet(t *Table, all bool) string {
	var ns []string
	for _, v := range t.Fields {
		if !all {
			if v.IsPrimaryKey || v.ColumnName == "ctime" || v.ColumnName == "mtime" {
				continue
			}
		}
		ns = append(ns, v.ColumnName+" = ? ")
	}
	return strings.Join(ns, ",")

}

// SQLUpdateGoValue SQLUpdateGoValue
func SQLUpdateGoValue(t *Table, all bool) string {
	var ns []string
	var prime *Column
	for _, v := range t.Fields {
		if !all {
			if v.ColumnName == "ctime" || v.ColumnName == "mtime" {
				continue
			}
		}
		if v.IsPrimaryKey {
			prime = v
			continue
		}
		ns = append(ns, "a."+v.GoName)
	}
	ns = append(ns, "a."+prime.GoName)

	return strings.Join(ns, ",")
}

// SQLIndexParamList SQLIndexParamList
func SQLIndexParamList(index *Index, needType bool) string {
	var ns []string
	for _, v := range index.IndexColumns {
		s := v.ColumnName
		if needType {
			s = s + " " + v.GoType
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
	return strings.Join(ns, " AND ")
}
