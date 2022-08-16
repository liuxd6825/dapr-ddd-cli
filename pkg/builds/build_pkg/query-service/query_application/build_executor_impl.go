package query_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildExecutorImpl struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildExecutorImpl(base builds.BaseBuild, entity *config.Entity, outFile string, tplName string) *BuildExecutorImpl {
	res := &BuildExecutorImpl{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/executor/executor_impl/" + tplName
	res.OutFile = outFile
	return res
}

func (b *BuildExecutorImpl) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)

}
