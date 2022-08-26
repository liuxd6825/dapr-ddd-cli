package config

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Property struct {
	Name          string ``                     // 属性名称
	Type          string `yaml:"type"`          // 数据类型
	ReferenceType string `yaml:"referenceType"` // 引用类型
	DefaultValue  any    `yaml:"defaultValue"`  // 默认值
	Validate      string `yaml:"validate"`      // 验证说明
	Description   string `yaml:"description"`   // 说明描述
	IsAggregateId bool   `yaml:"isAggregateId"` // 是聚合根ID
	IsArray       bool   `yaml:"isArray"`       // 是否循环 数组类型
	/*	Json           string            `yaml:"json"`          // JSON 属性
		Bson           string            `yaml:"bson"`          // Mongo属性*/
	Uses           []string          `yaml:"uses"`        // 使用范围 view, entity 等扩展
	Tags           map[string]string `yaml:"tags"`        // 扩展标记
	RelationTag    string            `yaml:"relationTag"` // DDD关系
	TypeDefinition *TypeDefinition
	Aggregate      *Aggregate
	Config         *Config
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
		/*		Json:           p.Json,
				Bson:           p.Bson,*/
		Aggregate:      p.Aggregate,
		TypeDefinition: p.TypeDefinition,
	}
	return t
}

func (p *Property) init(a *Aggregate, c *Config, name string) {
	p.Name = name
	p.Aggregate = a
	p.Config = c
	if len(p.Type) == 0 {
		p.Type = "string"
	}
	/*	p.Json = utils.FirstLower(name)
		p.Bson = utils.FirstLower(name)*/

	if td, ok := p.Config.TypeDefinitions.FindByName(p.Type); ok {
		p.TypeDefinition = td
	}
}

//
// LanType
// @Description: 当前语言下的类型
// @receiver p
// @return string
//
func (p *Property) LanType() string {
	return p.lanType(false)
}

func (p *Property) GoLanType() string {
	return p.lanType(true)
}

func (p *Property) IsObjectType() bool {
	t, ok := p.Config.TypeDefinitions[p.Type]
	if ok {
		return t.IsObjectType
	}
	return true
}

func (p *Property) lanType(field bool) string {
	dataType := ""
	if p == nil {
		return ""
	}
	if p.Config != nil {
		dataType = p.Config.GetType(p.Type)
	} else if p.Aggregate != nil && p.Aggregate.Config != nil {
		dataType = p.Aggregate.Config.GetType(p.Type)
	} else {
		dataType = p.Type
	}

	if p.Aggregate != nil && field {
		if v, ok := p.Aggregate.FieldsObjects.Find(dataType); ok {
			return "field." + v.Name
		}
		if v, ok := p.Aggregate.EnumObjects.Find(dataType); ok {
			return "field." + v.Name
		}
		if v, ok := p.Aggregate.ValueObjects.Find(dataType); ok {
			return "field." + v.Name
		}
	}

	/*	if p.IsArray && p.Config.lanType == Go {
		return fmt.Sprintf("[]*%s", dataType)
	}*/
	return dataType
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

func (p *Property) PluralName() string {
	return utils.Plural(p.Name)
}

func (p *Property) BsonName() string {
	v := utils.SnakeString(p.Name)
	if v == "id" {
		v = "_id"
	}
	return v
}

func (p *Property) GormTagName() string {
	fieldName := utils.SnakeString(p.Name)
	tagName := ""
	if p.NameIsId() {
		tagName += `;primaryKey`
	}
	if strings.HasSuffix(fieldName, "_id") {
		tagName += `;index:idx_` + fieldName
	}
	if len(tagName) > 0 {
		tagName = tagName[1:]
	}
	return fmt.Sprintf(`gorm:"%v"`, tagName)
}

func (p *Property) NameIsId() bool {
	if strings.ToLower(p.Name) == "id" {
		return true
	}
	return false
}
func (p *Property) IsItems() bool {
	return strings.ToLower(p.Name) == "items"
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

func (p *Property) GoTags() string {
	tags := fmt.Sprintf(`json:"%s" validate:"%s" `, p.JsonName(), p.GetValidate())
	for k, v := range p.Tags {
		tags = fmt.Sprintf(`%s %s:"%s"`, tags, k, v)
	}
	return tags
}

func (p *Property) GoDescription() string {
	if len(p.Description) > 0 {
		return "// " + p.Description
	}
	return ""
}

func (p *Property) IsTime() bool {
	return strings.ToLower(p.Type) == "time"
}

func (p *Property) IsPTime() bool {
	return strings.ToLower(p.Type) == "ptime"
}

func (p *Property) IsDate() bool {
	return strings.ToLower(p.Type) == "date"
}

func (p *Property) IsPDate() bool {
	return strings.ToLower(p.Type) == "pdate"
}

func (p *Property) IsTimes() bool {
	return p.IsTimeOrPTime()
}

func (p *Property) IsTimeOrPTime() bool {
	return p.IsTime() || p.IsPTime()
}

func (p *Property) IsDates() bool {
	return p.IsDateOrPDate()
}
func (p *Property) IsDateOrPDate() bool {
	return p.IsDate() || p.IsPDate()
}

func (p *Property) IsTimesOrDates() bool {
	return p.IsDateOrPDate() || p.IsTimeOrPTime()
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

func (p *Property) GetValidate() string {
	if len(p.Validate) == 0 {
		if strings.HasSuffix(p.Name, "Id") {
			return "required"
		} else {
			return "-"
		}
	}
	return p.Validate
}
