package config

import (
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type Properties map[string]*Property

const TenantId = "TenantId"

func (p *Properties) init(a *Aggregate) {
	if p != nil {
		for name, property := range *p {
			property.init(a, name)
		}
	}
}

func (p *Properties) AddTenantId(a *Aggregate) {
	if p != nil {
		return
	}
	m := *p
	if _, ok := m[TenantId]; !ok {
		m[TenantId] = &Property{
			Name:        TenantId,
			DataType:    "string",
			Validate:    "gt=0",
			Aggregate:   a,
			Description: "租户Id",
		}
	}
}

type Property struct {
	Name          string
	DataType      string `yaml:"type"`
	ReferenceType string `yaml:"referenceType"`
	DefaultValue  any    `yaml:"defaultValue"`
	Validate      string `yaml:"validate"`
	Description   string `yaml:"description"`
	IsAggregateId bool   `yaml:"isAggregateId"`
	IsArray       bool   `yaml:"isArray"`
	Aggregate     *Aggregate
}

func NewProperty(name string, dataType string) *Property {
	return &Property{
		Name:     name,
		DataType: dataType,
	}
}

func (p *Property) init(a *Aggregate, name string) {
	p.Name = name
	p.Aggregate = a
}

func (p *Property) HasValidate() bool {
	return len(p.Validate) > 0
}

func (p *Property) HasDescription() bool {
	return len(p.Description) > 0
}

func (p *Property) UpperName() string {
	return utils.FirstUpper(p.Name)
}

func (p *Property) LowerName() string {
	return utils.FirstLower(p.Name)
}

func (p *Property) JsonName() string {
	return utils.FirstLower(p.Name)
}
func (p *Property) BsonName() string {
	v := utils.FirstLower(p.Name)
	if v == "id" {
		v = "_id"
	}
	return v
}
