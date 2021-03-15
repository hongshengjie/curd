package user

import (
	"time"
)

// User represents a row from 'user'.
type User struct {
	ID    uint32    `json:"id"`    // id字段
	Name  string    `json:"name"`  // 名称
	Age   int32     `json:"age"`   // 年龄
	Ctime time.Time `json:"ctime"` // 创建时间
	Mtime time.Time `json:"mtime"` //
}

const (
	// table tableName is user
	table = "user"
	//ID id字段
	ID = "id"
	//Name 名称
	Name = "name"
	//Age 年龄
	Age = "age"
	//Ctime 创建时间
	Ctime = "ctime"
	//Mtime
	Mtime = "mtime"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	ID,
	Name,
	Age,
	Ctime,
	Mtime,
}

var dialect = "mysql"
