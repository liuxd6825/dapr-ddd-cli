package query_infrastructure

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildQueryServiceImplAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildQueryServiceImplAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildQueryServiceImplAggregate {
	res := &BuildQueryServiceImplAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/domain/service/query_service_impl_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryServiceImplAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.aggregate.Name
	return res
}
