package config

type Entities map[string]*Entity

type Entity struct {
	Name        string
	Description string     `yaml:"description"`
	IdInfo      *IdInfo    `yaml:"id"`
	Properties  Properties `yaml:"properties"`
	Aggregate   *Aggregate
}

func (e *Entities) init(a *Aggregate) {
	if e != nil {
		for name, entity := range *e {
			entity.int(a, name)
		}
	}
}

func (e *Entity) int(a *Aggregate, name string) {
	if e == nil {
		return
	}
	e.Name = name
	e.Aggregate = a
	e.Properties.init(a)
}
