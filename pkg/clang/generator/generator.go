package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func New(st []ast.Decl, file string) error {

	bbuf := bytes.NewBuffer(nil)
	fset := token.NewFileSet()

	p, err := filepath.Abs(file)
	pkgName := strings.Trim(filepath.Base(p), ".go")
	err = format.Node(bbuf, fset, &ast.File{Decls: st, Name: &ast.Ident{Name: pkgName}})
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
