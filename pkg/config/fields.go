package config

import "github.com/liuxd6825/dapr-ddd-cli/pkg/utils"

type FieldsObjects map[string]*Fields

type Fields struct {
	Name        string
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

func (f *FieldsObjects) init(a *Aggregate) {
	if f == nil {
		return
	}
	for name, fields := range *f {
		fields.Name = name
		fields.Properties.Init(a, a.Config)
	}
}

func (e *Fields) FileName() string {
	return utils.SnakeString(e.Name)
}
