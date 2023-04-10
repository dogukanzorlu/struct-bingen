package translator

import (
	"go/ast"
	"go/token"
	"modernc.org/cc/v3"
)

func Translate(rawAst *cc.AST) []ast.Decl {
	var resultDecl []ast.Decl
	tu := rawAst.TranslationUnit
	for tu != nil {
		d := tu.ExternalDeclaration
		tu = tu.TranslationUnit
		if d == nil {
			continue
		}
		switch d.Case {
		case cc.ExternalDeclarationDecl:
			sts := walkDecl(d.Declaration)
			if len(sts) > 0 {
				resultDecl = append(resultDecl, &ast.GenDecl{Specs: sts, Tok: token.TYPE})
			}
		default:
			continue
		}
	}

	return resultDecl
}
