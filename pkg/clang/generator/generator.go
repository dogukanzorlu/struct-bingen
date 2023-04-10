package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
)

func New(st []ast.Decl, file string) error {

	bbuf := bytes.NewBuffer(nil)
	err := format.Node(bbuf, token.NewFileSet(), &ast.File{Decls: st, Name: &ast.Ident{Name: "cpu"}})
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	source, err := format.Source(bbuf.Bytes())
	if err != nil {
		return err
	}
	errs := ioutil.WriteFile(file, source, 0644)
	if err != nil {
		return errs
	}

	return nil
}
