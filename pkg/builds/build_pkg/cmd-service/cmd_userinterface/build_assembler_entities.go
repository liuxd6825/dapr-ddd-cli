package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAssemblerEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildAssemblerEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildAssemblerEntity {
	res := &BuildAssemblerEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/assembler/assembler_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssemblerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultEntityProperties(), &b.entity.Properties)
	res["DefaultProperties"] = defaultProperties
	res["Commands"] = b.entity.GetCommands()
	return res
}
