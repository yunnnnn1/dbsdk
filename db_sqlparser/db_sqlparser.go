package db_sqlparser

import (
	"fmt"
	"github.com/jingmingyu/dbsdk/db_sqlparser/myast"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
)

type ParserResult struct {
	Method string
	Result string
}

func ChooseDoType(sql string) *ParserResult {
	p := parser.New()
	stmt, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		panic(err)
	}
	switch x := stmt.(type) {
	case *ast.InsertStmt:
		res := myast.StmtNode2JSON(sql, "", "")
		Result := ParserResult{
			Method: "INSERT",
			Result: res,
		}
		return &Result

	case *ast.DeleteStmt:
		res := myast.StmtNode2JSON(sql, "", "")
		Result := ParserResult{
			Method: "DELETE",
			Result: res,
		}
		return &Result
	case *ast.UpdateStmt:
		res := myast.StmtNode2JSON(sql, "", "")
		Result := ParserResult{
			Method: "UPDATE",
			Result: res,
		}
		return &Result
	case *ast.SelectStmt:
		res := myast.StmtNode2JSON(sql, "", "")
		Result := ParserResult{
			Method: "SELECT",
			Result: res,
		}
		return &Result
	default:
		fmt.Println("Now do not support this type", x.Text())
		return nil
	}

}
