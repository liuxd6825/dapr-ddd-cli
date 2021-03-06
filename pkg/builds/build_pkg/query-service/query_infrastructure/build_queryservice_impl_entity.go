package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildQueryServiceImplEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildQueryServiceImplEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildQueryServiceImplEntity {
	res := &BuildQueryServiceImplEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/service_impl/query_service_impl_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryServiceImplEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["name"] = b.entity.FirstLowerName()
	return res
}
