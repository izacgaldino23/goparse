package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
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
		erro error
		b    bytes.Buffer
		// funcoes = make([]funcao, 0)
	)

	fset := token.NewFileSet()
	filters := AppendFilters(PosFilter, KeywordFilter, ZeroFilter)

	expr, erro := parser.ParseExpr(codigo)
	if erro == nil {
		erro = Fprint(&b, fset, expr, filters)
		// fmt.Println(erro)
		// fmt.Println("teste")
		return b.String()
	}

	f, erro := parser.ParseFile(fset, "", codigo, parser.ParseComments)
	if erro != nil {
		// fmt.Println(erro)
		// panic(erro)
	}

	erro = Fprint(&b, fset, f.Decls, filters)
	if erro != nil {
		// fmt.Println(erro)
		// panic(erro)
	}

	estrutura := new(FinalStructure)

	estrutura.Generate(f.Decls, filters)

	// v, erro := tahwil.ToValue(f)
	// if erro != nil {
	// 	fmt.Println(erro)
	// 	panic(erro)
	// }

	res, err := json.Marshal(estrutura.corpo)
	if err != nil {
		panic(err)
	}

	// fmt.Println(b.String())
	fmt.Println(string(res))

	// fmt.Println(json.Marshal(f))
	// fmt.Println(b.String())

	// var temp interface{}
	// if erro = json.Unmarshal(b.Bytes(), &temp); erro != nil {
	// 	fmt.Println(erro)
	// 	panic(erro)
	// }
	// fmt.Println(temp)

	return string(res)
}

// Parse asdfsadf asdf
func Parse(codigo string) interface{} {
	return formatar(codigo)
}

func main() {
	Parse(`package main

	import "fmt"

	func foo(name, address string, idade int64) error {
		fmt.Println("Hello, World!")
		return nil
	}`)
	// js.Module.Get("exports").Set("Parse", Parse)
}
