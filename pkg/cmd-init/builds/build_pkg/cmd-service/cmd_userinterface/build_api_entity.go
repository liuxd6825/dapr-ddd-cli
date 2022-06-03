package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestApiEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildRestApiEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildRestApiEntity {
	res := &BuildRestApiEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/facade/api_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestApiEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["Commands"] = b.entity.GetCommands()
	res["AggregatePluralName"] = b.aggregate.PluralName()
	res["EntityPluralName"] = b.entity.PluralName()
	res["Name"] = b.entity.Name
	return res
}

func (b *BuildRestApiEntity) ClassName() string {
	return utils.FirstUpper(b.entity.Name + "CommandApi")
}
