package config

type Aggregates map[string]*Aggregate

func (a *Aggregates) init() {
	if a == nil {
		return
	}
	for name, agg := range *a {
		agg.Name = name
		agg.init()
		agg.Commands.init()
		agg.Events.init()
		agg.FieldsObjects.init()
		agg.Events.init()
		agg.EnumObjects.init()
		agg.Factory.init()
	}
}

type Aggregate struct {
	Name          string        `yaml:"name"`
	Id            *Property     `yaml:"id"`
	Description   string        `yaml:"description"`
	Properties    Properties    `yaml:"properties"`
	ValueObjects  ValueObjects  `yaml:"valueObject"`
	EnumObjects   EnumObjects   `yaml:"enumObjects"`
	Entities      Entities      `yaml:"entities"`
	FieldsObjects FieldsObjects `yaml:"fields"`
	Events        Events        `yaml:"events"`
	Commands      Commands      `yaml:"commands"`
	Factory       Factory       `yaml:"factory"`
}

func (a *Aggregate) init() {
	a.initId()
}

func (a *Aggregate) initId() *Property {
	if a.Id != nil {
		return a.Id
	}
	if a.Properties != nil {
		for _, p := range a.Properties {
			if p.IsAggregateId {
				a.Id = p
				return p
			}
		}
	}
	a.Id = NewProperty("Id", "string")
	return nil
}
