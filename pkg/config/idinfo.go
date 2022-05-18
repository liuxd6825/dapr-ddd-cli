package config

type IdInfo struct {
	Name string
}

func NewIdInfo() *IdInfo {
	return &IdInfo{}
}
