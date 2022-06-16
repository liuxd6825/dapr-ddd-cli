package config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Aggregates map[string]*Aggregate

func (a *Aggregates) init(config *Config) {
	if a == nil {
		return
	}
	for name, agg := range *a {
		agg.Name = name
		agg.Config = config
		agg.init()
	}
}

type Aggregate struct {
	Name              string        `yaml:"name"`
	Id                *Property     `yaml:"id"`
	Version           string        `yaml:"version"`
	Description       string        `yaml:"description"`
	Properties        Properties    `yaml:"properties"`
	ValueObjects      ValueObjects  `yaml:"valueObjects"`
	EnumObjects       EnumObjects   `yaml:"enumObjects"`
	Entities          Entities      `yaml:"entities"`
	FieldsObjects     FieldsObjects `yaml:"fields"`
	Events            Events        `yaml:"events"`
	Commands          Commands      `yaml:"commands"`
	Factory           Factory       `yaml:"factory"`
	AggregateCommands *[]Command
	Config            *Config
}

func (a *Aggregate) init() {
	a.initId()
	if a.Version == "" {
		a.Version = "v1.0"
	}
	if a.Properties == nil {
		a.Properties = Properties{}
	}
	// 添加聚合默认属性
	if a.Config != nil && a.Config.Configuration != nil && a.Config.Configuration.DefaultReservedProperties != nil {
		aggregateProperties := a.Config.Configuration.DefaultReservedProperties.AggregateProperties
		a.Properties.Adds(&aggregateProperties)
	}
	a.Properties.Init(a, a.Config)

	if a.Entities == nil {
		a.Entities = Entities{}
	}
	a.Entities.init(a)

	if a.ValueObjects == nil {
		a.ValueObjects = ValueObjects{}
	}
	a.ValueObjects.init(a)

	a.Events.init(a)
	a.Commands.init(a)
	a.FieldsObjects.init(a)
	a.EnumObjects.init(a)
	a.Factory.init(a)
	a.AggregateCommands = a.getAggregateCommands()

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

func (a *Aggregate) LowerName() string {
	return strings.ToLower(a.Name)
}

func (a *Aggregate) FirstLowerName() string {
	return utils.FirstLower(a.Name)
}

func (a *Aggregate) FirstUpperName() string {
	return utils.FirstUpper(a.Name)
}

func (a *Aggregate) FileName() string {
	return utils.SnakeString(a.Name)
}

func (a *Aggregate) SnakeName() string {
	return utils.SnakeString(a.Name)
}

func (a *Aggregate) MidlineName() string {
	return utils.MidlineString(a.Name)
}

func (a *Aggregate) PluralName() string {
	return utils.PluralMidline(a.Name)
}

func (a *Aggregate) getAggregateCommands() *[]Command {
	var commands []Command
	for _, event := range a.Events {
		if event.To == "" || event.To == a.Name {
			command := a.Commands.GetByEventName(event.Name)
			commands = append(commands, *command)
		}
	}
	return &commands
}
