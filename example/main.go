package main

import (
	"context"
	"time"

	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hongshengjie/crud/example/alltypetable"
	"github.com/hongshengjie/crud/example/user"
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
func main() {
	InitDB()
	//createAllType()
	updateAllType()
	//create()
	//findOne()
	//find()
	//updateUser()
	//count()
	//tx()
	db.Close()

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
	user, err := user.Find(ctx, db, user.Query(user.
		Where().
		NotIn(user.ID, []interface{}{1, 2})).
		OrderDesc(user.ID).
		OrderAsc(user.Age).
		Limit(2).Offset(2))
	b, _ := json.Marshal(user)
	fmt.Printf("%s, %+v \n", string(b), err)
}

func updateUser() {
	eff, err := user.Update(user.Where().Eq(user.ID, 1)).
		SetAge(10).
		SetName("python").
		SetCtime(time.Now()).
		Save(ctx, db)
	fmt.Println(eff, err)
}

func updateAllType() {
	ef, err := alltypetable.Update(alltypetable.Where().Eq(alltypetable.ID, 1)).
		SetBInt(11).SetBinaryM([]byte("java")).SetEnumM("n").SetSetM("c").Save(ctx, db)
	fmt.Println(ef, err)

}

func createAllType() {
	err := alltypetable.Insert(ctx, db, &alltypetable.AllTypeTable{
		ID:              0,
		TInt:            0,
		SInt:            0,
		MInt:            0,
		BInt:            0,
		F32:             0,
		F64:             0,
		DecimalMysql:    0,
		CharM:           "",
		VarcharM:        "",
		JSONM:           "{}",
		NvarcharM:       "",
		NcharM:          "",
		TimeM:           "23:59:59",
		DateM:           time.Time{},
		DataTimeM:       time.Time{},
		TimestampM:      time.Time{},
		TimestampUpdate: time.Time{},
		YearM:           "2021",
		TText:           "",
		MText:           "",
		TextM:           "",
		LText:           "",
		BinaryM:         []byte{},
		BlobM:           []byte{},
		LBlob:           []byte{},
		MBlob:           []byte{},
		TBlob:           []byte{},
		BitM:            []uint8{},
		EnumM:           "y",
		SetM:            "a",
		BoolM:           0,
	})
	fmt.Println(err)
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
