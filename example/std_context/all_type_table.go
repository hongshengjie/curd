package model

import (
	"context"
	"database/sql"
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

// Insert  Insert a record
func (a *AllTypeTable) Insert(ctx context.Context, db *sql.DB) error {
	var err error

	const sqlstr = `INSERT INTO  all_type_table (` +
		` t_int,s_int,m_int,b_int,f32,f64,decimal_mysql,char_m,varchar_m,json_m,nvarchar_m,nchar_m,time_m,date_m,data_time_m,year_m,t_text,m_text,text_m,l_text,binary_m,blob_m,l_blob,m_blob,t_blob,bit_m,enum_m,set_m,bool_m` +
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

// AllTypeTableDelete Delete by primary key:`id`
func AllTypeTableDelete(ctx context.Context, db *sql.DB, id uint32) error {
	var err error

	const sqlstr = `DELETE FROM  all_type_table WHERE id = ?`

	_, err = db.ExecContext(ctx, sqlstr, id)
	if err != nil {
		return err
	}
	return nil
}

// AllTypeTableByID   Select a record by primary key:`id`
func AllTypeTableByID(ctx context.Context, db *sql.DB, id uint32) (*AllTypeTable, error) {
	var err error

	const sqlstr = `SELECT ` +
		`id,t_int,s_int,m_int,b_int,f32,f64,decimal_mysql,char_m,varchar_m,json_m,nvarchar_m,nchar_m,time_m,date_m,data_time_m,timestamp_m,timestamp_update,year_m,t_text,m_text,text_m,l_text,binary_m,blob_m,l_blob,m_blob,t_blob,bit_m,enum_m,set_m,bool_m ` +
		`FROM  all_type_table ` +
		`WHERE id = ?`

	a := AllTypeTable{}

	err = db.QueryRowContext(ctx, sqlstr, id).Scan(&a.ID, &a.TInt, &a.SInt, &a.MInt, &a.BInt, &a.F32, &a.F64, &a.DecimalMysql, &a.CharM, &a.VarcharM, &a.JSONM, &a.NvarcharM, &a.NcharM, &a.TimeM, &a.DateM, &a.DataTimeM, &a.TimestampM, &a.TimestampUpdate, &a.YearM, &a.TText, &a.MText, &a.TextM, &a.LText, &a.BinaryM, &a.BlobM, &a.LBlob, &a.MBlob, &a.TBlob, &a.BitM, &a.EnumM, &a.SetM, &a.BoolM)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// Update Update a record
func (a *AllTypeTable) Update(ctx context.Context, db *sql.DB) error {
	var err error

	const sqlstr = `UPDATE all_type_table SET ` +
		` t_int = ? ,s_int = ? ,m_int = ? ,b_int = ? ,f32 = ? ,f64 = ? ,decimal_mysql = ? ,char_m = ? ,varchar_m = ? ,json_m = ? ,nvarchar_m = ? ,nchar_m = ? ,time_m = ? ,date_m = ? ,data_time_m = ? ,year_m = ? ,t_text = ? ,m_text = ? ,text_m = ? ,l_text = ? ,binary_m = ? ,blob_m = ? ,l_blob = ? ,m_blob = ? ,t_blob = ? ,bit_m = ? ,enum_m = ? ,set_m = ? ,bool_m = ? ` +
		` WHERE id = ?`

	_, err = db.ExecContext(ctx, sqlstr, a.TInt, a.SInt, a.MInt, a.BInt, a.F32, a.F64, a.DecimalMysql, a.CharM, a.VarcharM, a.JSONM, a.NvarcharM, a.NcharM, a.TimeM, a.DateM, a.DataTimeM, a.YearM, a.TText, a.MText, a.TextM, a.LText, a.BinaryM, a.BlobM, a.LBlob, a.MBlob, a.TBlob, a.BitM, a.EnumM, a.SetM, a.BoolM, a.ID)
	return err
}

// AllTypeTableByBIntCharM Select by index name:`ix_ccc`
func AllTypeTableByBIntCharM(ctx context.Context, db *sql.DB, b_int int64, char_m string) (*AllTypeTable, error) {

	const sqlstr = `SELECT ` +
		`id,t_int,s_int,m_int,b_int,f32,f64,decimal_mysql,char_m,varchar_m,json_m,nvarchar_m,nchar_m,time_m,date_m,data_time_m,timestamp_m,timestamp_update,year_m,t_text,m_text,text_m,l_text,binary_m,blob_m,l_blob,m_blob,t_blob,bit_m,enum_m,set_m,bool_m ` +
		`FROM  all_type_table ` +
		`WHERE b_int = ? AND char_m = ? `

	a := AllTypeTable{}

	err := db.QueryRowContext(ctx, sqlstr, b_int, char_m).Scan(&a.ID, &a.TInt, &a.SInt, &a.MInt, &a.BInt, &a.F32, &a.F64, &a.DecimalMysql, &a.CharM, &a.VarcharM, &a.JSONM, &a.NvarcharM, &a.NcharM, &a.TimeM, &a.DateM, &a.DataTimeM, &a.TimestampM, &a.TimestampUpdate, &a.YearM, &a.TText, &a.MText, &a.TextM, &a.LText, &a.BinaryM, &a.BlobM, &a.LBlob, &a.MBlob, &a.TBlob, &a.BitM, &a.EnumM, &a.SetM, &a.BoolM)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// AllTypeTableByMIntVarcharMTimestampM Select by index name:`ix_ddd`
func AllTypeTableByMIntVarcharMTimestampM(ctx context.Context, db *sql.DB, m_int int32, varchar_m string, timestamp_m time.Time) ([]*AllTypeTable, error) {

	const sqlstr = `SELECT ` +
		`id,t_int,s_int,m_int,b_int,f32,f64,decimal_mysql,char_m,varchar_m,json_m,nvarchar_m,nchar_m,time_m,date_m,data_time_m,timestamp_m,timestamp_update,year_m,t_text,m_text,text_m,l_text,binary_m,blob_m,l_blob,m_blob,t_blob,bit_m,enum_m,set_m,bool_m ` +
		`FROM  all_type_table ` +
		`WHERE m_int = ? AND varchar_m = ? AND timestamp_m = ? `
	q, err := db.QueryContext(ctx, sqlstr, m_int, varchar_m, timestamp_m)
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

// AllTypeTableByTIntSInt Select by index name:`uk_uuu`
func AllTypeTableByTIntSInt(ctx context.Context, db *sql.DB, t_int int8, s_int uint16) ([]*AllTypeTable, error) {

	const sqlstr = `SELECT ` +
		`id,t_int,s_int,m_int,b_int,f32,f64,decimal_mysql,char_m,varchar_m,json_m,nvarchar_m,nchar_m,time_m,date_m,data_time_m,timestamp_m,timestamp_update,year_m,t_text,m_text,text_m,l_text,binary_m,blob_m,l_blob,m_blob,t_blob,bit_m,enum_m,set_m,bool_m ` +
		`FROM  all_type_table ` +
		`WHERE t_int = ? AND s_int = ? `
	q, err := db.QueryContext(ctx, sqlstr, t_int, s_int)
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
