package query_userinterface

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
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/assembler/assembler_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssemblerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.entity.Name
	res["name"] = b.entity.FirstLowerName()
	res["Properties"] = b.entity.Properties
	res["Description"] = b.entity.Description
	return res
}
