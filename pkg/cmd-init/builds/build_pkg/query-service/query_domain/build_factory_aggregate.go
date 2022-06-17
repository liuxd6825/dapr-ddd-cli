package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildFactoryAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildFactoryAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildFactoryAggregate {
	res := &BuildFactoryAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/factory/factory_aggregate.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildFactoryAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultAggregateProperties(), &b.aggregate.Properties)
	res["Events"] = b.aggregate.Events.GetAggregateEvents()
	res["Name"] = b.aggregate.FirstUpperName()
	res["name"] = b.aggregate.FirstLowerName()
	res["DefaultProperties"] = defaultProperties
	return res
}
