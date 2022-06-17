package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildQueryServiceAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildQueryServiceAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildQueryServiceAggregate {
	res := &BuildQueryServiceAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/service/query_service_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryServiceAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.aggregate.Name
	res["name"] = b.aggregate.FirstLowerName()
	return res
}
