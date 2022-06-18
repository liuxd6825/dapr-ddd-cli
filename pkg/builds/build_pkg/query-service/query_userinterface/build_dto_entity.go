package query_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDtoEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildDtoEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildDtoEntity {
	res := &BuildDtoEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/dto/dto_entity.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDtoEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	props := config.NewProperties(b.aggregate, &b.entity.Properties, b.Config.GetDefaultViewProperties())
	res["Name"] = b.entity.Name
	res["Properties"] = props
	res["Description"] = b.entity.Description
	return res
}
