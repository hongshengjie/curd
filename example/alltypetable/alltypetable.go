package alltypetable

import (
	"bytes"
	"context"
	"errors"
	"github.com/hongshengjie/crud"
	"strings"
	"time"
)

// AllTypeTable represents a row from 'all_type_table'.
type AllTypeTable struct {
	ID              uint32    `json:"id"`               // 自增id
	TInt            int8      `json:"t_int"`            // 小整型
	SInt            uint16    `json:"s_int"`            //
	MInt            int32     `json:"m_int"`            // 中整数
	BInt            int64     `json:"b_int"`            //
	F32             float32   `json:"f32"`              //
	F64             float64   `json:"f64"`              //
	DecimalMysql    float64   `json:"decimal_mysql"`    //
	CharM           string    `json:"char_m"`           //
	VarcharM        string    `json:"varchar_m"`        //
	JSONM           string    `json:"json_m"`           //
	NvarcharM       string    `json:"nvarchar_m"`       //
	NcharM          string    `json:"nchar_m"`          //
	TimeM           string    `json:"time_m"`           //
	DateM           time.Time `json:"date_m"`           //
	DataTimeM       time.Time `json:"data_time_m"`      //
	TimestampM      time.Time `json:"timestamp_m"`      // 创建时间
	TimestampUpdate time.Time `json:"timestamp_update"` // 更新时间
	YearM           string    `json:"year_m"`           // 年
	TText           string    `json:"t_text"`           //
	MText           string    `json:"m_text"`           //
	TextM           string    `json:"text_m"`           //
	LText           string    `json:"l_text"`           //
	BinaryM         []byte    `json:"binary_m"`         //
	BlobM           []byte    `json:"blob_m"`           //
	LBlob           []byte    `json:"l_blob"`           //
	MBlob           []byte    `json:"m_blob"`           //
	TBlob           []byte    `json:"t_blob"`           //
	BitM            []uint8   `json:"bit_m"`            //
	EnumM           string    `json:"enum_m"`           //
	SetM            string    `json:"set_m"`            //
	BoolM           int8      `json:"bool_m"`           //
}

const (
	table  = "`all_type_table`"
	fields = "`id`,`t_int`,`s_int`,`m_int`,`b_int`,`f32`,`f64`,`decimal_mysql`,`char_m`,`varchar_m`,`json_m`,`nvarchar_m`,`nchar_m`,`time_m`,`date_m`,`data_time_m`,`timestamp_m`,`timestamp_update`,`year_m`,`t_text`,`m_text`,`text_m`,`l_text`,`binary_m`,`blob_m`,`l_blob`,`m_blob`,`t_blob`,`bit_m`,`enum_m`,`set_m`,`bool_m`"
	//ID 自增id
	ID = "`id`"
	//TInt 小整型
	TInt = "`t_int`"
	//SInt
	SInt = "`s_int`"
	//MInt 中整数
	MInt = "`m_int`"
	//BInt
	BInt = "`b_int`"
	//F32
	F32 = "`f32`"
	//F64
	F64 = "`f64`"
	//DecimalMysql
	DecimalMysql = "`decimal_mysql`"
	//CharM
	CharM = "`char_m`"
	//VarcharM
	VarcharM = "`varchar_m`"
	//JSONM
	JSONM = "`json_m`"
	//NvarcharM
	NvarcharM = "`nvarchar_m`"
	//NcharM
	NcharM = "`nchar_m`"
	//TimeM
	TimeM = "`time_m`"
	//DateM
	DateM = "`date_m`"
	//DataTimeM
	DataTimeM = "`data_time_m`"
	//TimestampM 创建时间
	TimestampM = "`timestamp_m`"
	//TimestampUpdate 更新时间
	TimestampUpdate = "`timestamp_update`"
	//YearM 年
	YearM = "`year_m`"
	//TText
	TText = "`t_text`"
	//MText
	MText = "`m_text`"
	//TextM
	TextM = "`text_m`"
	//LText
	LText = "`l_text`"
	//BinaryM
	BinaryM = "`binary_m`"
	//BlobM
	BlobM = "`blob_m`"
	//LBlob
	LBlob = "`l_blob`"
	//MBlob
	MBlob = "`m_blob`"
	//TBlob
	TBlob = "`t_blob`"
	//BitM
	BitM = "`bit_m`"
	//EnumM
	EnumM = "`enum_m`"
	//SetM
	SetM = "`set_m`"
	//BoolM
	BoolM = "`bool_m`"
)

