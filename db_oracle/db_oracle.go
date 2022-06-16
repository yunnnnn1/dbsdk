package db_oracle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	_ "github.com/sijms/go-ora/v2"
)

type ORACLE struct {
	Host        string
	Port        string
	User        string
	Password    string
	ServiceName string
}


type BaseInfo struct {
	String string
	Vaild  string
}

type CellData sql.NullString

type RowMap map[string]CellData

type RowData []CellData

type ResultData []RowData

type NamedResultData struct {
	Columns []string
	Data    ResultData
}

func rowToMap(row []CellData, columns []string) map[string]CellData {
	m := make(map[string]CellData)
	for k, data_col := range row {
		m[columns[k]] = data_col
	}
	return m
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

	if o.ServiceName == "" {
		o.ServiceName = "xe"
	}

	db, err := sql.Open("oracle", fmt.Sprintf("oracle://%s:%s@%s:%s/%s", o.User, o.Password, o.Host, o.Port, o.ServiceName))
	if err != nil {
		return nil, errors.Wrapf(err,"failed to connect to %s", o.String())
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

// for Select
func (o *ORACLE) SelectToJson(query string) ([]string, error) {

	var myres []string

	resultData, columns, _ := o.rowToStruct(query)

	for _, row := range resultData {
		rowmap := rowToMap(row, columns)
		result, _ := json.MarshalIndent(rowmap, "", "    ")
		myres = append(myres, string(result))

	}
	return myres, nil

}


// for show  engine innodb status
func (o *ORACLE) SelectToRowsData(query string) (NamedResultData, error) {

	resultData, columns, err := o.rowToStruct(query)

	return NamedResultData{Columns: columns, Data: resultData}, err

}
