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
	EventName          string     `yaml:"event"`
	AggregateId        string     `yaml:"aggregateId"`
	Properties         Properties `yaml:"properties"`
	Description        string     `yaml:"description"`
	IsAggregateCommand *bool      `yaml:"isAggregateCommand"`
	Aggregate          *Aggregate
	event              *Event
}

func (c *Commands) init(a *Aggregate) {
	if c != nil {
		for name, cmd := range *c {
			cmd.init(a, name)
		}
	}
}

func (c *Commands) GetByEventName(eventName string) *Command {
	if c == nil {
		return nil
	}
	for _, cmd := range *c {
		if strings.ToLower(cmd.EventName) == strings.ToLower(eventName) {
			return cmd
		}
	}
	return nil
}

func (c *Command) init(a *Aggregate, name string) {
	c.Aggregate = a
	c.Name = name
	if c.IsAggregateCommand == nil {
		f := strings.Contains(name, a.Name)
		c.IsAggregateCommand = &f
	}
	c.Properties.Init(a)
	if c.EventName != "" {
		c.event = c.Aggregate.Events[c.EventName]
	}
}

func (c *Command) ServiceFuncName() string {
	if strings.HasSuffix(c.Name, "Command") {
		return c.Name[0 : len(c.Name)-len("Command")]
	}
	return c.Name
}

func (c *Command) IsCreate() bool {
	if c.Action == "" && strings.Contains(c.Name, "CreateCommand") {
		return true
	}
	return strings.ToLower(c.Action) == "create"
}

func (c *Command) IsUpdate() bool {
	if c.Action == "" && strings.Contains(c.Name, "UpdateCommand") {
		return true
	}
	return strings.ToLower(c.Action) == "update"
}

func (c *Command) IsDelete() bool {
	if c.Action == "" && strings.Contains(c.Name, "DeleteCommand") {
		return true
	}
	return strings.ToLower(c.Action) == "delete"
}

func (c *Command) IsCreateOrUpdate() bool {
	if c.IsCreate() || c.IsUpdate() {
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

func (c *Command) IsAggregateCreateOrUpdate() bool {
	return c.IsCreateAggregate() || c.IsCreateOrUpdate()
}

func (c *Command) SnakeName() string {
	return utils.SnakeString(c.Name)
}

func (c *Command) HttpType() string {
	if c.IsCreate() {
		return "POST"
	} else if c.IsUpdate() {
		return "PATCH"
	} else if c.IsDelete() {
		return "DELETE"
	}
	return "POST"
}

func (c *Command) HttpMethod() string {
	methodName := c.Name
	if strings.HasSuffix(methodName, "Command") {
		methodName = methodName[0 : len(methodName)-7]
	}
	return utils.MidlineString(methodName)
}

func (c *Command) ControllerMethod() string {
	methodName := c.Name
	if strings.HasSuffix(methodName, "Command") {
		methodName = methodName[0 : len(methodName)-7]
	}
	return methodName
}

func (c *Command) IsAggregate() bool {
	return c.event.IsAggregate()
}

func (c *Command) IsEntity() bool {
	if c.event == nil {
		return false
	}
	return c.event.IsEntity(c.event.To)
}

func (c *Command) IsAggregateCreateOrUpdateCommand() bool {
	return c.event.IsAggregateUpdateEvent() || c.event.IsAggregateCreateEvent()
}

func (c *Command) IsAggregateDeleteByIdCommand() bool {
	return c.event.IsAggregateDeleteByIdEvent()
}

func (c *Command) IsAggregateCreateCommand() bool {
	return c.event.IsAggregateCreateEvent()
}

func (c *Command) IsAggregateUpdateCommand() bool {
	return c.event.IsAggregateUpdateEvent()
}

func (c *Command) IsAggregateCustomCommand() bool {
	return c.event.IsAggregateCustomEvent()
}

func (c *Command) IsEntityDeleteByIdCommand() bool {
	return c.event.IsEntityDeleteByIdEvent()
}

func (c *Command) IsEntityCreateCommand() bool {
	return c.event.IsEntityCreateEvent()
}

func (c *Command) IsEntityUpdateCommand() bool {
	return c.event.IsEntityUpdateEvent()
}

func (c *Command) IsEntityCustomCommand() bool {
	return c.event.IsEntityCustomEvent()
}

func (c *Command) IsEntityCreateOrUpdateCommand() bool {
	return c.event.IsEntityUpdateEvent() || c.event.IsEntityCreateEvent()
}
