package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/hongshengjie/curd/mytable"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string
var table string
var schema string
var tmpl string

// go:generate go-bindata -o mytable/templates.go -pkg mytable  templates
func main() {
	// 读取数据库连接地址，table
	flag.StringVar(&dsn, "dsn", "", "mysql connection url")
	flag.StringVar(&schema, "schema", "", "schema name")
	flag.StringVar(&table, "table", "", "table name")
	flag.StringVar(&tmpl, "tmpl", "", "template name")
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
	b, err := mytable.Asset("templates/" + tmpl)
	if err != nil {
		panic(err)
	}
	//b, _ := ioutil.ReadFile("templates/table_bilibili.tmpl")
	tpl, err := template.New("mysql").Funcs(f).Parse(string(b))
	if err != nil {
		panic(err)
	}
	bs := bytes.NewBufferString("")
	err = tpl.Execute(bs, &table)
	if err != nil {
		panic(err)
	}
	gofmt := exec.Command("gofmt")
	gofmt.Stdin = bs
	gofmt.Stdout = os.Stdout
	gofmt.Run()

}
