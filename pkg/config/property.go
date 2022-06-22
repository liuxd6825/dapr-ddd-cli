package config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Properties map[string]*Property

const TenantId = "TenantId"

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

func (p *Properties) HasDataTimeType() bool {
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

func NewProperties(agg *Aggregate, properties, delProperties *Properties) *Properties {
	res := &Properties{}
	res.Adds(properties)
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

type Property struct {
	Name          string            ``                     // 属性名称
	Type          string            `yaml:"type"`          // 数据类型
	ReferenceType string            `yaml:"referenceType"` // 引用类型
	DefaultValue  any               `yaml:"defaultValue"`  // 默认值
	Validate      string            `yaml:"validate"`      // 验证说明
	Description   string            `yaml:"description"`   // 说明描述
	IsAggregateId bool              `yaml:"isAggregateId"` // 是聚合根ID
	IsArray       bool              `yaml:"isArray"`       // 是否循环 数组类型
	Json          string            `yaml:"json"`          // JSON 属性
	Bson          string            `yaml:"bson"`          // Mongo属性
	Uses          []string          `yaml:"uses"`          // 使用范围 view, entity 等扩展
	Tags          map[string]string `yaml:"tags"`          // 扩展标记
	Aggregate     *Aggregate
	Config        *Config
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
		Json:          p.Json,
		Bson:          p.Bson,
		Aggregate:     p.Aggregate,
	}
	return t
}

func (p *Property) init(a *Aggregate, c *Config, name string) {
	p.Name = name
	p.Aggregate = a
	p.Config = c
}

//
// LanType
// @Description: 当前语言下的类型
// @receiver p
// @return string
//
func (p *Property) LanType() string {
	if p == nil {
		return ""
	}
	if p.Config != nil {
		return p.Config.GetType(p.Type)
	}
	if p.Aggregate != nil && p.Aggregate.Config != nil {
		return p.Aggregate.Config.GetType(p.Type)
	}
	return p.Type
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
	if p.Json == "" {
		return utils.FirstLower(p.Name)
	}
	return p.Json
}

func (p *Property) PluralName() string {
	return utils.Plural(p.Name)
}

func (p *Property) BsonName() string {
	if p.Bson != "" {
		return p.Bson
	}
	v := utils.SnakeString(p.Name)
	if v == "id" {
		v = "_id"
	}
	return v
}

func (p *Property) IsData() bool {
	return strings.ToLower(p.Name) == "data"
}

func (p *Property) IsUseView() bool {
	return p.IsUse("view")
}

func (p *Property) IsUse(useType string) bool {
	if p == nil {
		return false
	}
	if len(p.Uses) == 0 {
		return true
	}
	for _, ut := range p.Uses {
		if strings.ToLower(ut) == strings.ToLower(useType) {
			return true
		}
	}
	return true
}

func (p *Property) IsTimeType() bool {
	return p.IsDateTimeType()
}

func (p *Property) IsDateTimeType() bool {
	return strings.ToLower(p.Type) == "datetime"
}

func (p *Property) IsDateType() bool {
	return strings.ToLower(p.Type) == "date"
}

func (p *Property) IsEntityType() bool {
	if p == nil || p.Aggregate == nil || p.Aggregate.Entities == nil {
		return false
	}
	_, ok := p.Aggregate.Entities[p.Type]
	return ok
}

func (p *Property) IsEnumType() bool {
	if p == nil || p.Aggregate == nil || p.Aggregate.Entities == nil {
		return false
	}
	_, ok := p.Aggregate.EnumObjects[p.Type]
	return ok
}

func (p *Property) IsArrayEntityType() bool {
	if p.IsEntityType() && p.IsArray {
		return true
	}
	return false
}

func (p *Property) Entity() *Entity {
	if p == nil {
		return nil
	}
	entity, _ := p.Aggregate.Entities[p.Type]
	return entity
}
