package query_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildRestApiAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRestApiAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRestApiAggregate {
	res := &BuildRestApiAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/facade/api_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestApiAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.aggregate.Name
	res["AggregatePluralName"] = b.aggregate.PluralName()
	res["ServiceName"] = b.aggregate.FirstUpperName() + "AppService"
	return res
}
