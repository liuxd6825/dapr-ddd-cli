package config

type Namespace struct {
	Go     string `yaml:"go"`
	CSharp string `yaml:"c#"`
	Java   string `yaml:"java"`
	K8s    string `yaml:"k8s"`
}
