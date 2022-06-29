package db_mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MYSQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
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
	//m := make(map[string]CellData)
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

func (m *MYSQL) connect() (*sql.DB, error) {

	var Dsn string
	if len(m.Dbname) != 0 {
		Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Dbname)
	} else {
		Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)", m.Username, m.Password, m.Host, m.Port)
	}

	db, err := sql.Open("mysql", Dsn)

	db.SetConnMaxLifetime(100)

	db.SetMaxIdleConns(10)

	if err != nil {
		panic(err)
	}
	return db, nil
}

func (m *MYSQL) rowToStruct(query string) (resultData ResultData, columns []string, err error) {

	db, _ := m.connect()

	defer db.Close()

	var rows *sql.Rows

	rows, err = db.Query(query)

	if err != nil {
		panic(err)
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
func (m *MYSQL) SelectResToArray(query string) ([][]string, error) {

	var myres [][]string

	resultData, _, _ := m.rowToStruct(query)

	for _, row := range resultData {
		rowmap := rowToArry(row)
		myres = append(myres, rowmap)

	}
	return myres, nil
}

// SelectResToJson Format Result to Json
func (m *MYSQL) SelectResToJson(query string) ([]string, error) {

	var myres []string

	resultData, columns, _ := m.rowToStruct(query)

	for _, row := range resultData {
		rowmap := rowToMap(row, columns)
		result, _ := json.MarshalIndent(rowmap, "", "    ")
		myres = append(myres, string(result))

	}
	return myres, nil
}

// SelectResToMap Format Result to Map
func (m *MYSQL) SelectResToMap(query string) ([]interface{}, error) {
	var myres []interface{}
	resultData, columns, _ := m.rowToStruct(query)
	for _, row := range resultData {
		rowmap := rowToMap(row, columns)
		myres = append(myres, rowmap)
	}
	return myres, nil
}

// SelectToRowsData for show  engine innodb status
func (m *MYSQL) SelectToRowsData(query string) (NamedResultData, error) {

	resultData, columns, err := m.rowToStruct(query)

	return NamedResultData{Columns: columns, Data: resultData}, err

}

// DirectExec for DML and DDL
func (m *MYSQL) DirectExec(query string) (msg string, err error) {
	defer func() {
		if derr := recover(); derr != nil {
			err = errors.New(fmt.Sprintf("DirectExec unexpected error: %+v", derr))
		}
	}()
	db, _ := m.connect()
	defer db.Close()

	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	rownumaffected, _ := res.RowsAffected()
	resmsg := fmt.Sprintf("RowsAffected: %d", rownumaffected)
	return resmsg, err
}

// SingleTrxExec for TRX Exec
func (m *MYSQL) SingleTrxExec(query string) (msg string, err error) {
	defer func() {
		if derr := recover(); derr != nil {
			err = errors.New(fmt.Sprintf("SingleTrxExec unexpected error: %+v", derr))
		}
	}()
	db, _ := m.connect()

	defer db.Close()

	trx, err := db.Begin()

	if err != nil {
		panic(err)
	}
	res, err := trx.Exec(query)
	if err != nil {
		trx.Rollback()
		panic(err)
	}

	rownums, _ := res.RowsAffected()

	resmsg := fmt.Sprintf("Exec Succssfully! RowsAffected: %d", rownums)

	trx.Commit()
	return resmsg, err
}

// ComTrxExec for Combine TRX Exec
func (m *MYSQL) ComTrxExec(queryarry []string) (res []string, err error) {
	defer func() {
		if derr := recover(); derr != nil {
			err = errors.New(fmt.Sprintf("ComTrxExec unexpected error: %+v", derr))
		}
	}()

	db, _ := m.connect()

	defer db.Close()

	trx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	var resmsg []string
	for _, v := range queryarry {
		rowsres, err := trx.Exec(v)

		if err != nil {
			trx.Rollback()
			panic(err)
		}

		rownums, _ := rowsres.RowsAffected()
		msg := fmt.Sprintf("Exec Succssfully! RowsAffected: %d", rownums)
		resmsg = append(resmsg, msg)

	}

	trx.Commit()

	return resmsg, err

}
