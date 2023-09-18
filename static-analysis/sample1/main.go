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
	f, err := parser.ParseFile(fset, "../../sample/sample1/main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					fmt.Printf("variable name: %s\n", valueSpec.Names[0].Name)
					fmt.Printf("initial value: %s\n", valueSpec.Values[0].(*ast.BasicLit).Value)
					fmt.Printf("position: %v\n", spec.Pos())
				}
			}
		}
	}
}
