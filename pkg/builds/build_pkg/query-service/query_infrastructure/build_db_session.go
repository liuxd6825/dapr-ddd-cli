package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
)

type BuildDbSession struct {
	builds.BaseBuild
}

func NewBuildDbSession(base builds.BaseBuild, outFile string) *BuildDbSession {
	res := &BuildDbSession{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/db/session/session.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDbSession) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	database := b.Config.Configuration.Database
	values["Database"] = database
	values["IsMongo"] = database.Mongo
	values["IsNeo4j"] = database.Neo4j
	values["HaveDb"] = database.HaveDb()
	return values
}
