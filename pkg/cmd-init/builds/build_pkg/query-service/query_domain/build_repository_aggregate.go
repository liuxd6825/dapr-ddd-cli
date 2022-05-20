package query_domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
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
	res["ClassName"] = fmt.Sprintf("%sDomainService", utils.FirstUpper(b.AggregateName()))
	res["Commands"] = b.aggregate.Commands
	res["AggregateName"] = b.Aggregate.Name
	res["Package"] = fmt.Sprintf("%s_model", utils.ToLower(b.AggregateName()))
	res["CommandPackage"] = fmt.Sprintf("%s_commands", utils.ToLower(b.aggregate.Name))
	res["ModelPackage"] = fmt.Sprintf("%s_model", utils.ToLower(b.aggregate.Name))
	return res
}
