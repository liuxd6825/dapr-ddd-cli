package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDtoEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildDtoCommand(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildDtoEntity {
	res := &BuildDtoEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/dto/dto_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDtoEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["Properties"] = b.entity.Properties
	res["Description"] = b.Aggregate.Description
	res["Commands"] = b.entity.GetCommands()
	return res
}
