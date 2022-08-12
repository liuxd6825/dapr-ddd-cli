package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildFactory struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildFactory(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildFactory {
	res := &BuildFactory{
		BaseBuild: base,
		entity:    entity,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/factory/factory.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildFactory) Values() map[string]interface{} {
	values := b.BaseBuild.ValuesOfEntity(b.entity)
	if b.entity != nil {
		defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultViewProperties(), &b.Aggregate.Properties)
		values["DefaultProperties"] = defaultProperties
		values["Events"] = b.Aggregate.Events.GetEntityEvents(b.entity.Name)
		values["Name"] = b.entity.FirstUpperName()
		values["name"] = b.entity.FirstLowerName()
		values["DefaultProperties"] = defaultProperties
	} else {
		defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultViewProperties(), &b.Aggregate.Properties)
		values["DefaultProperties"] = defaultProperties
		values["Events"] = b.Aggregate.Events.GetAggregateEvents()
		values["Name"] = b.Aggregate.FirstUpperName()
		values["name"] = b.Aggregate.FirstLowerName()
	}
	return values
}
