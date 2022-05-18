package config

import (
	"strings"
)

type Commands map[string]*Command

type Command struct {
	Name        string
	IsHandler   bool       `yaml:"isHandler"`
	IsCreate    bool       `yaml:"isCreate"`
	Event       string     `yaml:"event"`
	AggregateId string     `yaml:"aggregateId"`
	Properties  Properties `yaml:"properties"`
	Description string     `yaml:"description"`
	Aggregate   *Aggregate
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
	c.Properties.init(a)
}

func (c *Command) ServiceFuncName() string {
	if strings.HasSuffix(c.Name, "Command") {
		return c.Name[0 : len(c.Name)-len("Command")]
	}
	return c.Name
}
