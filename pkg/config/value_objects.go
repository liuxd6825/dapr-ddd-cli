package config

type ValueObjects map[string]*ValueObject

func (vs *ValueObjects) init(a *Aggregate) {
	if vs == nil {
		return
	}

	valueObjects := *vs
	props := a.Config.GetDefaultEntityProperties()
	for _, value := range valueObjects {
		if value.Properties == nil {
			value.Properties = make(Properties)
		}
		value.Properties.Adds(props)
	}
	for name, item := range *vs {
		item.init(a, name)
	}
}

func (vs *ValueObjects) Find(name string) (*ValueObject, bool) {
	if vs == nil {
		return nil, false
	}
	m := *vs
	v, ok := m[name]
	return v, ok
}

type ValueObject struct {
	Name        string
	Aggregate   *Aggregate
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

func (v *ValueObject) init(a *Aggregate, name string) {
	v.Name = name
	v.Aggregate = a
	v.Properties.Init(a, a.Config)
	v.Properties.Adds(a.Config.GetDefaultValueProperties())
}
