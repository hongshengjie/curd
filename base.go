package crud

import (
	"fmt"
	"strings"
)

type Where struct {
	StringBuilder strings.Builder
	Args          []interface{}
	Methods       []func()
	whereString   string
	Table         string
	Fields        string
}

type Method func()

type WhereBuilder interface {
	Eq(field string, arg interface{}) WhereBuilder
	NotEq(field string, arg interface{}) WhereBuilder
	Gt(field string, arg interface{}) WhereBuilder
	GtE(field string, arg interface{}) WhereBuilder
	Lt(field string, arg interface{}) WhereBuilder
	LtE(field string, arg interface{}) WhereBuilder
	In(field string, args []interface{}) WhereBuilder
	NotIn(field string, args []interface{}) WhereBuilder
	And() WhereBuilder
	Or() WhereBuilder
	ArgSlice() []interface{}
	WhereString() string
	Query() QueryBuilder
}

func (w *Where) Query() QueryBuilder {
	return &Query{
		Table:        w.Table,
		QueryBuilder: strings.Builder{},
		Where:        w,
		Fields:       w.Fields,
	}

}

func (w *Where) ArgSlice() []interface{} {
	return w.Args
}

func (w *Where) Eq(field string, arg interface{}) WhereBuilder {

	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(" = ? ")

	w.Args = append(w.Args, arg)
	return w

}

func (w *Where) NotEq(field string, arg interface{}) WhereBuilder {

	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(" != ? ")

	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) Gt(field string, arg interface{}) WhereBuilder {

	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(" > ? ")

	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) GtE(field string, arg interface{}) WhereBuilder {
	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(" >= ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) Lt(field string, arg interface{}) WhereBuilder {
	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(" < ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) LtE(field string, arg interface{}) WhereBuilder {
	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(" <= ? ")
	w.Args = append(w.Args, arg)
	return w
}

func (w *Where) In(field string, args []interface{}) WhereBuilder {
	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(` IN (?` + strings.Repeat(",?", len(args)-1) + `)`)
	w.Args = append(w.Args, args...)
	return w
}

func (w *Where) NotIn(field string, args []interface{}) WhereBuilder {
	w.StringBuilder.WriteString(field)
	w.StringBuilder.WriteString(` NOT IN (?` + strings.Repeat(",?", len(args)-1) + `)`)
	w.Args = append(w.Args, args...)
	return w
}

func (w *Where) And() WhereBuilder {
	w.StringBuilder.WriteString(" AND ")
	return w
}

func (w *Where) Or() WhereBuilder {
	w.StringBuilder.WriteString(" OR ")
	return w
}

func (w *Where) WhereString() string {
	if w.whereString != "" {
		return w.whereString
	}
	w.whereString = w.StringBuilder.String()
	if strings.TrimSpace(w.whereString) == "" {
		w.whereString = " "
		return ""
	}
	w.whereString = " WHERE " + w.whereString

	return w.whereString
}

func (q *Query) QueryString() string {
	if q.Where.whereString == "" {
		q.Where.WhereString()
	}
	if q.queryString != "" {
		return q.queryString
	}
	q.QueryBuilder.WriteString(q.Where.whereString)

	if len(q.OrderBy) > 0 {
		q.QueryBuilder.WriteString(" ORDER BY ")
		q.QueryBuilder.WriteString(strings.Join(q.OrderBy, ","))
	}

	if q.Limit_ > 0 {
		if q.Offset_ < 0 {
			q.Offset_ = 0
		}
		q.QueryBuilder.WriteString(fmt.Sprintf(" LIMIT %d,%d", q.Offset_, q.Limit_))
	}
	sql := `SELECT ` + q.Fields + ` FROM ` + q.Table + q.QueryBuilder.String()
	q.queryString = sql
	return q.queryString
}

func (q *Query) Limit(l int64) QueryBuilder {
	q.Limit_ = l
	return q
}
func (q *Query) Offset(o int64) QueryBuilder {
	q.Offset_ = o
	return q
}

func (q *Query) OrderAsc(field string) QueryBuilder {
	q.OrderBy = append(q.OrderBy, field+" ASC")
	return q
}

func (q *Query) OrderDesc(field string) QueryBuilder {
	q.OrderBy = append(q.OrderBy, field+" DESC")
	return q
}

type Query struct {
	Table        string
	QueryBuilder strings.Builder
	Offset_      int64
	Limit_       int64
	OrderBy      []string
	Fields       string
	queryString  string
	*Where
}

type QueryBuilder interface {
	WhereBuilder
	Limit(l int64) QueryBuilder
	Offset(l int64) QueryBuilder
	OrderAsc(field string) QueryBuilder
	OrderDesc(field string) QueryBuilder
	QueryString() string
}
