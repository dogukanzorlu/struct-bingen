package generator

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"struct-bingen/pkg/clang/translator"
)

func New(st []translator.StructType, file string) {

	stToString := "package main;"
	for _, structType := range st {
		stIdent := fmt.Sprintf("type %s struct {", structType.Ident)
		for _, elem := range structType.Elems {
			stIdent += fmt.Sprintf("%s %s;", elem.Ident, elem.Type)
		}
		stToString += stIdent + "};"
	}

	fmtdata, err := format.Source([]byte(stToString))
	fmt.Println(err)
	_ = ioutil.WriteFile("/Users/dogukanzorlu/Projects/go-projects/struct-bingen/testdata/dummy.go", fmtdata, 0644)
}
