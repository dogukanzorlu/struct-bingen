package clang

import (
	"struct-bingen/pkg/clang/parser"
	"struct-bingen/pkg/clang/translator"
)

func Convert(cfg parser.PreProcessConfig) ([]translator.TranslatedUnit, error) {

	unit, err := parser.Parse(cfg)
	if err != nil {
		return nil, err
	}

	t := translator.New(unit.TranslationUnit)

	return t, nil
}
