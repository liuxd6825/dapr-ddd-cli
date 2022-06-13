package config

import "fmt"

type K8S struct {
	Namespace    string `yaml:"namespace"`
	CommandImage string `yaml:"commandImage"`
	QueryImage   string `yaml:"queryImage"`
}

type Configuration struct {
	BoundedContextName        string                     `yaml:"boundedContextName"`        // 领域上下文名称
	DefaultModule             string                     `yaml:"defaultModule"`             // 领域模块
	ServiceName               string                     `yaml:"serviceName"`               // 服务名
	Namespace                 Namespace                  `yaml:"namespace"`                 // 命名空间
	ApiVersion                string                     `yaml:"apiVersion"`                // API 版本号
	Description               string                     `yaml:"description"`               // 领域上下文说明
	DefaultReservedProperties *DefaultReservedProperties `yaml:"defaultReservedProperties"` // 系统中固定的字段
	Metadata                  Metadata                   `yaml:"metadata"`                  // 领域元数据
	CSharp                    Metadata                   `yaml:"c#"`                        // C# 元数据
	Java                      Metadata                   `yaml:"java"`                      // Java 元数据
	Go                        Metadata                   `yaml:"go"`                        // Go 元数据
	K8s                       K8S                        `yaml:"k8s"`                       // K8s 元数据
	GoUtil                    *MetadataUtil
	CSharpUtil                *MetadataUtil
	JavaUtil                  *MetadataUtil
	LangType                  LangType
}

type DefaultReservedProperties struct {
	AggregateProperties Properties `yaml:"aggregate"`
	EntityProperties    Properties `yaml:"entity"`
	ValueProperties     Properties `yaml:"value"`
	ViewProperties      Properties `yaml:"view"`
	FieldProperties     Properties `yaml:"field"`
}

func (c *Configuration) Init(config *Config, langType LangType) {
	c.LangType = langType
	c.GoUtil = NewMetadataUtil(c.Go)
	c.DefaultReservedProperties.init(config)
}

func (c *Configuration) GetK8sNamespace() string {
	return c.K8s.Namespace
}

func (c *Configuration) GetK8sQueryImage() string {
	return c.K8s.QueryImage
}

func (c *Configuration) GetK8sCommandImage() string {
	return c.K8s.CommandImage
}

func (c *Configuration) GetNamespace() string {
	switch c.LangType {
	case Go:
		return c.Namespace.Go
	case Java:
		return c.Namespace.Java
	case CSharp:
		return c.Namespace.CSharp
	}
	return "{{.Namespace}}"
}

func (c *Configuration) QueryServiceName() string {
	return fmt.Sprintf("%s-query-service", c.ServiceName)
}

func (c *Configuration) CommandServiceName() string {
	return fmt.Sprintf("%s-command-service", c.ServiceName)
}

func (p *DefaultReservedProperties) init(c *Config) {
	p.ViewProperties.Init(nil, c)
	p.FieldProperties.Init(nil, c)
	p.AggregateProperties.Init(nil, c)
	p.EntityProperties.Init(nil, c)
	p.ValueProperties.Init(nil, c)
}
