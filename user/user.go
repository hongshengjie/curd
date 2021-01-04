package user

import (
	"bytes"
	"context"
	"crud"
	"errors"
	"fmt"
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
	// ID id字段
	ID = "`id`"
	// Name 名称
	Name = "`name`"
	// Age 年龄
	Age = "`age`"
	// Ctime 创建时间
	Ctime = "`ctime`"
)

// Insert  Insert a record
func Insert(ctx context.Context, db crud.DB, a *User) error {
	const sqlstr = "INSERT INTO  `user` (" +
		"`name`,`age`,`ctime`" +
		`) VALUES (` +
		` ?,?,?` +
		`)`

	result, err := db.ExecContext(ctx, sqlstr, &a.Name, &a.Age, &a.Ctime)
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

// UserDelete Delete by primary key:`id`
func Delete(ctx context.Context, db crud.DB, id uint32) error {
	const sqlstr = "DELETE FROM `user` WHERE id = ?"
	_, err := db.ExecContext(ctx, sqlstr, id)
	if err != nil {
		return err
	}
	return nil
}

func Where() *crud.Where {
	return &crud.Where{}
}

func Query(where crud.WhereBuilder) crud.QueryBuilder {
	return &crud.Query{
		Table:     table,
		AllFields: fields,
		Where:     where,
	}
}

func Find(ctx context.Context, db crud.DB, build crud.Builder) ([]*User, error) {
	sql, args := build.Build()
	return FindRaw(ctx, db, sql, args...)
}

func FindRaw(ctx context.Context, db crud.DB, sql string, args ...interface{}) ([]*User, error) {
	fmt.Println(sql, args)
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

func FindOne(ctx context.Context, db crud.DB, build crud.Builder) (*User, error) {
	sql, args := build.Build()
	return FindOneRaw(ctx, db, sql, args...)

}

func FindOneRaw(ctx context.Context, db crud.DB, sql string, args ...interface{}) (*User, error) {
	fmt.Println(sql, args)
	a := User{}
	err := db.QueryRowContext(ctx, sql, args...).Scan(&a.ID, &a.Name, &a.Age, &a.Ctime)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func Count(ctx context.Context, db crud.DB, build crud.Builder) (int64, error) {

	sql, args := build.Build()
	fmt.Println(sql, args)
	var a int64
	if err := db.QueryRowContext(ctx, "SELECT COUNT(1) FROM `user` "+sql, args...).Scan(&a); err != nil {
		return 0, err
	}
	return a, nil
}

type UpdateBuilder interface {
	SetName(arg interface{}) UpdateBuilder
	SetAge(arg interface{}) UpdateBuilder
	SetCtime(arg interface{}) UpdateBuilder
	Save(ctx context.Context, db crud.DB) (int64, error)
}

type Updater struct {
	where  crud.WhereBuilder
	Args   []interface{}
	Fields []string
}

func Update(where crud.WhereBuilder) *Updater {
	return &Updater{
		where: where,
	}
}

func (u *Updater) SetName(arg interface{}) UpdateBuilder {
	u.Fields = append(u.Fields, "`name` = ? ")
	u.Args = append(u.Args, arg)
	return u

}
func (u *Updater) SetAge(arg interface{}) UpdateBuilder {
	u.Fields = append(u.Fields, "`age` = ? ")
	u.Args = append(u.Args, arg)
	return u

}
func (u *Updater) SetCtime(arg interface{}) UpdateBuilder {
	u.Fields = append(u.Fields, "`ctime` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// TODO 校验条件
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
	fmt.Println(b.String(), u.Args)
	result, err := db.ExecContext(ctx, b.String(), u.Args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
