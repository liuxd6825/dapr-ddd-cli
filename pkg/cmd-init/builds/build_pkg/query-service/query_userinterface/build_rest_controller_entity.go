package query_userinterface

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildRestControllerEntity struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRestControllerEntity(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRestControllerEntity {
	res := &BuildRestControllerEntity{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/controller/controller.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestControllerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["Properties"] = b.entity.Properties
	return res
}
