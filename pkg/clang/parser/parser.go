package clang

import (
	"modernc.org/cc"
)

func ParseWith(includePaths []string, sourcesPaths []string, targetArch string) (*cc.TranslationUnit, error) {
	PredefinedVariables += targetArch

	return cc.Parse(PredefinedVariables, sourcesPaths, Model,
		cc.SysIncludePaths(includePaths),
		cc.EnableAnonymousStructFields(),
		cc.EnableAsm(),
		cc.EnableAlternateKeywords(),
		cc.EnableIncludeNext(),
		cc.EnableNoreturn(),
		cc.EnableEmptyDeclarations(),
		cc.EnableWideEnumValues(),
		cc.EnableWideBitFieldTypes(),
		cc.EnableParenthesizedCompoundStatemen(),
		cc.AllowCompatibleTypedefRedefinitions(),
	)
}
