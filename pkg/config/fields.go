package config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type FieldsObjects map[string]*Fields

type Fields struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

func (f *FieldsObjects) init(a *Aggregate) {
	if f == nil || *f == nil {
		return
	}
	mf := *f
	for _, cmd := range a.Commands {
		if cmd.Fields != nil {
			mf[cmd.Fields.Name] = cmd.Fields
		}
	}

	for name, fields := range mf {
		fields.init(a, name)
	}
}

func (f *FieldsObjects) Adds(fields FieldsObjects) {
	if f == nil || fields == nil {
		return
	}
	em := *f
	for name, field := range fields {
		if field != nil {
			if _, ok := em[name]; !ok {
				em[name] = field
			}
		}
	}
}

func (f *FieldsObjects) Find(name string) (*Fields, bool) {
	if f == nil {
		return nil, false
	}
	m := *f
	v, ok := m[name]
	return v, ok
}

func (e *Fields) init(a *Aggregate, name string) {
	if e.Properties == nil {
		e.Properties = make(Properties)
	}

	e.Name = name
	e.Properties.Init(a, a.Config)
	isItems := e.Properties.IsItems()
	if !isItems {
		e.Properties.Adds(a.Config.GetDefaultFieldProperties())
	}

	if !strings.Contains(name, a.Name) && !isItems {
		aggregateId := a.Name + "Id"
		_, ok := e.Properties[aggregateId]
		if !ok {
			aggIdProp := NewProperty(aggregateId, "string")
			aggIdProp.Validate = "gt=0"
			e.Properties[aggregateId] = aggIdProp
		}
	}

}

func (e *Fields) FileName() string {
	return utils.SnakeString(e.Name)
}
