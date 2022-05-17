package config

type Configuration struct {
	BoundedContextName        string   `yaml:"boundedContextName"`
	DefaultModule             string   `yaml:"defaultModule"`
	DefaultReservedProperties Metadata `yaml:"defaultReservedProperties"`
	Metadata                  Metadata `yaml:"metadata"`
	CSharp                    Metadata `yaml:"cSharp"`
	Java                      Metadata `yaml:"java"`
	Go                        Metadata `yaml:"go"`

	GoUtil     *MetadataUtil
	CSharpUtil *MetadataUtil
	JavaUtil   *MetadataUtil

	LangType LangType
}

func (c *Configuration) Init(langType LangType) {
	c.LangType = langType
	c.GoUtil = NewMetadataUtil(c.Go)
}

func (c *Configuration) Namespace() string {
	switch c.LangType {
	case Go:
		return c.GoUtil.Namespace()
	case Java:
		return c.JavaUtil.Namespace()
	case CShape:
		return c.CSharpUtil.Namespace()
	}
	return "{{.namespace.}}"
}
