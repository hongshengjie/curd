package crud

import (
	"bytes"
	"fmt"
	"strings"
)

type Where struct {
	Fields    []string
	Args      []interface{}
	statement string
}

type WhereBuilder interface {
	Eq(field string, arg interface{}) WhereBuilder
	NotEq(field string, arg interface{}) WhereBuilder
	Gt(field string, arg interface{}) WhereBuilder
	GtE(field string, arg interface{}) WhereBuilder
	Lt(field string, arg interface{}) WhereBuilder
	LtE(field string, arg interface{}) WhereBuilder
	In(field string, args []interface{}) WhereBuilder
	NotIn(field string, args []interface{}) WhereBuilder

	Between(field string, arg1, arg2 interface{}) WhereBuilder
	NotBetween(field string, arg1, arg2 interface{}) WhereBuilder
	Like(field string, arg string) WhereBuilder
	NotLike(filed string, arg string) WhereBuilder

	And() WhereBuilder
	Or() WhereBuilder
	Builder
}

func (w *Where) Eq(field string, arg interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+" = ? ")
	w.Args = append(w.Args, arg)
	return w

}

func (w *Where) NotEq(field string, arg interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+" <> ? ")
	w.Args = append(w.Args, arg)

	return w
}

func (w *Where) Gt(field string, arg interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+" > ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) GtE(field string, arg interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+" >= ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) Lt(field string, arg interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+" < ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) LtE(field string, arg interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+" <= ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) In(field string, args []interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+` IN (?`+strings.Repeat(",?", len(args)-1)+`)`)
	w.Args = append(w.Args, args...)
	return w
}

func (w *Where) NotIn(field string, args []interface{}) WhereBuilder {
	w.Fields = append(w.Fields, field+` NOT IN (?`+strings.Repeat(",?", len(args)-1)+`)`)
	w.Args = append(w.Args, args...)
	return w
}

func (w *Where) And() WhereBuilder {
	w.Fields = append(w.Fields, " AND ")
	return w
}

func (w *Where) Or() WhereBuilder {
	w.Fields = append(w.Fields, " OR ")
	return w
}

func (w *Where) Between(field string, arg1, arg2 interface{}) WhereBuilder {
	w.Fields = append(w.Fields, fmt.Sprintf(" %s BETWEEN  %v AND %v", field, arg1, arg2))
	return w
}

func (w *Where) NotBetween(field string, arg1, arg2 interface{}) WhereBuilder {
	w.Fields = append(w.Fields, fmt.Sprintf(" %s NOT BETWEEN  %v AND %v", field, arg1, arg2))
	return w
}

func (w *Where) Like(field string, arg string) WhereBuilder {
	w.Fields = append(w.Fields, fmt.Sprintf(" %s LIKE '%s'", field, arg))
	return w
}

func (w *Where) NotLike(field, arg string) WhereBuilder {
	w.Fields = append(w.Fields, fmt.Sprintf(" %s NOT LIKE '%s'", field, arg))
	return w
}

func (w *Where) Build() (string, []interface{}) {
	if w.statement != "" {
		return w.statement, w.Args
	}
	if len(w.Fields) <= 0 {
		return w.statement, w.Args
	}
	var build bytes.Buffer
	build.WriteString(" WHERE ")
	for _, v := range w.Fields {
		build.WriteString(v)
	}
	w.statement = build.String()

	return w.statement, w.Args
}
