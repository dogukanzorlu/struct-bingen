package translator

import (
	"go/ast"
)

func SetField(ident []*ast.Ident, typ ast.Expr) *ast.Field {
	field := ast.Field{Names: ident, Type: typ}

	return &field
}

func SetIdent(obj *ast.Object, name string) *ast.Ident {
	ident := ast.Ident{Name: name, Obj: obj}

	return &ident
}

func SetObject(kind ast.ObjKind, name string) *ast.Object {
	object := ast.Object{Kind: kind, Name: name}

	return &object
}
