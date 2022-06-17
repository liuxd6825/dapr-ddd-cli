package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildFactoryEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildFactoryEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildFactoryEntity {
	res := &BuildFactoryEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/factory/factory_entity.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildFactoryEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultViewProperties(), &b.aggregate.Properties)
	res["Events"] = b.aggregate.Events.GetEntityEvents(b.entity.Name)
	res["Name"] = b.entity.FirstUpperName()
	res["name"] = b.entity.FirstLowerName()
	res["Entity"] = b.entity
	res["DefaultProperties"] = defaultProperties
	return res
}