// Insert  Insert a record
func Insert(ctx context.Context, db crud.DB, a *AllTypeTable) error {
	const sqlstr = "INSERT INTO  `all_type_table` (" +
		" `t_int`,`s_int`,`m_int`,`b_int`,`f32`,`f64`,`decimal_mysql`,`char_m`,`varchar_m`,`json_m`,`nvarchar_m`,`nchar_m`,`time_m`,`date_m`,`data_time_m`,`year_m`,`t_text`,`m_text`,`text_m`,`l_text`,`binary_m`,`blob_m`,`l_blob`,`m_blob`,`t_blob`,`bit_m`,`enum_m`,`set_m`,`bool_m`" +
		`) VALUES (` +
		` ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?` +
		`)`

	result, err := db.ExecContext(ctx, sqlstr, &a.TInt, &a.SInt, &a.MInt, &a.BInt, &a.F32, &a.F64, &a.DecimalMysql, &a.CharM, &a.VarcharM, &a.JSONM, &a.NvarcharM, &a.NcharM, &a.TimeM, &a.DateM, &a.DataTimeM, &a.YearM, &a.TText, &a.MText, &a.TextM, &a.LText, &a.BinaryM, &a.BlobM, &a.LBlob, &a.MBlob, &a.TBlob, &a.BitM, &a.EnumM, &a.SetM, &a.BoolM)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	a.ID = uint32(id)

	return nil
}

