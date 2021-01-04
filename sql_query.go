package crud

import (
	"bytes"
	"fmt"
	"strings"
)

type Query struct {
	Where       WhereBuilder
	Table       string
	AllFields   string
	offset      int64
	limit       int64
	orderBy     []string
	selectField []string
	sql         string
}

type QueryBuilder interface {
	// Select(fields ...string) QueryBuilder
	Limit(l int64) QueryBuilder
	Offset(l int64) QueryBuilder
	OrderAsc(field string) QueryBuilder
	OrderDesc(field string) QueryBuilder
	Builder
}

type Builder interface {
	Build() (string, []interface{})
}

func (s *Query) Limit(l int64) QueryBuilder {
	s.limit = l
	return s
}
func (s *Query) Offset(o int64) QueryBuilder {
	s.offset = o
	return s
}

func (s *Query) OrderAsc(field string) QueryBuilder {
	s.orderBy = append(s.orderBy, field+" ASC")
	return s
}

func (s *Query) OrderDesc(field string) QueryBuilder {
	s.orderBy = append(s.orderBy, field+" DESC")
	return s
}

func (s *Query) Build() (string, []interface{}) {
	var where string
	var args []interface{}
	if s.Where != nil {
		where, args = s.Where.Build()
	}

	if s.sql != "" {
		return s.sql, args
	}
	var b bytes.Buffer
	b.WriteString("SELECT ")
	b.WriteString(s.AllFields)
	b.WriteString(" FROM ")
	b.WriteString(s.Table)
	if where != "" {
		b.WriteString(where)
	}

	if len(s.orderBy) > 0 {
		b.WriteString(" ORDER BY ")
		b.WriteString(strings.Join(s.orderBy, ","))
	}

	if s.limit > 0 {
		if s.offset < 0 {
			s.offset = 0
		}
		b.WriteString(fmt.Sprintf(" LIMIT %d , OFFSET %d", s.limit, s.offset))
	}
	return b.String(), args
}
