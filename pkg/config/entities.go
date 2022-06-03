package config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type Entities map[string]*Entity

type Entity struct {
	Name        string
	Description string     `yaml:"description"`
	IdInfo      *IdInfo    `yaml:"id"`
	Properties  Properties `yaml:"properties"`
	Aggregate   *Aggregate
}

func (e *Entities) init(a *Aggregate) {
	if e == nil {
		return
	}
	entities := *e
	props := a.Config.GetDefaultEntityProperties()
	for _, entity := range entities {
		if entity.Properties != nil {
			entity.Properties.Adds(props)
		}
	}
	for name, entity := range *e {
		entity.int(a, name)
	}
}

func (e *Entity) int(a *Aggregate, name string) {
	if e == nil {
		return
	}
	e.Name = name
	e.Aggregate = a
	e.Properties.Init(a)
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
	return utils.Plural(e.Name)
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
