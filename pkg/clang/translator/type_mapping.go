package translator

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/token"
	"modernc.org/cc/v3"
	"reflect"
)

func detectKinds(t cc.Type) ast.Expr {
	switch kind := t.Kind(); kind {
	case cc.UInt64, cc.UInt32, cc.UInt16, cc.UInt8:
		detectedKind := UintT(int(t.Size()))
		return &ast.Ident{Name: detectedKind.String()}
	case cc.Int64, cc.Int32, cc.Int16, cc.Int8, cc.Int:
		detectedKind := IntT(int(t.Size()))
		return &ast.Ident{Name: detectedKind.String()}
	case cc.SChar, cc.Short, cc.Long, cc.LongLong, cc.Char:
		detectedKind := IntT(int(t.Size()))
		return &ast.Ident{Name: detectedKind.String()}
	case cc.UChar, cc.UShort, cc.UInt, cc.ULong, cc.ULongLong:
		detectedKind := UintT(int(t.Size()))
		return &ast.Ident{Name: detectedKind.String()}
	case cc.Float, cc.Double, cc.LongDouble:
		detectedKind := FloatT(int(t.Size()))
		return &ast.Ident{Name: detectedKind.String()}
	case cc.Bool:
		detectedKind := reflect.TypeOf(true).Kind()
		return &ast.Ident{Name: detectedKind.String()}
	case cc.Ptr:
		detectedKind := detectKinds(t.Elem())
		return &ast.StarExpr{X: detectedKind}
	case cc.Array:
		detectedKind := detectKinds(t.Elem())
		length := fmt.Sprintf("%d", t.Len())
		return &ast.ArrayType{Len: &ast.BasicLit{Kind: token.INT, Value: length}, Elt: detectedKind}
	case cc.Struct, cc.Union:
		return &ast.Ident{Name: strcase.ToCamel(t.String())}
	default:
		panic(fmt.Errorf("%T, %s (%s)", t, kind, t.String()))
	}
}
