package config

import (
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Properties map[string]*Property

const TenantId = "TenantId"

func (p *Properties) Init(a *Aggregate) {
	if p != nil {
		for name, property := range *p {
			property.init(a, name)
		}
	}
}

func (p *Properties) Adds(sources *Properties) {
	if p != nil && sources != nil {
		for name, property := range *sources {
			m := *p
			if _, ok := m[name]; !ok {
				m[name] = property
			}
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
			Type:        "string",
			Validate:    "gt=0",
			Aggregate:   a,
			Description: "租户Id",
		}
	}
}

func NewProperties(agg *Aggregate, properties, delProperties *Properties) *Properties {
	res := &Properties{}
	res.Adds(properties)
	m := *res
	for k, _ := range *delProperties {
		if _, ok := m[k]; ok {
			delete(*res, k)
		}
	}
	res.Init(agg)
	return res
}

type Property struct {
	Name          string
	Type          string `yaml:"type"`
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
		Name: name,
		Type: dataType,
	}
}

func (p *Property) Copy() *Property {
	t := &Property{
		Name:          p.Name,
		Type:          p.Type,
		ReferenceType: p.ReferenceType,
		DefaultValue:  p.DefaultValue,
		Validate:      p.Validate,
		Description:   p.Description,
		IsAggregateId: p.IsAggregateId,
		IsArray:       p.IsArray,
		Aggregate:     p.Aggregate,
	}
	return t
}
func (p *Property) init(a *Aggregate, name string) {
	p.Name = name
	p.Aggregate = a
}

func (p *Property) DataType() string {
	if p == nil {
		return ""
	}
	if p.Aggregate == nil || p.Aggregate.Config == nil {
		return p.Type
	}
	return p.Aggregate.Config.GetType(p.Type)
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

func (p *Property) IsData() bool {
	return strings.ToLower(p.Name) == "data"
}
