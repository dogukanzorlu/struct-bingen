package clang

import (
	"modernc.org/cc"
	clang "struct-bingen/pkg/clang/parser"
)

func Convert(sourcesPath []string) (*cc.TranslationUnit, error) {
	includePaths := []string{"/usr/lib/gcc/x86_64-linux-gnu/9/include",
		"/usr/local/include",
		"/usr/include/x86_64-linux-gnu",
		"/usr/include"}

	parser, err := clang.ParseWith(includePaths, sourcesPath, "#define __x86_64__ 1")
	if err != nil {
		return nil, err
	}

	return parser, nil
}
