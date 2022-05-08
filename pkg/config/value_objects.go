package config

type ValueObject struct {
	Description string     `yaml:"description"`
	Properties  Properties `yaml:"properties"`
}

type ValueObjects map[string]*ValueObject
