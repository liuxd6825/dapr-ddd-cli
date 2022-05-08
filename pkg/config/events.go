package config

type Event struct {
	Name        string
	AggregateId string     `yaml:"aggregateId"`
	Version     string     `yaml:"version"`
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

type Events map[string]*Event

func (e *Events) init() {
	if e != nil {
		for name, event := range *e {
			event.Name = name
			if len(event.Version) == 0 {
				event.Version = "1"
			}
			event.Properties.init()
		}
	}
}