// Delete Delete by primary key:`id`
func Delete(ctx context.Context, db crud.DB, id uint32) (int64, error) {
	const sqlstr = `DELETE FROM  all_type_table WHERE id = ?`
	result, err := db.ExecContext(ctx, sqlstr, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Where  create where builder for update or query
func Where() *crud.Where {
	return &crud.Where{}
}

// Query create query builder  maybe with where builder
func Query(where crud.WhereBuilder) crud.QueryBuilder {
	return &crud.Query{
		Table:     table,
		AllFields: fields,
		Where:     where,
	}
}

// FindOne find a record
func FindOne(ctx context.Context, db crud.DB, build crud.Builder) (*AllTypeTable, error) {
	sql, args := build.Build()
	return FindOneRaw(ctx, db, sql, args...)
}

// FindOneRaw find a record by raw sql
func FindOneRaw(ctx context.Context, db crud.DB, sql string, args ...interface{}) (*AllTypeTable, error) {
	a := AllTypeTable{}
	err := db.QueryRowContext(ctx, sql, args...).Scan(&a.ID, &a.TInt, &a.SInt, &a.MInt, &a.BInt, &a.F32, &a.F64, &a.DecimalMysql, &a.CharM, &a.VarcharM, &a.JSONM, &a.NvarcharM, &a.NcharM, &a.TimeM, &a.DateM, &a.DataTimeM, &a.TimestampM, &a.TimestampUpdate, &a.YearM, &a.TText, &a.MText, &a.TextM, &a.LText, &a.BinaryM, &a.BlobM, &a.LBlob, &a.MBlob, &a.TBlob, &a.BitM, &a.EnumM, &a.SetM, &a.BoolM)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// Find Find many record by where statment or query statment
func Find(ctx context.Context, db crud.DB, build crud.Builder) ([]*AllTypeTable, error) {
	sql, args := build.Build()
	return FindRaw(ctx, db, sql, args...)
}

// FindRaw FindRaw many record by raw sql
func FindRaw(ctx context.Context, db crud.DB, sql string, args ...interface{}) ([]*AllTypeTable, error) {
	q, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	res := []*AllTypeTable{}
	for q.Next() {
		a := AllTypeTable{}
		err = q.Scan(&a.ID, &a.TInt, &a.SInt, &a.MInt, &a.BInt, &a.F32, &a.F64, &a.DecimalMysql, &a.CharM, &a.VarcharM, &a.JSONM, &a.NvarcharM, &a.NcharM, &a.TimeM, &a.DateM, &a.DataTimeM, &a.TimestampM, &a.TimestampUpdate, &a.YearM, &a.TText, &a.MText, &a.TextM, &a.LText, &a.BinaryM, &a.BlobM, &a.LBlob, &a.MBlob, &a.TBlob, &a.BitM, &a.EnumM, &a.SetM, &a.BoolM)
		if err != nil {
			return nil, err
		}
		res = append(res, &a)
	}
	if q.Err() != nil {
		return nil, err
	}

	return res, nil
}

// Count Count return number of rows that fit the where statment
func Count(ctx context.Context, db crud.DB, build crud.Builder) (int64, error) {
	sql, args := build.Build()
	var a int64
	if err := db.QueryRowContext(ctx, "SELECT COUNT(1) FROM `all_type_table` "+sql, args...).Scan(&a); err != nil {
		return 0, err
	}
	return a, nil
}

// UpdateBuilder UpdateBuilder
type UpdateBuilder interface {
	SetID(arg uint32) UpdateBuilder
	SetTInt(arg int8) UpdateBuilder
	SetSInt(arg uint16) UpdateBuilder
	SetMInt(arg int32) UpdateBuilder
	SetBInt(arg int64) UpdateBuilder
	SetF32(arg float32) UpdateBuilder
	SetF64(arg float64) UpdateBuilder
	SetDecimalMysql(arg float64) UpdateBuilder
	SetCharM(arg string) UpdateBuilder
	SetVarcharM(arg string) UpdateBuilder
	SetJSONM(arg string) UpdateBuilder
	SetNvarcharM(arg string) UpdateBuilder
	SetNcharM(arg string) UpdateBuilder
	SetTimeM(arg string) UpdateBuilder
	SetDateM(arg time.Time) UpdateBuilder
	SetDataTimeM(arg time.Time) UpdateBuilder
	SetTimestampM(arg time.Time) UpdateBuilder
	SetTimestampUpdate(arg time.Time) UpdateBuilder
	SetYearM(arg string) UpdateBuilder
	SetTText(arg string) UpdateBuilder
	SetMText(arg string) UpdateBuilder
	SetTextM(arg string) UpdateBuilder
	SetLText(arg string) UpdateBuilder
	SetBinaryM(arg []byte) UpdateBuilder
	SetBlobM(arg []byte) UpdateBuilder
	SetLBlob(arg []byte) UpdateBuilder
	SetMBlob(arg []byte) UpdateBuilder
	SetTBlob(arg []byte) UpdateBuilder
	SetBitM(arg []uint8) UpdateBuilder
	SetEnumM(arg string) UpdateBuilder
	SetSetM(arg string) UpdateBuilder
	SetBoolM(arg int8) UpdateBuilder
	Save(ctx context.Context, db crud.DB) (int64, error)
}

// Updater Updater
type Updater struct {
	where  crud.WhereBuilder
	Args   []interface{}
	Fields []string
}

// Update return a Updater
func Update(where crud.WhereBuilder) *Updater {
	return &Updater{
		where: where,
	}
}

// SetID set id
func (u *Updater) SetID(arg uint32) UpdateBuilder {
	u.Fields = append(u.Fields, "`id` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTInt set t_int
func (u *Updater) SetTInt(arg int8) UpdateBuilder {
	u.Fields = append(u.Fields, "`t_int` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetSInt set s_int
func (u *Updater) SetSInt(arg uint16) UpdateBuilder {
	u.Fields = append(u.Fields, "`s_int` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetMInt set m_int
func (u *Updater) SetMInt(arg int32) UpdateBuilder {
	u.Fields = append(u.Fields, "`m_int` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetBInt set b_int
func (u *Updater) SetBInt(arg int64) UpdateBuilder {
	u.Fields = append(u.Fields, "`b_int` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetF32 set f32
func (u *Updater) SetF32(arg float32) UpdateBuilder {
	u.Fields = append(u.Fields, "`f32` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetF64 set f64
func (u *Updater) SetF64(arg float64) UpdateBuilder {
	u.Fields = append(u.Fields, "`f64` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetDecimalMysql set decimal_mysql
func (u *Updater) SetDecimalMysql(arg float64) UpdateBuilder {
	u.Fields = append(u.Fields, "`decimal_mysql` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetCharM set char_m
func (u *Updater) SetCharM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`char_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetVarcharM set varchar_m
func (u *Updater) SetVarcharM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`varchar_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetJSONM set json_m
func (u *Updater) SetJSONM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`json_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetNvarcharM set nvarchar_m
func (u *Updater) SetNvarcharM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`nvarchar_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetNcharM set nchar_m
func (u *Updater) SetNcharM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`nchar_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTimeM set time_m
func (u *Updater) SetTimeM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`time_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetDateM set date_m
func (u *Updater) SetDateM(arg time.Time) UpdateBuilder {
	u.Fields = append(u.Fields, "`date_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetDataTimeM set data_time_m
func (u *Updater) SetDataTimeM(arg time.Time) UpdateBuilder {
	u.Fields = append(u.Fields, "`data_time_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTimestampM set timestamp_m
func (u *Updater) SetTimestampM(arg time.Time) UpdateBuilder {
	u.Fields = append(u.Fields, "`timestamp_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTimestampUpdate set timestamp_update
func (u *Updater) SetTimestampUpdate(arg time.Time) UpdateBuilder {
	u.Fields = append(u.Fields, "`timestamp_update` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetYearM set year_m
func (u *Updater) SetYearM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`year_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTText set t_text
func (u *Updater) SetTText(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`t_text` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetMText set m_text
func (u *Updater) SetMText(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`m_text` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTextM set text_m
func (u *Updater) SetTextM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`text_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetLText set l_text
func (u *Updater) SetLText(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`l_text` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetBinaryM set binary_m
func (u *Updater) SetBinaryM(arg []byte) UpdateBuilder {
	u.Fields = append(u.Fields, "`binary_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetBlobM set blob_m
func (u *Updater) SetBlobM(arg []byte) UpdateBuilder {
	u.Fields = append(u.Fields, "`blob_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetLBlob set l_blob
func (u *Updater) SetLBlob(arg []byte) UpdateBuilder {
	u.Fields = append(u.Fields, "`l_blob` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetMBlob set m_blob
func (u *Updater) SetMBlob(arg []byte) UpdateBuilder {
	u.Fields = append(u.Fields, "`m_blob` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetTBlob set t_blob
func (u *Updater) SetTBlob(arg []byte) UpdateBuilder {
	u.Fields = append(u.Fields, "`t_blob` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetBitM set bit_m
func (u *Updater) SetBitM(arg []uint8) UpdateBuilder {
	u.Fields = append(u.Fields, "`bit_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetEnumM set enum_m
func (u *Updater) SetEnumM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`enum_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetSetM set set_m
func (u *Updater) SetSetM(arg string) UpdateBuilder {
	u.Fields = append(u.Fields, "`set_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// SetBoolM set bool_m
func (u *Updater) SetBoolM(arg int8) UpdateBuilder {
	u.Fields = append(u.Fields, "`bool_m` = ? ")
	u.Args = append(u.Args, arg)
	return u
}

// Save do a update statment
func (u *Updater) Save(ctx context.Context, db crud.DB) (int64, error) {
	var where string
	var args []interface{}
	if u.where != nil {
		where, args = u.where.Build()
	}
	if len(u.Fields) <= 0 {
		return 0, errors.New("not set update fields")
	}

	var b bytes.Buffer
	b.WriteString("UPDATE ")
	b.WriteString(table)
	b.WriteString(" SET ")
	b.WriteString(strings.Join(u.Fields, ","))
	b.WriteString(where)
	u.Args = append(u.Args, args...)
	result, err := db.ExecContext(ctx, b.String(), u.Args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

