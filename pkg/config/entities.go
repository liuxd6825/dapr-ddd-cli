package config

type Entities map[string]*Entity

type Entity struct {
	Name       string
	Properties Properties `yaml:"properties"`
}

func (e *Entities) init() {
	if e != nil {
		for name, entity := range *e {
			entity.Name = name
		}
	}

}
