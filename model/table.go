package model

type Table struct {
	Name       string // table name
	GoName     string // go struct name
	Fields     []*Column
	Indexes    []*Index
	PrimaryKey *Column
}
