package main

import (
	"context"
	"crud/user"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dsn = "root:root@tcp(127.0.0.1:3306)/example?parseTime=true"
var ctx = context.Background()

func init() {

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
func main() {
	//create()
	findOne()
	find()
	update()
	count()
	tx()

}
func create() {

	user.Insert(ctx, db, &user.User{
		Name: "xiaoming",
		Age:  12,
	})
}

func findOne() {
	user, err := user.FindOne(ctx, db, user.Query(user.Where().Eq(user.ID, 1)))
	fmt.Println(user, err)
}

func find() {
	user, err := user.Find(ctx, db, user.Query(user.Where().Gt(user.ID, 0)))
	b, _ := json.Marshal(user)
	fmt.Printf("%s, %+v", string(b), err)
}

func update() {
	user.Update(user.Where().Eq(user.ID, 1)).SetAge(1).SetName("java").Save(ctx, db)
}

func count() {
	c, err := user.Count(ctx, db, user.Where().Eq(user.Age, 12).And().NotEq(user.ID, 1))
	fmt.Println(c, err)
}

func tx() {
	tx, _ := db.Begin()
	_, err := user.Update(user.Where().Eq(user.ID, 1)).SetAge(120).Save(ctx, tx)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	use, err := user.FindOne(ctx, tx, user.Query(user.Where().Eq(user.ID, 1)))
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	fmt.Println(*use)
	tx.Commit()
}
