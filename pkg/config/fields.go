package config

type FieldsObjects map[string]*Fields

type Fields struct {
	Name        string
	Description string
	Properties  Properties `yaml:"properties"`
}

func (f *FieldsObjects) init() {
	if f == nil {
		return
	}
	for name, fields := range *f {
		fields.Name = name
		fields.Properties.init()
	}
}
