package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type LangType int

type Config struct {
	Aggregates      Aggregates      `yaml:"aggregates"`
	ValueObjects    ValueObjects    `yaml:"valueObjects"`
	TypeDefinitions TypeDefinitions `yaml:"typeDefinitions"`
	Configuration   *Configuration  `yaml:"configuration"`
	//  当前语言类型
	lanType LangType
}

const (
	Go LangType = iota
	Java
	CSharp
	Sql
)

func NewConfig(lang string) (*Config, error) {
	res := &Config{
		Aggregates:      make(map[string]*Aggregate),
		ValueObjects:    make(map[string]*ValueObject),
		TypeDefinitions: make(map[string]*TypeDefinition),
		Configuration:   &Configuration{},
	}
	if err := res.setLangType(lang); err != nil {
		return nil, err
	}
	return res, nil
}

func NewConfigWidthFile(fileName string) (*Config, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return NewConfigWithByte(bytes)
}

func NewConfigWithDir(dirName string, lang string) (*Config, error) {
	configs := make([]*Config, 0)
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	if len(fileInfos) == 0 {
		return nil, NewReadDirError(dirName)
	}

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if !fileInfo.IsDir() && strings.HasSuffix(fileName, ".yaml") {
			config, err := NewConfigWidthFile(dirName + "/" + fileName)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("field:%s/%s, error: %s", dirName, fileName, err.Error()))
			}
			configs = append(configs, config)
		}
	}

	config, err := NewConfig(lang)
	if err != nil {
		return nil, err
	}
	for _, c := range configs {
		config.merge(c)
	}

	if config.Aggregates != nil {
		config.Aggregates.init(config)
	}
	if config.Configuration != nil {
		config.Configuration.Init(config, config.lanType)
	}

	return config, nil

}

func NewConfigWithByte(bytes []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) merge(source *Config) {
	if source.Aggregates != nil && len(source.Aggregates) > 0 {
		for k, v := range source.Aggregates {
			v.Name = k
			v.initValues()
			agg, ok := c.Aggregates[k]
			if !ok {
				c.Aggregates[k] = v
			} else {
				if agg.Id == nil {
					agg.Id = v.Id
				}
				if len(agg.Name) == 0 {
					agg.Name = v.Name
				}
				if len(agg.Description) == 0 {
					agg.Description = v.Description
				}
				if len(v.Properties) > 0 {
					agg.Properties.Adds(&v.Properties)
				}
				if len(v.Entities) > 0 {
					agg.Entities.Adds(v.Entities)
				}
				if len(v.Commands) > 0 {
					agg.Commands.Adds(v.Commands)
				}
				if len(v.FieldsObjects) > 0 {
					agg.FieldsObjects.Adds(v.FieldsObjects)
				}
				if len(v.EnumObjects) > 0 {
					agg.EnumObjects.Adds(v.EnumObjects)
				}
			}
		}
	}
	if source.TypeDefinitions != nil && len(source.TypeDefinitions) > 0 {
		for k, v := range source.TypeDefinitions {
			v.Name = k
			c.TypeDefinitions[k] = v
		}
	}
	if source.ValueObjects != nil && len(source.ValueObjects) > 0 {
		for k, v := range source.ValueObjects {
			v.Name = k
			c.ValueObjects[k] = v
		}
	}
	if source.Configuration != nil {
		c.Configuration = source.Configuration
	}
}

//
// setLangType
// @Description:
// @param lang
// @return error
//
func (c *Config) setLangType(lang string) error {
	l := strings.ToLower(lang)
	switch l {
	case "go":
		c.lanType = Go
	case "java":
		c.lanType = Java
	case "c#":
		c.lanType = CSharp
	case "csharp":
		c.lanType = CSharp
	default:
		return NewLangTypeError(lang)
	}
	return nil
}

func (c *Config) GetDefaultEntityProperties() *Properties {
	if c != nil && c.Configuration != nil && c.Configuration.DefaultReservedProperties != nil {
		return &c.Configuration.DefaultReservedProperties.EntityProperties
	}
	return nil
}

func (c *Config) GetDefaultAggregateProperties() *Properties {
	if c != nil && c.Configuration != nil && c.Configuration.DefaultReservedProperties != nil {
		return &c.Configuration.DefaultReservedProperties.AggregateProperties
	}
	return nil
}

func (c *Config) GetDefaultValueProperties() *Properties {
	if c != nil && c.Configuration != nil && c.Configuration.DefaultReservedProperties != nil {
		return &c.Configuration.DefaultReservedProperties.ValueProperties
	}
	return nil
}

func (c *Config) GetDefaultViewProperties() *Properties {
	if c != nil && c.Configuration != nil && c.Configuration.DefaultReservedProperties != nil {
		return &c.Configuration.DefaultReservedProperties.ViewProperties
	}
	return nil
}

func (c *Config) GetDefaultFieldProperties() *Properties {
	if c != nil && c.Configuration != nil && c.Configuration.DefaultReservedProperties != nil {
		return &c.Configuration.DefaultReservedProperties.FieldProperties
	}
	return nil
}

func (c *Config) GetType(value string) string {
	if c == nil {
		return value
	}
	tds := c.TypeDefinitions
	if t, ok := tds[value]; ok {
		switch c.lanType {
		case Go:
			return t.GoType
		case Java:
			return t.JavaType
		case CSharp:
			return t.CSharpType
		case Sql:
			return t.SqlType
		}
	}
	return value
}

func (c *Config) GetLanType() LangType {
	return c.lanType
}
