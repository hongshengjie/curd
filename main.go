package main

import (
	"curd/model"
	"curd/tmpl"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/knq/snaker"
	"os"
	"text/template"
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

	mytable := model.Table{
		Name:   table,
		GoName: snaker.SnakeToCamel(table),
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	columns, err := model.MyTableColumns(db, schema, table)
	if err != nil {
		panic(err)
	}

	indexs, err := model.MyTableIndexes(db, schema, table)
	if err != nil {
		panic(err)
	}

	for _, v := range indexs {
		// index column
		indexcolumn, err := model.MyIndexColumns(db, schema, table, v.IndexName)
		if err != nil {
			panic(err)
		}
		var indexGoName string
		for _, d := range indexcolumn {
			for _, c := range columns {
				if c.ColumnName == d.IndexColumnName {
					d.Column = c
					break
				}
			}
			indexGoName = indexGoName + d.GoName

		}

		v.IndexColumns = indexcolumn
		v.IndexGoName = indexGoName

	}
	mytable.Indexes = indexs
	mytable.Fields = columns
	for _, v := range columns {
		if v.IsPrimaryKey {
			mytable.PrimaryKey = v
			break
		}
	}

	f := template.FuncMap{
		"field":         model.SqlInsertFields,
		"value":         model.SqlInsertValue,
		"govalue":       model.SqlInsertGoValue,
		"updateset":     model.SqlUpdateSet,
		"updategovalue": model.SqlUpdateGoValue,
		"goparamlist":   model.SqlIndexParamList,
		"query":         model.SqlIndexQuery,
	}
	b, _ := tmpl.Asset("table.tmpl")
	tpl, err := template.New("mysql").Funcs(f).Parse(string(b))
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, &mytable)
	if err != nil {
		panic(err)
	}
}
