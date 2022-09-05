package main

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"struct-bingen/pkg/clang"
	_ "struct-bingen/pkg/clang"
	parser "struct-bingen/pkg/clang/parser"
)

var (
	file string
)

func main() {
	flag.StringVar(&file, "file", "", "target c file")

	flag.Parse()

	config := readYaml()

	err := clang.Convert(config, file)
	if err != nil {
		return
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
