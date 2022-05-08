package config

type TypeDefinitions map[string]*TypeDefinition

type TypeDefinition struct {
	Name       string `yaml:"name"`
	SqlType    string `yaml:"sqlType"`
	CSharpType string `yaml:"cSharpType"`
	JavaType   string `yaml:"javaType"`
	GoType     string `yaml:"goType"`
}

func (t *TypeDefinitions) init() {
	if t == nil {
		return
	}
}
