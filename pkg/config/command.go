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
}

func (c *Commands) init() {
	if c != nil {
		for name, cmd := range *c {
			cmd.Name = name
			cmd.Properties.init()
		}
	}
}

func (c *Command) ServiceFuncName() string {
	if strings.HasSuffix(c.Name, "Command") {
		return c.Name[0 : len(c.Name)-len("Command")]
	}
	return c.Name
}
