package config

type Configuration struct {
	BoundedContextName        string                     `yaml:"boundedContextName"`
	DefaultModule             string                     `yaml:"defaultModule"`
	ServiceName               string                     `yaml:"serviceName"`
	Namespace                 Namespace                  `yaml:"namespace"`
	Description               string                     `yaml:"description"`
	DefaultReservedProperties *DefaultReservedProperties `yaml:"defaultReservedProperties"`
	Metadata                  Metadata                   `yaml:"metadata"`
	CSharp                    Metadata                   `yaml:"c#"`
	Java                      Metadata                   `yaml:"java"`
	Go                        Metadata                   `yaml:"go"`

	GoUtil     *MetadataUtil
	CSharpUtil *MetadataUtil
	JavaUtil   *MetadataUtil

	LangType LangType
}

type DefaultReservedProperties struct {
	AggregateProperties Properties `yaml:"aggregate"`
	EntityProperties    Properties `yaml:"entity"`
	ValueProperties     Properties `yaml:"value"`
	ViewProperties      Properties `yaml:"view"`
}

func (c *Configuration) Init(langType LangType) {
	c.LangType = langType
	c.GoUtil = NewMetadataUtil(c.Go)
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
