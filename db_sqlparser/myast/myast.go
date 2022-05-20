package myast

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
)

// TiDBParse TiDB 语法解析
func TiDBParse(sql, charset, collation string) ([]ast.StmtNode, error) {
	p := parser.New()
	stmt, warn, err := p.Parse(sql, charset, collation)
	// TODO: bypass warning info
	for _, w := range warn {
		fmt.Println(w.Error())
	}
	return stmt, err
}

// PrintPrettyStmtNode 打印TiParse语法树
func PrintPrettyStmtNode(sql, charset, collation string) {
	tree, err := TiDBParse(sql, charset, collation)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = pretty.Println(tree)
		fmt.Println(err.Error())
	}
}

// StmtNode2JSON TiParse AST tree into json format
func StmtNode2JSON(sql, charset, collation string) string {
	var str string
	tree, err := TiDBParse(sql, charset, collation)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, err := json.MarshalIndent(tree[0], "", "  ")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			str = string(b)
		}
	}

	return str
}
