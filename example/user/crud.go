package user

import (
	"context"
	"errors"
	"time"

	"github.com/hongshengjie/crud/xsql"
)

// InsertBuilder InsertBuilder
type InsertBuilder struct {
	dt      xsql.DBTX
	builder *xsql.InsertBuilder
	a       *User
}

// Create Create
func Create(dt xsql.DBTX) *InsertBuilder {
	return &InsertBuilder{
		builder: xsql.Dialect(dialect).Insert(table),
		dt:      dt,
	}
}

// Table  涉及到分表查询的时候指定表名，通常情况下不需要调用此方法
func (in *InsertBuilder) Table(name string) *InsertBuilder {
	in.builder.Table(name)
	return in
}

// SetUser SetUser
func (in *InsertBuilder) SetUser(a *User) *InsertBuilder {
	in.a = a
	return in
}

// Save Save
func (in *InsertBuilder) Save(ctx context.Context) error {
	if in.a == nil {
		return errors.New("please set a User")
	}
	ins, args := in.builder.Columns(Name, Age).
		Values(&in.a.Name, &in.a.Age).
		Query()
	result, err := in.dt.ExecContext(ctx, ins, args...)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	in.a.ID = uint32(id)

	return nil
}

// DeleteBuilder DeleteBuilder
type DeleteBuilder struct {
	builder *xsql.DeleteBuilder
	dt      xsql.DBTX
}

// Delete Delete
func Delete(dt xsql.DBTX) *DeleteBuilder {
	return &DeleteBuilder{
		builder: xsql.Dialect(dialect).Delete(table),
		dt:      dt,
	}
}

// Table  涉及到分表查询的时候指定表名，通常情况下不需要调用此方法
func (d *DeleteBuilder) Table(name string) *DeleteBuilder {
	d.builder.Table(name)
	return d
}

// Where  UserWhere
func (d *DeleteBuilder) Where(p ...UserWhere) *DeleteBuilder {
	s := &xsql.Selector{}
	for _, v := range p {
		v(s)
	}
	d.builder = d.builder.Where(s.P())
	return d
}

// WhereP arg is Predicate
func (d *DeleteBuilder) WhereP(p *xsql.Predicate) *DeleteBuilder {
	d.builder = d.builder.Where(p)
	return d
}

// ByID  delete by primary key
func (d *DeleteBuilder) ByID(id uint32) *DeleteBuilder {
	d.builder = d.builder.Where(xsql.EQ(ID, id))
	return d
}

