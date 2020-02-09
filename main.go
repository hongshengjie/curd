package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/hongshengjie/curd/mytable"
	"github.com/hongshengjie/curd/templates"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string
var table string
var schema string
var tmpl string

// go:generate go-bindata -o templates/templates.go -pkg templates  templates
func main() {
	// 读取数据库连接地址，table
	flag.StringVar(&dsn, "dsn", "", "mysql connection url")
	flag.StringVar(&schema, "schema", "", "schema name")
	flag.StringVar(&table, "table", "", "table name")
	flag.StringVar(&tmpl, "tmpl", "", "table name")
	flag.Parse()
	if dsn == "" || table == "" || schema == "" {
		fmt.Println("dns or schema or table is empty")
		os.Exit(0)
	}
	if tmpl == "" {
		tmpl = "table_std.tmpl"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	table := mytable.NewTable(db, schema, table)
	f := template.FuncMap{
		"sqltool":     mytable.SQLTool,
		"sqlupdate":   mytable.SQLUpdate,
		"goparamlist": mytable.SQLIndexParamList,
		"query":       mytable.SQLIndexQuery,
	}
	b, err := templates.Asset("templates/" + tmpl)
	if err != nil {
		panic(err)
	}
	//b, _ := ioutil.ReadFile("templates/table_bilibili.tmpl")
	tpl, err := template.New("mysql").Funcs(f).Parse(string(b))
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, &table)
	if err != nil {
		panic(err)
	}
}
