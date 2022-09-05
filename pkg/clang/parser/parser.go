package parser

import (
	"modernc.org/cc/v3"
)

var (
	tokLBrace = cc.String("(")
	tokRBrace = cc.String(")")
)

type PreProcessConfig struct {
	IncludePaths    []string `yaml:"include_paths"`
	SysIncludePaths []string `yaml:"sys_include_paths"`
	Predefine       []string `yaml:"predefine"`
}

func Parse(cfg PreProcessConfig, file string) (*cc.AST, error) {

	// Using standard parser config as context
	var ccConfig cc.Config

	cfg64 := Config64()
	ccConfig.ABI = NewABI(cfg64)
	var sourcePaths []cc.Source

	for _, s := range cfg.Predefine {
		sourcePaths = append(sourcePaths, cc.Source{Value: s})
	}

	sourcePaths = append(sourcePaths, cc.Source{Name: file})

	return cc.Translate(&ccConfig, cfg.IncludePaths, cfg.SysIncludePaths, sourcePaths)
}
