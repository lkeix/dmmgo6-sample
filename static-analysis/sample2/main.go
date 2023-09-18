package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../sample/sample4/main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok {
			checkDeferCallClosure(f.Body)
		}
	}
}

func checkDeferCallClosure(body *ast.BlockStmt) {
	for _, stmt := range body.List {
		switch s := stmt.(type) {
		case *ast.DeferStmt:
			if _, ok := s.Call.Fun.(*ast.FuncLit); !ok {
				fmt.Println("defer call not closure")
				fmt.Println("position:", s.Pos())
			}
		}
	}
}
