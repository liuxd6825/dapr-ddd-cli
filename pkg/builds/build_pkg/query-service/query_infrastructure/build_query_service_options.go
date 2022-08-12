package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildQueryServiceOptionsImpl struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildQueryServiceOptionsImpl(base builds.BaseBuild, outFile string) *BuildQueryServiceOptionsImpl {
	res := &BuildQueryServiceOptionsImpl{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/service_impl/options.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryServiceOptionsImpl) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
