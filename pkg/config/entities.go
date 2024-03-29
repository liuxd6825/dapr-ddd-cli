package config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type Entities map[string]*Entity

type Entity struct {
	Name           string
	Description    string     `yaml:"description"`
	IdInfo         *IdInfo    `yaml:"id"`
	Properties     Properties `yaml:"properties"`
	Aggregate      *Aggregate
	entityCommands []*Command
	entityEvents   []*Event
}

func (e *Entities) init(a *Aggregate) {
	if e == nil {
		return
	}
	entities := *e
	props := a.Config.GetDefaultEntityProperties()
	for _, entity := range entities {
		if entity.Properties == nil {
			entity.Properties = make(Properties)
		}
		entity.Properties.Adds(props)
	}
	for name, entity := range *e {
		entity.int(a, name)
	}
}

func (e *Entities) Adds(entities Entities) {
	if e == nil || entities == nil {
		return
	}
	em := *e
	for name, entity := range entities {
		if entity != nil {
			if _, ok := em[name]; !ok {
				em[name] = entity
			}
		}
	}
}

func (e *Entities) Find(name string) (*Entity, bool) {
	if e == nil {
		return nil, false
	}
	em := *e
	entity, ok := em[name]
	return entity, ok
}

func (e *Entity) int(a *Aggregate, name string) {
	if e == nil {
		return
	}
	e.Name = name
	e.Aggregate = a
	e.Properties.Init(a, a.Config)
	e.Properties.Adds(a.Config.GetDefaultEntityProperties())

	aggregateId := a.Name + "Id"
	_, ok := e.Properties[aggregateId]
	if !ok {
		aidProp := NewProperty(aggregateId, "string")
		e.Properties[aggregateId] = aidProp
		aidProp.Validate = "gt=0"
	}
}

func (e *Entities) Count() int {
	return len(*e)
}

func (e *Entities) Empty() bool {
	return e.Count() == 0
}

//
// FirstLowerName
// @Description: 首字母小字名称
// @receiver e
// @return string
//
func (e *Entity) FirstLowerName() string {
	return utils.FirstLower(e.Name)
}

//
// FirstUpperName
// @Description: 首字母大字名称
// @receiver e
// @return string
//
func (e *Entity) FirstUpperName() string {
	return utils.FirstUpper(e.Name)
}

//
// FileName
// @Description: 源代码文件名
// @receiver e
// @return string
//
func (e *Entity) FileName() string {
	return utils.SnakeString(e.Name)
}

//
// SnakeName
// @Description: 蛇形名称
// @receiver e
// @return string
//
func (e *Entity) SnakeName() string {
	return utils.SnakeString(e.Name)
}

//
// PluralName
// @Description: 复数名称
// @receiver e
// @return string
//
func (e *Entity) PluralName() string {
	return utils.PluralMidline(e.Name)
}

func (e *Entity) GetCommands() *[]Command {
	var commands []Command
	for _, event := range e.Aggregate.Events {
		if event.To == e.Name {
			command := e.Aggregate.Commands.GetByEventName(event.Name)
			commands = append(commands, *command)
		}
	}
	return &commands
}

func (e *Entity) EntityCommands() []*Command {
	return e.entityCommands
}

func (e *Entity) EntityEvents() []*Event {
	return e.entityEvents
}

func (e *Entity) initEventCommands() {
	var commands []*Command
	commands = []*Command{}
	agg := e.Aggregate
	for _, event := range agg.Events {
		if event.To == "" || event.To == e.Name {
			command := agg.Commands.GetByEventName(event.Name)
			if command != nil {
				commands = append(commands, command)
			}
		}
	}
	e.entityCommands = commands
}

func (e *Entity) initEntityEvents() {
	var events []*Event
	events = []*Event{}
	agg := e.Aggregate
	for _, event := range agg.Events {
		if event.To == "" || event.To == event.Name {
			events = append(events, event)
		}
	}
	e.entityEvents = events
}
