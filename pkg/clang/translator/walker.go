package translator

import (
	"github.com/iancoleman/strcase"
	"go/ast"
	"modernc.org/cc/v3"
)

func walkDecl(d *cc.Declaration) []ast.Spec {
	spec := d.DeclarationSpecifiers
	if spec.Case == cc.DeclarationSpecifiersStorage &&
		spec.StorageClassSpecifier.Case == cc.StorageClassSpecifierTypedef {
		spec = spec.DeclarationSpecifiers
	}
	var specs []ast.Spec
	for sp := spec; sp != nil; sp = sp.DeclarationSpecifiers {
		switch sp.Case {
		case cc.DeclarationSpecifiersTypeSpec:
			ds := sp.TypeSpecifier
			switch ds.Case {
			case cc.TypeSpecifierStructOrUnion:
				su := ds.StructOrUnionSpecifier
				switch su.Case {
				case cc.StructOrUnionSpecifierDef:
					name := su.Type().Name().String()
					if name == "" {
						name = d.InitDeclaratorList.InitDeclarator.Declarator.DirectDeclarator.Token.Value.String()
					}
					fields := convertStruct(su.Type())
					specs = append(specs, &ast.TypeSpec{Name: &ast.Ident{
						Name: strcase.ToCamel(name),
						Obj:  &ast.Object{Kind: ast.Typ, Name: strcase.ToCamel(name)},
					}, Type: &ast.StructType{Fields: &ast.FieldList{List: fields}}})
				default:
					continue
				}
			default:
				continue
			}
		default:
			continue
		}
	}

	return specs
}

func convertStruct(t cc.Type) []*ast.Field {
	var fields []*ast.Field
	for i := 0; i < t.NumField(); i++ {
		var idents []*ast.Ident
		f := t.FieldByIndex([]int{i})
		identObj := SetObject(ast.Var, strcase.ToCamel(f.Name().String()))
		idents = append(idents, SetIdent(identObj, strcase.ToCamel(f.Name().String())))

		field := detectKinds(f.Type())
		fields = append(fields, SetField(idents, field))
	}
	return fields
}
