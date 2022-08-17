package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAssembler struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildAssembler(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAssembler {
	res := &BuildAssembler{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/assembler/assembler.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssembler) Values() map[string]interface{} {
	values := b.BaseBuild.ValuesOfEntity(b.entity)
	if b.entity != nil {
		defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultEntityProperties(), &b.entity.Properties)
		values["Commands"] = b.entity.GetCommands()
		values["DefaultProperties"] = defaultProperties
	} else {
		values["Commands"] = b.Aggregate.AggregateCommands()
		values["DefaultProperties"] = b.Aggregate.Properties
	}
	return values
}
