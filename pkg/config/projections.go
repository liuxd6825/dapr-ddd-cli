package config

type Projections map[string]Projection

type Projection struct {
	Name       string
	Properties Properties `yaml:"properties"`
}

func (p *Projections) init(a *Aggregate) {
	if p == nil {
		return
	}
	for name, item := range *p {
		item.init(a, name)
	}
}

func (p *Projection) init(a *Aggregate, name string) {
	p.Name = name
	p.Properties.init(a)
}
