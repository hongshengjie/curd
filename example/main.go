package main

import (
	"context"
	"fmt"
	"time"

	"database/sql"

	"github.com/hongshengjie/crud/example/user"
	"github.com/hongshengjie/crud/xsql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dsn = "root:root@tcp(127.0.0.1:3306)/example?parseTime=true"
var ctx = context.Background()

func InitDB() {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func UserExample() {

	u := &user.User{
		Name:  "testA",
		Age:   22,
		Ctime: time.Now(),
	}
	u1 := &user.User{
		Name:  "testb",
		Age:   22,
		Ctime: time.Now(),
	}
	err := user.Create(db).SetUser(u).Save(ctx)
	fmt.Println(err)

	err = user.Create(db).SetUser(u1).Save(ctx)
	fmt.Println(err)

	au, err := user.Find(db).Where(
		user.Or(
			user.IDEQ(10),
			user.NameEQ("testb"),
		)).
		Offset(0).
		Limit(3).
		OrderAsc("name").
		One(ctx)
	fmt.Println(au, err)

	effect, err := user.Update(db).SetAge(10).WhereP(xsql.EQ(user.ID, 1)).Save(ctx)

	effect, err = user.Update(db).SetAge(100).SetName("java").Where(user.IDEQ(1)).Save(ctx)

	effect, err = user.Update(db).AddAge(-100).SetName("java").ByID(5).Save(ctx)

	effect, err = user.Delete(db).Where(user.And(user.IDEQ(3), user.IDIn(1, 3))).Exec(ctx)

	effect, err = user.Delete(db).WhereP(xsql.EQ(user.ID, 32)).Exec(ctx)

	effect, err = user.Delete(db).ByID(2).Exec(ctx)

	tx, _ := db.Begin()
	u2 := &user.User{
		ID:    0,
		Name:  "dfasd",
		Age:   2,
		Ctime: time.Now(),
	}
	err = user.Create(tx).SetUser(u2).Save(ctx)
	tx.Rollback()

	effect, err = user.Update(tx).SetAge(100).ByID(u2.ID).Save(ctx)
	tx.Commit()
	fmt.Println(effect, err)

}
func main() {
	InitDB()
	xsql.Debug()
	//sqlbuild()
	UserExample()
}
