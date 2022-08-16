package query_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAppQuery struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildAppQuery(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAppQuery {
	res := &BuildAppQuery{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/appquery/appquery.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAppQuery) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)

}
