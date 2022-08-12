package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildQueryServiceImpl struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildQueryServiceImpl(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildQueryServiceImpl {
	res := &BuildQueryServiceImpl{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/service_impl/query_service_impl.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryServiceImpl) Values() map[string]interface{} {
	values := b.BaseBuild.ValuesOfEntity(b.entity)
	return values
}
