package config

type ValueObjects map[string]*ValueObject

type ValueObject struct {
	Name        string
	Aggregate   *Aggregate
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

func (v *ValueObjects) init(a *Aggregate) {
	if v == nil {
		return
	}

	valueObjects := *v
	props := a.Config.GetDefaultEntityProperties()
	for _, value := range valueObjects {
		if value != nil && value.Properties != nil {
			value.Properties.Adds(props)
		}
	}
	for name, item := range *v {
		item.init(a, name)
	}
}

func (v *ValueObject) init(a *Aggregate, name string) {
	v.Name = name
	v.Aggregate = a
	v.Properties.Init(a, a.Config)
}
