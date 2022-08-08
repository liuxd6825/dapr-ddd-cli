package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestApi struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildRestApi(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildRestApi {
	res := &BuildRestApi{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	tmplFile := "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/facade/"
	res.TmplFile = tmplFile + "api_aggregate.go.tpl"
	if entity != nil {
		res.TmplFile = tmplFile + "api_entity.go.tpl"
	}
	res.OutFile = outFile
	return res
}

func (b *BuildRestApi) Values() map[string]interface{} {
	values := b.ValuesOfEntity(b.entity)
	if b.entity != nil {
		values["ClassName"] = utils.FirstUpper(b.entity.Name + "CommandApi")
		values["Commands"] = b.entity.GetCommands()
		values["EntityPluralName"] = b.entity.PluralName()
	} else {
		values["ClassName"] = utils.FirstUpper(b.Aggregate.Name + "CommandApi")
		values["Commands"] = b.Aggregate.Commands
		values["AggregatePluralName"] = b.Aggregate.PluralName()
	}
	return values
}
