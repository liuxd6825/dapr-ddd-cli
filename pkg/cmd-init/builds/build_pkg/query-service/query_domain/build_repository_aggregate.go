package query_domain

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildRepositoryAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRepositoryAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRepositoryAggregate {
	res := &BuildRepositoryAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/repository/repository_aggregate.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRepositoryAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.AggregateName()
	return res
}
