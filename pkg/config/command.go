package config

type Commands map[string]*Command

type Command struct {
	Name        string
	IsHandler   bool       `yaml:"isHandler"`
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
