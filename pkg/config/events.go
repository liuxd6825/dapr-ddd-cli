package config

type Event struct {
	Name         string
	AggregateId  string     `yaml:"aggregateId"`
	Version      string     `yaml:"version"`
	Description  string     `yaml:"description"`
	Properties   Properties `yaml:"properties"`
	DataProperty *Property
	DataFields   *Fields
}

type Events map[string]*Event

func (e *Events) init(a *Aggregate) {
	if e != nil {
		for name, event := range *e {
			event.init(a, name)
		}
	}
}

func (e *Event) init(a *Aggregate, name string) {
	e.Name = name
	if len(e.Version) == 0 {
		e.Version = "1"
	}
	e.Properties.init(a)
	if data := e.Properties["data"]; data != nil {
		e.DataProperty = data
		if a.FieldsObjects != nil && e.DataProperty.DataType != "" {
			fields := a.FieldsObjects[e.DataProperty.DataType]
			if fields != nil {
				e.DataFields = fields
			}
		}
	}
}

func (e *Event) ClassName() string {
	return e.Name
}
