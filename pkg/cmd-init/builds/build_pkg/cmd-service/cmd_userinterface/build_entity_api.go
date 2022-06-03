package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestEntityApi struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildRestEntityApi(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildRestEntityApi {
	res := &BuildRestEntityApi{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/facade/api_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestEntityApi) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["Commands"] = b.entity.GetCommands()
	res["AggregatePluralName"] = b.aggregate.PluralName()
	res["EntityPluralName"] = b.entity.PluralName()
	res["Name"] = b.entity.Name
	return res
}

func (b *BuildRestEntityApi) ClassName() string {
	return utils.FirstUpper(b.entity.Name + "CommandApi")
}
