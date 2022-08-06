package parser

import (
	"fmt"
	"modernc.org/cc/v3"
)

type PreProcessConfig struct {
	IncludePaths    []string `yaml:"include_paths"`
	SysIncludePaths []string `yaml:"sys_include_paths"`
	Sources         []struct {
		Name  string `yaml:"name,omitempty"`
		Value string `yaml:"value,omitempty"`
	}
}

func Parse(cfg PreProcessConfig) (*cc.AST, error) {

	// Using standard parser config as context
	var ccConfig cc.Config
	
	var sourcePaths []cc.Source

	for _, c := range cfg.Sources {
		if c.Name != "" {
			fmt.Println(c.Name)
			sourcePaths = append(sourcePaths, cc.Source{Name: c.Name})
			continue
		}
		sourcePaths = append(sourcePaths, cc.Source{Value: c.Value})
	}

	return cc.Parse(&ccConfig, cfg.IncludePaths, cfg.SysIncludePaths, sourcePaths)
}
