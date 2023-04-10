package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"struct-bingen/pkg/clang"
	_ "struct-bingen/pkg/clang"
	"struct-bingen/pkg/clang/parser"
)

var (
	file string
)

var (
	source string
	target string
)

func main() {
	flag.StringVar(&source, "source", "", "source file for c binding")
	flag.StringVar(&target, "target", "", "target file for will be generate go file")
	flag.Parse()

	yamlCfg := readYaml()

	err := clang.Convert(yamlCfg, source, target)
	if err != nil {
		os.Exit(1)
	}

}

func readYaml() parser.PreProcessConfig {
	f, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var c parser.PreProcessConfig
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
