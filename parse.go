package main

import (
	"bytes"
	"encoding/json"
	"go/parser"
	"go/token"

	"github.com/gopherjs/gopherjs/js"
)

type funcao struct {
	nome       string
	retornos   []parametro
	parametros []parametro
}

type parametro struct {
	nome string
	tipo string
}

func formatar(codigo string) interface{} {
	var (
		err error
		b   bytes.Buffer
	)

	fset := token.NewFileSet()
	filters := AppendFilters(PosFilter, KeywordFilter, ZeroFilter)

	expr, err := parser.ParseExpr(codigo)
	if err == nil {
		err = Fprint(&b, fset, expr, filters)
		// fmt.Println(err)
		// fmt.Println("teste")
		return b.String()
	}

	f, err := parser.ParseFile(fset, "", codigo, parser.ParseComments)
	if err != nil {
		// fmt.Println(err)
		// panic(err)
	}

	// For retrieve all struct parse f variable only
	err = Fprint(&b, fset, f.Decls, filters)
	if err != nil {
		// fmt.Println(err)
		// panic(err)
	}

	estrutura := new(FinalStructure)

	estrutura.Generate(f.Decls, filters)

	res, err := json.Marshal(estrutura.body)
	if err != nil {
		// fmt.Println(err)
		// panic(err)
	}

	// fmt.Println(b.String())
	// fmt.Println(string(res))

	return string(res)
}

// Parse asdfsadf asdf
func Parse(codigo string) interface{} {
	return formatar(codigo)
}

func main() {
	// Parse(`package main

	// import "fmt"

	// func foo(name, address string, idade int64) error {
	// 	fmt.Println("Hello, World!")
	// 	return nil
	// }`)
	js.Module.Get("exports").Set("Parse", Parse)
}
