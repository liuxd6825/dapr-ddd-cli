package query_application

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildAppServiceEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRestControllerEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAppServiceEntity {
	res := &BuildAppServiceEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internales/query_appservice/app_service.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAppServiceEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["Properties"] = b.entity.Properties
	return res
}
