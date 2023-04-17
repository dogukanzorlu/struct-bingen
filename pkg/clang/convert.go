package clang

import (
	"fmt"
	"struct-bingen/pkg/clang/generator"
	"struct-bingen/pkg/clang/parser"
	"struct-bingen/pkg/clang/translator"
)

func Convert(cfg parser.PreProcessConfig, source, target string) error {
	unit, err := parser.Parse(cfg, source)
	if err != nil {
		fmt.Println(err)
		return err
	}

	structs := translator.Translate(unit)
	if len(structs) == 0 {
		panic("Translate Error, no struct or struct could not be defined\n")
	}
	//output := render.AsCode(structs)
	//fmt.Println(output)
	err = generator.New(structs, target)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
