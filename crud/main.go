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

	"github.com/hongshengjie/crud/mytable"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string
var table string

var schema string

//var tmpl string

// go:generate go-bindata -o mytable/templates.go -pkg mytable  templates
func main() {
	// 读取数据库连接地址，table
	flag.StringVar(&dsn, "dsn", "", "mysql connection url")
	flag.StringVar(&table, "table", "", "table name")
	//flag.StringVar(&schema, "schema", "", "schema name")
	//flag.StringVar(&tmpl, "tmpl", "", "template name")
	flag.Parse()
	if dsn == "" || table == "" {
		fmt.Println("dns or schema or table is empty")
		os.Exit(0)
	}
	// if tmpl == "" {
	// 	tmpl = "default.tmpl"
	// }
	//   -dsn='root:root@tcp(127.0.0.1:3306)/example?parseTime=true'
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
		// "sqlupdate":   mytable.SQLUpdate,
		// "goparamlist": mytable.SQLIndexParamList,
		// "query":       mytable.SQLIndexQuery,
	}
	b, err := mytable.Asset("templates/default.tmpl")
	if err != nil {
		panic(err)
	}
	//b, _ := ioutil.ReadFile("templates/table_bilibili.tmpl")
	tpl, err := template.New("mysql").Funcs(f).Parse(string(b))
	if err != nil {
		panic(err)
	}
	bs := bytes.NewBuffer(nil)
	err = tpl.Execute(bs, &table)
	if err != nil {
		panic(err)
	}
	//copy := bs.String()
	// gofmt := exec.Command("gofmt")
	// gofmt.Stdin = bs
	// gofmt.Stdout = os.Stdout
	// gofmt.Stderr = os.Stderr
	// gofmt.Run()
	//fmt.Println("end")
	result, err := format.Source(bs.Bytes())
	if err != nil {
		panic(err)
	}
	//创建目录
	os.Mkdir(table.PackageName, os.ModePerm)
	//写文件
	fileName := table.PackageName + "/" + table.PackageName + ".go"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	}
	file.Write(result)
	file.Close()

}
