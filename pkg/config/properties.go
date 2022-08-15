package config

import "strings"

type Properties map[string]*Property

const TenantId = "TenantId"

func NewProperties(agg *Aggregate, properties, delProperties *Properties) *Properties {
	res := &Properties{}
	if properties != nil {
		res.Adds(properties)
	}

	if delProperties != nil {
		m := *res
		for k := range *delProperties {
			if _, ok := m[k]; ok {
				delete(*res, k)
			}
		}
	}
	res.Init(agg, agg.Config)
	return res
}

func (p *Properties) IsNull() bool {
	return p == nil || *p == nil
}

func (p *Properties) Init(a *Aggregate, c *Config) {
	if p.IsNull() {
		return
	}
	for name, property := range *p {
		property.init(a, c, name)
	}
}

func (p *Properties) Find(name string) (*Property, bool) {
	if p == nil {
		return nil, false
	}
	m := *p
	v, ok := m[name]
	return v, ok
}

func (p *Properties) IsItems() bool {
	_, ok := p.Find("Items")
	return ok
}

func (p *Properties) Adds(sources *Properties) {
	if p.IsNull() || sources.IsNull() {
		return
	}

	for name, property := range *sources {
		m := *p
		if m != nil {
			if _, ok := m[name]; !ok {
				m[name] = property
			}
		}
	}
}

func (p *Properties) Add(property *Property) {
	if p.IsNull() || property == nil {
		return
	}
	m := *p
	m[property.Name] = property
}

func (p *Properties) AddTenantId(a *Aggregate) {
	if p.IsNull() {
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

func (p *Properties) HasType(typeName string) bool {
	if p.IsNull() {
		return false
	}
	m := *p
	for _, item := range m {
		if strings.ToLower(item.Type) == strings.ToLower(typeName) {
			return true
		}
	}
	return false
}

func (p *Properties) HasDateTimeType() bool {
	return p.HasType("dateTime")
}

func (p *Properties) HasDateType() bool {
	return p.HasType("date")
}

func (p *Properties) HasTimeType() bool {
	return p.HasType("time")
}

func (p *Properties) GetDataFieldProperties() *Properties {
	if p.IsNull() {
		return &Properties{}
	}
	for _, item := range *p {
		if item.IsData() {
			fieldName := item.Type
			field := item.Aggregate.FieldsObjects[fieldName]
			if field != nil {
				return &field.Properties
			}
		}
	}
	return &Properties{}
}

func (p *Properties) NewViewProject(config *Config) *Properties {
	properties := &Properties{}
	properties.Adds(p)
	properties.Adds(config.GetDefaultViewProperties())
	return properties
}
