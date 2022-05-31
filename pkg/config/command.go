package config

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Commands map[string]*Command

type Command struct {
	Name               string
	IsHandler          bool       `yaml:"isHandler"`
	Action             string     `yaml:"action"`
	Event              string     `yaml:"event"`
	AggregateId        string     `yaml:"aggregateId"`
	Properties         Properties `yaml:"properties"`
	Description        string     `yaml:"description"`
	IsAggregateCommand *bool      `yaml:"isAggregateCommand"`
	Aggregate          *Aggregate
}

func (c *Commands) init(a *Aggregate) {
	if c != nil {
		for name, cmd := range *c {
			cmd.init(a, name)
		}
	}
}

func (c *Command) init(a *Aggregate, name string) {
	c.Aggregate = a
	c.Name = name
	if c.IsAggregateCommand == nil {
		f := strings.Contains(name, a.Name)
		c.IsAggregateCommand = &f
	}
	c.Properties.Init(a)
}

func (c *Command) ServiceFuncName() string {
	if strings.HasSuffix(c.Name, "Command") {
		return c.Name[0 : len(c.Name)-len("Command")]
	}
	return c.Name
}

func (c *Command) IsCreate() bool {
	if c.Action == "" && strings.HasSuffix(strings.ToLower(c.Name), "createcommand") {
		return true
	}
	return strings.ToLower(c.Action) == "create"
}

func (c *Command) IsUpdate() bool {
	if c.Action == "" && strings.HasSuffix(strings.ToLower(c.Name), "updatecommand") {
		return true
	}
	return strings.ToLower(c.Action) == "update"
}

func (c *Command) IsDelete() bool {
	if c.Action == "" && strings.HasSuffix(strings.ToLower(c.Name), "deletecommand") {
		return true
	}
	return strings.ToLower(c.Action) == "delete"
}

func (c *Command) IsCreateOrUpdate() bool {
	if c.IsUpdate() || c.IsUpdate() {
		return true
	}
	return false
}

func (c *Command) IsCreateAggregate() bool {
	if *c.IsAggregateCommand && c.IsCreate() {
		return true
	}
	return false
}

func (c *Command) IsUpdateAggregate() bool {
	if *c.IsAggregateCommand && c.IsUpdate() {
		return true
	}
	return false
}

func (c *Command) IsDeleteAggregate() bool {
	if *c.IsAggregateCommand && c.IsDelete() {
		return true
	}
	return false
}

func (c *Command) SnakeName() string {
	return utils.SnakeString(c.Name)
}
