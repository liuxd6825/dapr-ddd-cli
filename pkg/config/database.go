package config

type Database struct {
	Mongo bool `yaml:"mongo"`
	Neo4j bool `yaml:"neo4j"`
	MySql bool `json:"mysql"`
}

func (d *Database) HaveDb() bool {
	if d.Neo4j || d.Mongo || d.MySql {
		return true
	}
	return false
}
