package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dsn = "root:1234@tcp(127.0.0.1:3306)/stock?parseTime=true"
var ctx = context.Background()

func init() {

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func TestAllTypeTable_Insert(t *testing.T) {

	a := &AllTypeTable{
		TInt:            9,
		SInt:            3,
		MInt:            4,
		BInt:            5,
		F32:             5.12121,
		F64:             6.33,
		DecimalMysql:    9.3333,
		CharM:           "v",
		VarcharM:        "fsdf",
		JSONM:           `{"java":1}`,
		NvarcharM:       "dd",
		NcharM:          "dd",
		TimeM:           "20:00:01",
		DateM:           time.Now(),
		DataTimeM:       time.Now(),
		TimestampM:      time.Now(),
		TimestampUpdate: time.Now(),
		YearM:           "2021",
		TText:           "dfsdf",
		MText:           "fdsfs",
		TextM:           "fsdfs",
		LText:           "sfsdf",
		BinaryM:         []byte("jsdjfs"),
		BlobM:           []byte("jsdjfs"),
		LBlob:           []byte("jsdjfs"),
		MBlob:           []byte("jsdjfs"),
		TBlob:           []byte("jsdjfs"),
		BitM:            []uint8{64},
		EnumM:           "n",
		SetM:            "d,d",
	}
	err := a.Insert(db)
	fmt.Println(err, a.ID)

}

func TestAllTypeTableByID(t *testing.T) {

	got, err := AllTypeTableByID(db, 8)
	fmt.Println(got, err)

	gots, err := AllTypeTableByTIntSInt(db, 2, 3)
	b, _ := json.MarshalIndent(gots, "	", "	")
	fmt.Println(string(b))

	gots2, err := AllTypeTableByBIntCharM(db, 5, "j")
	b, _ = json.MarshalIndent(gots2, "	", "	")
	fmt.Println(string(b))

	t2, _ := time.Parse("2006-01-02 15:04:05", "2020-02-09 14:02:45")
	gots3, err := AllTypeTableByMIntVarcharMTimestampM(db, 4, "fsdf", t2)
	b, _ = json.MarshalIndent(gots3, "	", "	")
	fmt.Println(string(b), err)

}
