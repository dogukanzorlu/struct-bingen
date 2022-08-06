package clang_test

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"struct-bingen/pkg/clang"
	"struct-bingen/pkg/clang/parser"
	"testing"
)

func TestConvert(t *testing.T) {

	config := readYaml()

	tranlationUnit, err := clang.Convert(config)

	r := fmt.Sprintf("%v", tranlationUnit)
	output := "[{[{ Student [{false {  []} char name 2} {false {  []} int age 3} {false {  []} int year 4} {false {  []} float gpa 5}]}]} {[{ Organisation [{false {  []} char organisation_name 10} {false {  []} char org_number 11} {true { Student []}  emp 0}]}]} {[{ mabbas [{false {  []} int mnum 21} {false {  []} char mlet 22}]}]}]"

	if r != output || err != nil {
		t.Fatalf(`Convert Clang expected: %s, but got: %s-%v`, output, r, err)
	}

	op := prettyPrint(tranlationUnit)

	fmt.Println(op)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func readYaml() parser.PreProcessConfig {
	absConfig, _ := filepath.Abs("../../config-test.yaml")
	f, err := os.ReadFile(absConfig)
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
