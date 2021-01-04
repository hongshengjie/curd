package user

import (
	"bytes"
	"context"
	"errors"
	"github.com/hongshengjie/crud"
	"strings"
	"time"
)

// User represents a row from 'user'.
type User struct {
	ID    uint32    `json:"id"`    // id字段
	Name  string    `json:"name"`  // 名称
	Age   int32     `json:"age"`   // 年龄
	Ctime time.Time `json:"ctime"` // 创建时间
}

const (
	table  = "`user`"
	fields = "`id`,`name`,`age`,`ctime`"
	//ID id字段
	ID = "`id`"
	//Name 名称
	Name = "`name`"
	//Age 年龄
	Age = "`age`"
	//Ctime 创建时间
	Ctime = "`ctime`"
)

// Insert  Insert a record
func Insert(ctx context.Context, db crud.DB, a *User) error {
	const sqlstr = "INSERT INTO  `user` (" +
		" `name`,`age`" +
		`) VALUES (` +
		` ?,?` +
		`)`

	result, err := db.ExecContext(ctx, sqlstr, &a.Name, &a.Age)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	a.ID = uint32(id)

	return nil
}

// Delete Delete by primary key:`id`
func Delete(ctx context.Context, db crud.DB, id uint32) (int64, error) {
	const sqlstr = `DELETE FROM  user WHERE id = ?`
	result, err := db.ExecContext(ctx, sqlstr, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Where  create where builder for update or query
func Where() *crud.Where {
	return &crud.Where{}
}

// Query create query builder  maybe with where builder
func Query(where crud.WhereBuilder) crud.QueryBuilder {
	return &crud.Query{
		Table:     table,
		AllFields: fields,
		Where:     where,
	}
}

// FindOne find a record
func FindOne(ctx context.Context, db crud.DB, build crud.Builder) (*User, error) {
	sql, args := build.Build()
	return FindOneRaw(ctx, db, sql, args...)
}

// FindOneRaw find a record by raw sql
func FindOneRaw(ctx context.Context, db crud.DB, sql string, args ...interface{}) (*User, error) {
	a := User{}
	err := db.QueryRowContext(ctx, sql, args...).Scan(&a.ID, &a.Name, &a.Age, &a.Ctime)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// Find Find many record by where statment or query statment
func Find(ctx context.Context, db crud.DB, build crud.Builder) ([]*User, error) {
	sql, args := build.Build()
	return FindRaw(ctx, db, sql, args...)
}

// FindRaw FindRaw many record by raw sql
func FindRaw(ctx context.Context, db crud.DB, sql string, args ...interface{}) ([]*User, error) {
	q, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*User{}
	for q.Next() {
		a := User{}
		err = q.Scan(&a.ID, &a.Name, &a.Age, &a.Ctime)
		if err != nil {
			return nil, err
		}
		res = append(res, &a)
	}
	if q.Err() != nil {
		return nil, err
	}

	return res, nil
}

// Count Count return number of rows that fit the where statment
func Count(ctx context.Context, db crud.DB, build crud.Builder) (int64, error) {
	sql, args := build.Build()
	var a int64
	if err := db.QueryRowContext(ctx, "SELECT COUNT(1) FROM `user` "+sql, args...).Scan(&a); err != nil {
		return 0, err
	}
	return a, nil
}

// UpdateBuilder UpdateBuilder
type UpdateBuilder interface {
	SetID(arg uint32) UpdateBuilder
	SetName(arg string) UpdateBuilder
	SetAge(arg int32) UpdateBuilder
	SetCtime(arg time.Time) UpdateBuilder
	Save(ctx context.Context, db crud.DB) (int64, error)
}

// Updater Updater
type Updater struct {
	where  crud.WhereBuilder
	Args   []interface{}
	Fields []string
}

// Update return a Updater
func Update(where crud.WhereBuilder) *Updater {
	return &Updater{
		where: where,
	}
}

// SetID set id
func (u *Updater) SetID(arg uint32) UpdateBuilder {
	u.Fields = append(u.Fields, "`id` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetName set name
func (u *Updater) SetName(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`name` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetAge set age
func (u *Updater) SetAge(arg int32) UpdateBuilder {
	u.Fields = append(u.Fields, "`age` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetCtime set ctime
func (u *Updater) SetCtime(arg time.Time) UpdateBuilder {
	u.Fields = append(u.Fields, "`ctime` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// Save do a update statment
func (u *Updater) Save(ctx context.Context, db crud.DB) (int64, error) {
	var where string
	var args []interface{}
	if u.where != nil {
		where, args = u.where.Build()
	}
	if len(u.Fields) <= 0 {
		return 0, errors.New("not set update fields")
	}

	var b bytes.Buffer
	b.WriteString("UPDATE ")
	b.WriteString(table)
	b.WriteString(" SET ")
	b.WriteString(strings.Join(u.Fields, ","))
	b.WriteString(where)
	u.Args = append(u.Args, args...)
	result, err := db.ExecContext(ctx, b.String(), u.Args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

