package config

type EnumObjects map[string]*EnumObject

type EnumObject struct {
	Name        string
	Description string `yaml:"description"`
}

func (e *EnumObjects) init() {
	if e == nil {
		return
	}
	for name, enumObjects := range *e {
		enumObjects.Name = name
	}
}
