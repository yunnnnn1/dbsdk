
# dbsdk
![LICENSE](https://img.shields.io/badge/license-AGPL%20-blue.svg)
![](https://img.shields.io/badge/build-release-brightgreen.svg)  
![](https://img.shields.io/badge/version-v1.1.0-brightgreen.svg)

#  Installation

```go
$ go get -u github.com/jingmingyu/dbsdk 
```

# MYSQL sdk Usage: 
##### Init MySQL DB
```go
    mysqldb := db_mysql.MYSQL{
		Host:     "IP",
		Port:     "3306",
		Username: "xxxxxx",
		Password: "xxxxxx",
		Dbname:   "xxxxxxx",
	}
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
https://segmentfault.com/a/1190000020105590
https://www.cnblogs.com/taadis/p/12132809.html
```
### Select  Method:
 1. SelectToJson 
 2. SelectToRowsData

##### SelectToJson : Ordinary select  like select / show master status 
```go
	querysql := fmt.Sprintf("%s", "SELECT * FROM `dbtool_accesslog`")
	res, _ := mysqldb.SelectToJson(querysql)
	fmt.Println(res)
```
###### result :
```azure
[{
    "created_by": {
        "String": "",
        "Valid": false
    },
    "id": {
        "String": "10",
        "Valid": true
    },
    "operate_page": {
        "String": "zzzzzzzz",
        "Valid": true
    },
    "when_created": {
        "String": "2022-03-17T22:37:23+08:00",
        "Valid": true
    }
}]
```

#### SelectToRowsData:  show engine innodb status 
```go
	querysql := fmt.Sprintf("%s", "show engine innodb status")
	res, _ := mysqldb.SelectToRowsData(querysql)
	fmt.Println(res.Data)
```
###### res:
```azure
[[{InnoDB true} { true} {
=====================================
2022-03-18 10:04:00 140672645170944 INNODB MONITOR OUTPUT
=====================================
Per second averages calculated from the last 59 seconds
-----------------
BACKGROUND THREAD
-----------------
srv_master_thread loops: 5073576 srv_active, 0 srv_shutdown, 1710 srv_idle
srv_master_thread log flush and writes: 0
----------
...
```
### DML  Method:
 1. insert
 ```go
	insertsql := fmt.Sprintf("%s", "insert into `dbtool_accesslog`( `operate_page` , `when_created` ) VALUE ('zzzzzzzz',now())")
	res, _ := mysqldb.DirectExec(insertsql)
	fmt.Println(res)
```
 2. delete
```go
    deletesql := fmt.Sprintf("%s", "delete from `dbtool_accesslog`")
    res, _ := mysqldb.DirectExec(deletesql)
    fmt.Println(res)
```
 3. update
```go
	updatesql := fmt.Sprintf("%s", "update `dbtool_accesslog` set operate_page='ascd' ")
	res, _ := mysqldb.DirectExec(updatesql)
	fmt.Println(res)
```
### TRX Exec  Method:
 1. DirectExec
```go
	updatesql := fmt.Sprintf("%s", "update `dbtool_accesslog` set operate_page='ascd' ")
	res, _ := mysqldb.DirectExec(updatesql)
	fmt.Println(res)

	exec add/drop ddl
	addsql := fmt.Sprintf("%s", "alter table dbtool_accesslog add zzz int(2)")
	res, _ := mysqldb.DirectExec(addsql)
	fmt.Println(res)

	dropsql := fmt.Sprintf("%s", "alter table dbtool_accesslog drop  COLUMN zzz ")
	res, _ := mysqldb.DirectExec(dropsql)
	fmt.Println(res)
```
 2. SingleTrxExec
```go
	//single trx exec
	// 单个事务执行
	trxsql := fmt.Sprintf("%s", "insert into `dbtool_accesslog`( `operate_page` , `when_created` ) VALUE ('zzzzzzzz',now())")
	res, err := mysqldb.SingleTrxExec(trxsql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
```

 4. ComTrxExec
```go
	// combine trx exec
	// 组合事务顺序执行
	var trx []string
	trxsql1 := fmt.Sprintf("%s", "insert into `dbtool_accesslog`( `operate_page` , `when_created` ) VALUE ('zzzzzzzz',now())")
	trxsql2 := fmt.Sprintf("%s", "insert into `dbtool_accesslog`( `operate_page` , `when_created` ) VALUE ('bbbbb',now())")
	trxsql3 := fmt.Sprintf("%s", "insert into `dbtool_accesslog`( `operate_page` , `when_created` ) VALUE ('bbbbb',now())")
	trxsql4 := fmt.Sprintf("%s", "update `dbtool_accesslog` set operate_page='zzzzzzzz' where operate_page='ascd'")
	trxsql5 := fmt.Sprintf("%s", "insert into `dbtool_accesslog`( `operate_page` , `when_created` ) VALUE ('zzzzzzzz',now())")
	
	trx = append(trx,trxsql1,trxsql2,trxsql3,trxsql4,trxsql5)
	
	res, err := mysqldb.ComTrxExec(trx)
	
	if err != nil {
		fmt.Println(err)
	}
	
	for _,v := range res{
		fmt.Println(v)
	}
```
