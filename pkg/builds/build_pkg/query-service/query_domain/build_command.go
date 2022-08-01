package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildCommand struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildCommand(base builds.BaseBuild, agg *config.Aggregate, entity *config.Entity, outFile string) *BuildCommand {
	res := &BuildCommand{
		BaseBuild: base,
		entity:    entity,
		aggregate: agg,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/command/query.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	res.entity = entity
	return res
}

func (b *BuildCommand) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