// Exec Exec
func (d *DeleteBuilder) Exec(ctx context.Context) (int64, error) {
	del, args := d.builder.Query()
	res, err := d.dt.ExecContext(ctx, del, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// SelectBuilder SelectBuilder
type SelectBuilder struct {
	builder *xsql.Selector
	dt      xsql.DBTX
}

// Find Find
func Find(dt xsql.DBTX) *SelectBuilder {
	sel := &SelectBuilder{
		builder: xsql.Dialect(dialect).Select(),
		dt:      dt,
	}
	sel.builder = sel.builder.From(xsql.Table(table))
	return sel
}

// Select Select
func (s *SelectBuilder) Select(columns ...string) *SelectBuilder {
	s.builder.Select(columns...)
	return s
}

// Count Count
func (s *SelectBuilder) Count(columns ...string) *SelectBuilder {
	s.builder.Count(columns...)
	return s
}

// Table  涉及到分表查询的时候指定表名，通常情况下不需要调用此方法
func (s *SelectBuilder) Table(name string) *SelectBuilder {
	s.builder.From(xsql.Table(name))
	return s
}

// WhereP WhereP
func (s *SelectBuilder) WhereP(p *xsql.Predicate) *SelectBuilder {
	s.builder = s.builder.Where(p)
	return s
}

//Where where
func (s *SelectBuilder) Where(p ...UserWhere) *SelectBuilder {
	sel := &xsql.Selector{}
	for _, v := range p {
		v(sel)
	}
	s.builder = s.builder.Where(sel.P())
	return s
}

// ByID select by primary key
func (s *SelectBuilder) ByID(id uint32) *SelectBuilder {
	s.builder.Where(xsql.EQ(ID, id))
	return s
}

// Offset Offset
func (s *SelectBuilder) Offset(offset int) *SelectBuilder {
	s.builder = s.builder.Offset(offset)
	return s
}

// Limit Limit
func (s *SelectBuilder) Limit(limit int) *SelectBuilder {
	s.builder = s.builder.Limit(limit)
	return s
}

// OrderDesc OrderDesc
func (s *SelectBuilder) OrderDesc(field string) *SelectBuilder {
	s.builder = s.builder.OrderBy(xsql.Desc(field))
	return s
}

// OrderAsc OrderAsc
func (s *SelectBuilder) OrderAsc(field string) *SelectBuilder {
	s.builder = s.builder.OrderBy(xsql.Asc(field))
	return s
}

// One One
func (s *SelectBuilder) One(ctx context.Context) (*User, error) {
	results, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(results) <= 0 {
		return nil, xsql.ErrNoRows
	}
	return results[0], nil
}

// Int64 count or select only one int field
func (s *SelectBuilder) Int64(ctx context.Context) (int64, error) {
	return xsql.Int64(ctx, s.builder, s.dt)
}

// Int64 count or select only one int field
func (s *SelectBuilder) Int64s(ctx context.Context) ([]int64, error) {
	return xsql.Int64s(ctx, s.builder, s.dt)
}

// Query FindRaw many record by raw sql
func (s *SelectBuilder) String(ctx context.Context) (string, error) {
	return xsql.String(ctx, s.builder, s.dt)
}

// Query FindRaw many record by raw sql
func (s *SelectBuilder) Strings(ctx context.Context) ([]string, error) {
	return xsql.Strings(ctx, s.builder, s.dt)
}

// All FindRaw many record by raw sql
func (s *SelectBuilder) All(ctx context.Context) ([]*User, error) {
	if len(s.builder.Columns()) <= 0 {
		s.builder.Columns(Columns...)
	}
	sqlstr, args := s.builder.Query()
	q, err := s.dt.QueryContext(ctx, sqlstr, args...)
	if err != nil {
		return nil, err
	}
	defer q.Close()
	res := []*User{}
	if err := xsql.ScanSlice(q, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateBuilder UpdateBuilder
type UpdateBuilder struct {
	builder *xsql.UpdateBuilder
	dt      xsql.DBTX
}

// Update return a UpdateBuilder
func Update(dt xsql.DBTX) *UpdateBuilder {
	return &UpdateBuilder{
		dt:      dt,
		builder: xsql.Dialect(dialect).Update(table),
	}
}

// Table  for custom demand  change table name  for example   sub table system
// In the case of no sub-table you do not need this method
func (u *UpdateBuilder) Table(name string) *UpdateBuilder {
	u.builder.Table(name)
	return u
}

// WhereP WhereP
func (u *UpdateBuilder) WhereP(p *xsql.Predicate) *UpdateBuilder {
	u.builder.Where(p)
	return u
}

// Where Where
func (u *UpdateBuilder) Where(p ...UserWhere) *UpdateBuilder {
	s := &xsql.Selector{}
	for _, v := range p {
		v(s)
	}
	u.builder = u.builder.Where(s.P())
	return u
}

// SetID  set id
func (u *UpdateBuilder) SetID(arg uint32) *UpdateBuilder {
	u.builder.Set(ID, arg)
	return u
}

// SetName  set name
func (u *UpdateBuilder) SetName(arg string) *UpdateBuilder {
	u.builder.Set(Name, arg)
	return u
}

// SetAge  set age
func (u *UpdateBuilder) SetAge(arg int32) *UpdateBuilder {
	u.builder.Set(Age, arg)
	return u
}

// AddAge  add  age set x = x + arg
func (u *UpdateBuilder) AddAge(arg interface{}) *UpdateBuilder {
	u.builder.Add(Age, arg)
	return u
}

// SetCtime  set ctime
func (u *UpdateBuilder) SetCtime(arg time.Time) *UpdateBuilder {
	u.builder.Set(Ctime, arg)
	return u
}

// SetMtime  set mtime
func (u *UpdateBuilder) SetMtime(arg time.Time) *UpdateBuilder {
	u.builder.Set(Mtime, arg)
	return u
}

// ByID  update by primary key
func (u *UpdateBuilder) ByID(id uint32) *UpdateBuilder {
	u.builder.Where(xsql.EQ(ID, id))
	return u
}

// Save do a update statment  if tx can without context
func (u *UpdateBuilder) Save(ctx context.Context) (int64, error) {
	up, args := u.builder.Query()
	result, err := u.dt.ExecContext(ctx, up, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
