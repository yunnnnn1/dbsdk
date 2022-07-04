package db_oracle

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	_ "github.com/sijms/go-ora/v2"
)

type OracleInterface interface {
	SelectResToArray(query string) ([][]string, error)
	SelectResToJson(query string) ([]string, error)
	SelectResToMap(query string) ([]interface{}, error)
	DirectExec(query string) (msg string, err error)
	SingleTrxExec(query string) (msg string, err error)
}

type ORACLE struct {
	Host        string
	Port        string
	Username    string
	Password    string
	ServiceName string
}

type CellData sql.NullString

type RowMap map[string]CellData

type RowData []CellData

type ResultData []RowData

type NamedResultData struct {
	Columns []string
	Data    ResultData
}

func rowToMap(row []CellData, columns []string) map[string]string {
	m := make(map[string]string)
	for k, data_col := range row {
		m[columns[k]] = data_col.String
	}
	return m
}

func rowToArry(row []CellData) []string {
	var s []string
	for _, dataCol := range row {
		s = append(s, dataCol.String)
	}
	return s
}

func (c *CellData) NullString() *sql.NullString {
	return (*sql.NullString)(c)
}

func RowToArray(rows *sql.Rows, columns []string) []CellData {
	buff := make([]interface{}, len(columns))
	data := make([]CellData, len(columns))
	for i, _ := range buff {
		buff[i] = data[i].NullString()
	}
	rows.Scan(buff...)
	return data
}

func ScanRowsToArrays(rows *sql.Rows, on_row func([]CellData) error) error {
	columns, _ := rows.Columns()
	for rows.Next() {
		arr := RowToArray(rows, columns)
		err := on_row(arr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *ORACLE) String() string {
	return fmt.Sprintf("%s:%s/%s", o.Host, o.Port, o.ServiceName)
}

func (o *ORACLE) connect() (*sql.DB, error) {
	var Dsn string
	if o.ServiceName == "" {
		o.ServiceName = "xe"
	}

	if passwd, err := base64.StdEncoding.DecodeString(o.Password); err != nil {
		Dsn = fmt.Sprintf("oracle://%s:%s@%s:%s/%s", o.Username, passwd, o.Host, o.Port, o.ServiceName)
	} else {
		Dsn = fmt.Sprintf("oracle://%s:%s@%s:%s/%s", o.Username, o.Password, o.Host, o.Port, o.ServiceName)
	}

	db, err := sql.Open("oracle", Dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to %s", o.String())
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ping %s", o.String())
	}

	return db, nil
}

func (o *ORACLE) rowToStruct(query string) (resultData ResultData, columns []string, err error) {

	db, _ := o.connect()

	defer db.Close()

	var rows *sql.Rows

	rows, err = db.Query(query)

	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to query %s", query)
	}

	columns, _ = rows.Columns()

	resultData = ResultData{}

	err = ScanRowsToArrays(rows, func(rowData []CellData) error {
		resultData = append(resultData, rowData)
		return nil
	})

	return resultData, columns, err

}

// SelectResToArray  Format Result to Array [[1,2,3],[4,5,6]]
func (o *ORACLE) SelectResToArray(query string) ([][]string, error) {

	var myres [][]string

	resultData, _, _ := o.rowToStruct(query)

	for _, row := range resultData {
		rowmap := rowToArry(row)
		myres = append(myres, rowmap)

	}
	return myres, nil
}

// SelectResToJson Format Result to Json
func (o *ORACLE) SelectResToJson(query string) ([]string, error) {

	var myres []string

	resultData, columns, _ := o.rowToStruct(query)

	for _, row := range resultData {
		rowmap := rowToMap(row, columns)
		result, _ := json.MarshalIndent(rowmap, "", "    ")
		myres = append(myres, string(result))

	}
	return myres, nil

}

// SelectResToMap Format Result to Map
func (o *ORACLE) SelectResToMap(query string) ([]interface{}, error) {
	var myres []interface{}
	resultData, columns, _ := o.rowToStruct(query)
	for _, row := range resultData {
		rowmap := rowToMap(row, columns)
		myres = append(myres, rowmap)

	}
	return myres, nil
}

// DirectExec for DML and DDL
func (o *ORACLE) DirectExec(query string) (msg string, err error) {
	defer func() {
		if derr := recover(); derr != nil {
			err = errors.New(fmt.Sprintf("DirectExec unexpected error: %+v", derr))
		}
	}()
	db, _ := o.connect()
	defer db.Close()

	res, err := db.Exec(query)
	if err != nil {
		return "", errors.Wrapf(err, "failed to exec %s", query)
	}
	rownumaffected, _ := res.RowsAffected()
	resmsg := fmt.Sprintf("RowsAffected: %d", rownumaffected)
	return resmsg, err
}

// SingleTrxExec for TRX Exec
func (o *ORACLE) SingleTrxExec(query string) (msg string, err error) {
	defer func() {
		if derr := recover(); derr != nil {
			err = errors.New(fmt.Sprintf("SingleTrxExec unexpected error: %+v", derr))
		}
	}()
	db, _ := o.connect()

	defer db.Close()

	trx, err := db.Begin()

	if err != nil {
		return "", errors.Wrapf(err, "failed to begin trx")
	}
	res, err := trx.Exec(query)
	if err != nil {
		trx.Rollback()
		return "", errors.Wrapf(err, "failed to exec %s", query)
	}

	rownums, _ := res.RowsAffected()

	resmsg := fmt.Sprintf("Exec Succssfully! RowsAffected: %d", rownums)

	trx.Commit()
	return resmsg, err
}
