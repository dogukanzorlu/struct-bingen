package translator

import (
	"fmt"
	"modernc.org/cc/v3"
)

func Translate(ast *cc.AST) []StructType {
	var resultStruct []StructType
	tu := ast.TranslationUnit
	for tu != nil {
		d := tu.ExternalDeclaration
		tu = tu.TranslationUnit
		if d == nil {
			continue
		}
		switch d.Case {
		case cc.ExternalDeclarationDecl:
			sts := convertDecl(d.Declaration)
			if sts.Ident != "" {
				resultStruct = append(resultStruct, sts)
			}
		default:
			continue
		}
	}

	return resultStruct
}

func convertDecl(d *cc.Declaration) StructType {
	spec := d.DeclarationSpecifiers
	if spec.Case == cc.DeclarationSpecifiersStorage &&
		spec.StorageClassSpecifier.Case == cc.StorageClassSpecifierTypedef {
		spec = spec.DeclarationSpecifiers
	}

	var sts StructType
	for sp := spec; sp != nil; sp = sp.DeclarationSpecifiers {
		switch sp.Case {
		case cc.DeclarationSpecifiersTypeSpec:
			ds := sp.TypeSpecifier
			switch ds.Case {
			case cc.TypeSpecifierStructOrUnion:
				su := ds.StructOrUnionSpecifier
				switch su.Case {
				case cc.StructOrUnionSpecifierDef:
					fmt.Println(su.Position().String())
					sts = convertStructElemType(su.Type())
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

	return sts
}
