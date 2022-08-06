package translator

import (
	"go/token"
	"modernc.org/cc/v3"
)

type TranslatedUnit struct {
	DeclarationStructs []DeclarationStruct
}

type DeclarationStruct struct {
	Type       string
	Identifier string
	Childs     []DeclarationType
}

type DeclarationType struct {
	IsStruct          bool
	DeclarationStruct DeclarationStruct
	Type              string
	Identifier        string
	PositionLine      int
}

func New(t *cc.TranslationUnit) []TranslatedUnit {
	var tus []TranslatedUnit
	// Iterate and Find all struct header
	for t != nil {
		var tu TranslatedUnit
		it := filterExternalDeclaration(t.ExternalDeclaration)
		tu.DeclarationStructs = it
		tus = append(tus, tu)
		t = t.TranslationUnit
	}

	return tus
}

func filterExternalDeclaration(e *cc.ExternalDeclaration) []DeclarationStruct {
	// Declaration Filter
	if e.Case == 1 {
		a := e.Declaration.DeclarationSpecifiers

		var allStructs []DeclarationStruct
		for a != nil {
			single := filterExternalTypeSpecifier(a.TypeSpecifier)

			allStructs = append(allStructs, single)
			a = a.DeclarationSpecifiers
		}

		return allStructs
	}

	return []DeclarationStruct{}
}

func filterExternalTypeSpecifier(e *cc.TypeSpecifier) DeclarationStruct {
	a := e.StructOrUnionSpecifier

	c := filterExternalStructOrUnionSpecifier(a)

	return c

}

func filterExternalStructOrUnionSpecifier(e *cc.StructOrUnionSpecifier) DeclarationStruct {
	a := e.StructDeclarationList

	var structInit DeclarationStruct

	if e.StructOrUnion.Token.String() == cc.Struct.String() {
		structInit.Identifier = e.Token.String()
	}

	for a != nil {
		sd := a.StructDeclaration
		dt := appendTypeSpecifier(sd.SpecifierQualifierList.TypeSpecifier)
		ident := appendStructDeclarator(sd.StructDeclaratorList.StructDeclarator)
		dt.Identifier = ident
		structInit.Childs = append(structInit.Childs, dt)
		a = a.StructDeclarationList

	}
	return structInit
}

func appendStructDeclarator(sd *cc.StructDeclarator) string {
	if sd.Declarator.DirectDeclarator.DirectDeclarator != nil {
		return sd.Declarator.DirectDeclarator.DirectDeclarator.Token.String()
	} else {
		return sd.Declarator.DirectDeclarator.Token.String()
	}
}

func appendTypeSpecifier(ts *cc.TypeSpecifier) DeclarationType {
	var dt DeclarationType
	switch ts.Case {
	case cc.TypeSpecifierStructOrUnion:

		dt.IsStruct = true

		a := filterExternalStructOrUnionSpecifier(ts.StructOrUnionSpecifier)

		dt.DeclarationStruct = a
	default:

		dt.IsStruct = false
		dt.Type = ts.Token.String()
		dt.PositionLine = token.Position(ts.Position()).Line

	}

	return dt
}
