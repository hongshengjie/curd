package main

import (
	"bytes"
	"database/sql"
	"flag"

	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	_ "embed"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hongshengjie/crud/mytable"
)

//go:embed "templates/model.go.tmpl"
var modelTmpl []byte

//go:embed "templates/crud_mysql.go.tmpl"
var crudTmpl []byte

//go:embed "templates/where.go.tmpl"
var whereTmpl []byte

var dsn string
var table string
var tidb bool
var fields string

func init() {
	flag.StringVar(&dsn, "dsn", "", "mysql connection url")
	flag.StringVar(&table, "table", "", "table name")
	flag.StringVar(&fields, "fields", "", "split by space, mark table‘s fields that can generate where condition method，default generate all index fields ; if fields = all generate all fields ;if fileds = id,xx,xxx,ctime generate id xx xxx citme fileds ")
}

func main() {

	flag.Parse()
	if dsn == "" || table == "" {
		log.Fatal("dns or schema or table is empty")
	}

	temps := strings.Split(dsn, "/")
	if len(temps) < 2 {
		log.Fatal("dsn not hava /")
	}
	temps2 := strings.Split(temps[1], "?")
	if len(temps2) < 2 {
		log.Fatal("dsn not hava ?")
	}
	schema := temps2[0]

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	var isAll bool

	if strings.TrimSpace(fields) == "all" {
		isAll = true
	}
	conditionFields := strings.Split(fields, ",")
	table := mytable.NewTable(db, schema, table, conditionFields, isAll)
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
		log.Fatal(err)
	}
	bs := bytes.NewBuffer(nil)
	err = tpl.Execute(bs, table)
	if err != nil {
		log.Fatal(err)
	}

	result, err := format.Source(bs.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	//写文件
	fileName := table.PackageName + "/" + name + ".go"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(result)
	file.Close()
}
