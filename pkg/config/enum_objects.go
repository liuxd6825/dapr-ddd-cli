package config

type EnumObjects map[string]*EnumObject

func (e *EnumObjects) init(a *Aggregate) {
	if e == nil {
		return
	}
	for name, enumObjects := range *e {
		enumObjects.init(a, name)
	}
}

func (e *EnumObjects) Adds(sources EnumObjects) {
	if e == nil || sources == nil {
		return
	}
	m := *e
	for name, item := range sources {
		if m != nil {
			if _, ok := m[name]; !ok {
				m[name] = item
			}
		}
	}
}

func (e *EnumObjects) Find(name string) (*EnumObject, bool) {
	if e == nil {
		return nil, false
	}
	m := *e
	v, ok := m[name]
	return v, ok
}

type EnumObject struct {
	Name        string
	EnumValues  EnumValues `yaml:"values"`
	Description string     `yaml:"description"`
	Aggregate   *Aggregate
}

func (e *EnumObject) init(a *Aggregate, name string) {
	if e == nil {
		return
	}
	e.Name = name
	e.Aggregate = a
	e.EnumValues.init(a)
}

type EnumValues map[string]*EnumValue

func (e *EnumValues) init(a *Aggregate) {
	if e == nil {
		return
	}
	for name, enumValue := range *e {
		enumValue.init(a, name)
	}
}

type EnumValue struct {
	Name        string
	Value       uint   `yaml:"value"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Aggregate   *Aggregate
}

func (e *EnumValue) init(a *Aggregate, name string) {
	if e == nil {
		return
	}
	e.Name = name
	e.Aggregate = a
}

func (e *EnumValue) HasTitle() bool {
	return len(e.Title) > 0
}
