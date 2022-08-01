package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildQueryService struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildQueryService(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildQueryService {
	res := &BuildQueryService{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/service/query_service.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryService) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
