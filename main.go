package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/hongshengjie/curd/model"
	"github.com/hongshengjie/curd/tmpl"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string
var table string
var schema string

func main() {
	// 读取数据库连接地址，table
	flag.StringVar(&dsn, "dsn", "", "mysql connection url")
	flag.StringVar(&schema, "schema", "", "schema name")
	flag.StringVar(&table, "table", "", "table name")
	flag.Parse()
	if dsn == "" || table == "" || schema == "" {
		fmt.Println("dns or schema or table is empty")
		os.Exit(0)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	mytable := model.NewTable(db, schema, table)
	f := template.FuncMap{
		"field":         model.SQLInsertFields,
		"value":         model.SQLInsertValue,
		"govalue":       model.SQLInsertGoValue,
		"updateset":     model.SQLUpdateSet,
		"updategovalue": model.SQLUpdateGoValue,
		"goparamlist":   model.SQLIndexParamList,
		"query":         model.SQLIndexQuery,
	}
	b, _ := tmpl.Asset("table.tmpl")
	//b,_:=ioutil.ReadFile("tmpl/table.tmpl")
	tpl, err := template.New("mysql").Funcs(f).Parse(string(b))
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, &mytable)
	if err != nil {
		panic(err)
	}
}
