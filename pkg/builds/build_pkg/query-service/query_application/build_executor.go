package query_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildExecutor struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildExecutor(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildExecutor {
	res := &BuildExecutor{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/executor/executor.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildExecutor) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
