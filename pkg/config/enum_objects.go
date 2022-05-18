package config

type EnumObjects map[string]*EnumObject

type EnumObject struct {
	Name        string
	Description string `yaml:"description"`
	Aggregate   *Aggregate
}

func (e *EnumObjects) init(a *Aggregate) {
	if e == nil {
		return
	}
	for name, enumObjects := range *e {
		enumObjects.init(a, name)
	}
}

func (e *EnumObject) init(a *Aggregate, name string) {
	if e == nil {
		return
	}
	e.Name = name
	e.Aggregate = a
}
