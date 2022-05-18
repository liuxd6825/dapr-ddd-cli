package config

type ValueObjects map[string]*ValueObject

type ValueObject struct {
	Name        string
	Aggregate   *Aggregate
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

func (v *ValueObjects) init(a *Aggregate) {
	for name, item := range *v {
		item.init(a, name)
	}
}

func (v *ValueObject) init(a *Aggregate, name string) {
	v.Name = name
	v.Aggregate = a
	v.Properties.init(a)
}
