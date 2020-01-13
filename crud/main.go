package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"

	"go/format"
	"os"
	"strings"
	"text/template"

	_ "embed"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hongshengjie/crud/mytable"
)

var dsn string
var table string

var schema string

//go:embed "templates/model.go.tmpl"
var modelTmpl []byte

//go:embed "templates/crud_mysql.go.tmpl"
var crudTmpl []byte

//go:embed "templates/where.go.tmpl"
var whereTmpl []byte

func main() {
	flag.StringVar(&dsn, "dsn", "", "mysql connection url")
	flag.StringVar(&table, "table", "", "table name")
	flag.Parse()

	if dsn == "" || table == "" {
		fmt.Println("dns or schema or table is empty")
		os.Exit(0)
	}

	temps := strings.Split(dsn, "/")
	if len(temps) < 2 {
		panic("dsn not hava /")
	}
	temps2 := strings.Split(temps[1], "?")
	if len(temps2) < 2 {
		panic("dsn not hava ?")
	}
	schema = temps2[0]

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	table := mytable.NewTable(db, schema, table)
	f := template.FuncMap{
		"sqltool": mytable.SQLTool,
	}
	//创建目录
	os.Mkdir(table.PackageName, os.ModePerm)
	generateFile("model", string(modelTmpl), f, table)
	generateFile("where", string(whereTmpl), f, table)
	generateFile("crud", string(crudTmpl), f, table)

}

func generateFile(name, tmpl string, f template.FuncMap, table *mytable.Table) {
	tpl, err := template.New(name).Funcs(f).Parse(string(tmpl))
	if err != nil {
		panic(err)
	}
	bs := bytes.NewBuffer(nil)
	err = tpl.Execute(bs, table)
	if err != nil {
		panic(err)
	}

	result, err := format.Source(bs.Bytes())
	if err != nil {
		panic(err)
	}
	//写文件
	fileName := table.PackageName + "/" + name + ".go"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	}
	file.Write(result)
	file.Close()
}
