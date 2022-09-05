package clang

import (
	"fmt"
	"struct-bingen/pkg/clang/generator"
	"struct-bingen/pkg/clang/parser"
	"struct-bingen/pkg/clang/translator"
)

func Convert(cfg parser.PreProcessConfig, file string) error {

	unit, err := parser.Parse(cfg, file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	structs := translator.Translate(unit)
	if len(structs) == 0 {
		panic("Translate Error, no struct or struct could not be defined\n")
	}
	generator.New(structs, file)

	return nil
}
