package config

type Database struct {
	Mongo bool `yaml:"mongo"`
	Neo4j bool `yaml:"neo4j"`
}

func (d *Database) HaveDb() bool {
	if d.Neo4j || d.Mongo {
		return true
	}
	return false
}
