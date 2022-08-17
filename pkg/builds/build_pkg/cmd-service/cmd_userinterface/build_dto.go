package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDto struct {
	builds.BaseBuild
	entity    *config.Entity
	aggregate *config.Aggregate
}

func NewBuildDto(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildDto {
	res := &BuildDto{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/dto/dto.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDto) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	if b.entity != nil {
		b.AddTimePackageValue(values, &b.entity.Properties)
		values["Properties"] = b.entity.Properties
		values["Name"] = b.entity.Name
		values["Description"] = b.entity.Description
		values["Commands"] = b.entity.GetCommands()
	} else {
		b.AddTimePackageValue(values, &b.Aggregate.Properties)
		values["Name"] = b.Aggregate.Name
		values["Properties"] = b.Aggregate.Properties
		values["Description"] = b.Aggregate.Description
		values["Commands"] = b.Aggregate.AggregateCommands()
		b.AddTimePackageValue(values, &b.Aggregate.Properties)
	}
	return values
}
