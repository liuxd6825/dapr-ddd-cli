package config

type Configuration struct {
	BoundedContextName        string         `yaml:"boundedContextName"`
	DefaultModule             string         `yaml:"defaultModule"`
	DefaultReservedProperties Metadata       `yaml:"defaultReservedProperties"`
	Metadata                  Metadata       `yaml:"metadata"`
	CSharp                    CSharpMetadata `yaml:"cSharp"`
	Java                      JavaMetadata   `yaml:"java"`
	Go                        GoMetadata     `yaml:"go"`
}

func (c *Configuration) Namespace() string {
	return c.BoundedContextName
}
