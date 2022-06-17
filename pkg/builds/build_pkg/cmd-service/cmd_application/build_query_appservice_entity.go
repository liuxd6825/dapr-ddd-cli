package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildQueryAppServiceEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildQueryAppServiceEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildQueryAppServiceEntity {
	res := &BuildQueryAppServiceEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/service/query_app_service_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryAppServiceEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["name"] = utils.FirstLower(b.entity.Name)
	res["ResourceName"] = b.entity.SnakeName()
	res["Description"] = b.entity.Description
	return res
}
