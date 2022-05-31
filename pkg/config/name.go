package config

import "github.com/liuxd6825/dapr-ddd-cli/pkg/utils"

type BaseProperty struct {
	Name string
}

func (p BaseProperty) UpperName() string {
	return utils.FirstUpper(p.Name)
}

func (p BaseProperty) LowerName() string {
	return utils.FirstLower(p.Name)
}
