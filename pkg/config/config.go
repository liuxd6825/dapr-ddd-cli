package config

import (
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
	langType LangType
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
				return nil, err
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
		config.Configuration.Init(config.langType)
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
			c.Aggregates[k] = v
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
		c.langType = Go
	case "java":
		c.langType = Java
	case "c#":
		c.langType = CSharp
	case "csharp":
		c.langType = CSharp
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

func (c *Config) GetType(value string) string {
	if c == nil {
		return value
	}
	tds := c.TypeDefinitions
	if t, ok := tds[value]; ok {
		switch c.langType {
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
