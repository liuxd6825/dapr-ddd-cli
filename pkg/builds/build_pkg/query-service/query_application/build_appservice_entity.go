package query_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAppServiceEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildAppService(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAppServiceEntity {
	res := &BuildAppServiceEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/service/app_service.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAppServiceEntity) Values() map[string]interface{} {
	res := b.BaseBuild.ValuesOfEntity(b.entity)
	if b.entity != nil {
		res["Name"] = b.entity.FirstUpperName()
		res["name"] = b.entity.FirstLowerName()
		res["Properties"] = b.entity.Properties
	} else {
		res["Name"] = b.Aggregate.FirstUpperName()
		res["name"] = b.Aggregate.FirstLowerName()
		res["Properties"] = b.Aggregate.Properties
	}
	return res
}
