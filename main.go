package main

import (
	"flag"
	"fmt"
	"struct-bingen/pkg/clang"
)

type sourcePathFlags []string

func (i *sourcePathFlags) String() string {
	return ""
}

func (i *sourcePathFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var sourcePaths sourcePathFlags

func main() {

	flag.Var(&sourcePaths, "sourcePath", "add your source path for code")

	flag.Parse()

	parsed, err := clang.Convert(sourcePaths)
	if err != nil {
		return
	}

	fmt.Println(parsed.Declarations, err)
}
