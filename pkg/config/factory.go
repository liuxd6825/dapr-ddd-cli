package config

type Factory map[string]*FactoryFunc

type FactoryMappings map[string]FactoryMapping

type FactoryFuncParameters map[string]FactoryFuncParameter

type FactoryFunc struct {
	Name       string                 `yaml:"name"`
	Parameters *FactoryFuncParameters `yaml:"parameters"`
	Result     *FactoryFuncResult     `yaml:"result"`
	Mappings   *FactoryMappings       `yaml:"mappings"`
}

type FactoryFuncParameter struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type FactoryFuncResult struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type FactoryMapping struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func (f *Factory) init() {
	if f == nil {
		return
	}
	for name, fun := range *f {
		fun.Name = name
		fun.Parameters.init()
		fun.Mappings.init()
		fun.Result.init()
	}
}

func (p *FactoryFuncParameters) init() {
	if p == nil {
		return
	}
	for name, item := range *p {
		item.Name = name
	}
}

func (m *FactoryMappings) init() {
	if m == nil {
		return
	}
	for name, item := range *m {
		item.Name = name
	}
}

func (f *FactoryFuncResult) init() {
	if f == nil {
		return
	}
	if f.Name == "" {
		f.Name = "res"
	}
}
