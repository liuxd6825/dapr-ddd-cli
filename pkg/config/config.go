package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type LongType int

type Config struct {
	Aggregates      Aggregates      `yaml:"aggregates"`
	ValueObjects    ValueObjects    `yaml:"valueObjects"`
	TypeDefinitions TypeDefinitions `yaml:"typeDefinitions"`
	Configuration   *Configuration  `yaml:"configuration"`
}

const (
	Go LongType = iota
	Java
	CShape
)

//  当前语言类型
var _longType LongType

//
// setLangType
// @Description:
// @param lang
// @return error
//
func setLangType(lang string) error {
	l := strings.ToLower(lang)
	switch l {
	case "go":
		_longType = Go
	case "java":
		_longType = Java
	case "c#":
	case "cshape":
		_longType = CShape
	default:
		return NewLangTypeError(lang)
	}
	return nil
}

func NewConfig(fileName string) (*Config, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return NewConfigWithByte(bytes)
}

func NewConfigEmpty() *Config {
	return &Config{
		Aggregates:      make(map[string]*Aggregate),
		ValueObjects:    make(map[string]*ValueObject),
		TypeDefinitions: make(map[string]*TypeDefinition),
		Configuration:   &Configuration{},
	}
}

func NewConfigWithDir(dirName string, lang string) (*Config, error) {
	if err := setLangType(lang); err != nil {
		return nil, err
	}

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
			config, err := NewConfig(dirName + "/" + fileName)
			if err != nil {
				return nil, err
			}
			configs = append(configs, config)
		}
	}

	config := NewConfigEmpty()
	for _, c := range configs {
		config.merge(c)
	}
	return config, nil

}

func NewConfigWithByte(bytes []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	if config.Aggregates != nil {
		config.Aggregates.init()
	}

	return &config, nil
}

func (c *Config) merge(source *Config) {
	if source.Aggregates != nil && len(source.Aggregates) > 0 {
		for k, v := range source.Aggregates {
			c.Aggregates[k] = v
		}
	}
	if source.TypeDefinitions != nil && len(source.TypeDefinitions) > 0 {
		for k, v := range source.TypeDefinitions {
			c.TypeDefinitions[k] = v
		}
	}
	if source.ValueObjects != nil && len(source.ValueObjects) > 0 {
		for k, v := range source.ValueObjects {
			c.ValueObjects[k] = v
		}
	}
	if source.Configuration != nil {
		c.Configuration = source.Configuration
	}
}
