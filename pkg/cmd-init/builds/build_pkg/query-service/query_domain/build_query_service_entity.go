package query_domain

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildQueryServiceEntity struct {
	builds.BaseBuild
	Entity *config.Entity
}

func NewBuildQueryServiceEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildQueryServiceEntity {
	res := &BuildQueryServiceEntity{
		BaseBuild: base,
		Entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/queryservice/query_service_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryServiceEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.Entity.Name
	return res
}
