package user

import (
	"context"
	"crud"
	"database/sql"
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
	table  = "user"
	fields = "`id`,`name`,`age`,`ctime`"
	// ID id字段
	ID = "id"
	// Name 名称
	Name = "name"
	// Age 年龄
	Age = "age"
	// Ctime 创建时间
	Ctime = "ctime"
)

// Insert  Insert a record
func Insert(ctx context.Context, db *sql.DB, a *User) error {
	var err error

	const sqlstr = `INSERT INTO  user (` +
		` name,age,ctime` +
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

func Query() crud.QueryBuilder {
	return &crud.Query{
		Table:        table,
		QueryBuilder: strings.Builder{},
		Fields:       fields,
		Where:        Where(),
	}
}

func Where() *crud.Where {
	return &crud.Where{
		Fields:        fields,
		Table:         table,
		StringBuilder: strings.Builder{},
	}
}

// UserDelete Delete by primary key:`id`
func Delete(ctx context.Context, db *sql.DB, id uint32) error {
	var err error

	const sqlstr = `DELETE FROM  user WHERE id = ?`

	_, err = db.ExecContext(ctx, sqlstr, id)
	if err != nil {
		return err
	}
	return nil
}

type UpdateBuilder interface {
	SetName(arg interface{}) UpdateBuilder
	SetAge(arg interface{}) UpdateBuilder
	SetCtime(arg interface{}) UpdateBuilder
	Update(ctx context.Context, db *sql.DB, selecter crud.WhereBuilder) error
}

type Updater struct {
	Builder strings.Builder
	Args    []interface{}
	Fields  []string
}

func NewUpdater() *Updater {
	u := &Updater{
		Builder: strings.Builder{},
		Args:    []interface{}{},
		Fields:  []string{},
	}

	return u
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
func (u *Updater) Update(ctx context.Context, db *sql.DB, selecter crud.WhereBuilder) error {
	u.Builder.WriteString("UPDATE user SET ")
	u.Builder.WriteString(strings.Join(u.Fields, ","))
	u.Builder.WriteString(selecter.WhereString())
	fmt.Println(u.Builder.String())
	_, err := db.ExecContext(ctx, u.Builder.String(), append(u.Args, selecter.ArgSlice()...)...)
	return err
}

func Find(ctx context.Context, db *sql.DB, build crud.QueryBuilder) ([]*User, error) {
	fmt.Println(build.QueryString())
	q, err := db.QueryContext(ctx, build.QueryString(), build.ArgSlice()...)
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

func FindOne(ctx context.Context, db *sql.DB, build crud.QueryBuilder) (*User, error) {
	var err error
	a := User{}
	fmt.Println(build.QueryString())
	err = db.QueryRowContext(ctx, build.QueryString(), build.ArgSlice()...).Scan(&a.ID, &a.Name, &a.Age, &a.Ctime)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func Count(ctx context.Context, db *sql.DB, build crud.WhereBuilder) (int64, error) {
	var err error
	var a int64
	fmt.Println(build.WhereString())
	err = db.QueryRowContext(ctx, "SELECT COUNT(1) FROM user "+build.WhereString(), build.ArgSlice()...).Scan(&a)
	if err != nil {
		return 0, err
	}

	return a, nil
}
