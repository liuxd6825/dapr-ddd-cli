package query_domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildQueryHandlerAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildQueryHandler(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildQueryHandlerAggregate {
	res := &BuildQueryHandlerAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/queryhandler/aggregate/aggregate_query_handler.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildQueryHandlerAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = fmt.Sprintf("%sQueryHandler", utils.FirstUpper(b.AggregateName()))
	res["AggregateName"] = utils.FirstUpper(b.Aggregate.Name)
	res["aggregateName"] = utils.FirstLower(b.Aggregate.Name)
	res["Entities"] = b.Aggregate.Entities
	res["Events"] = b.Aggregate.Events
	res["Commands"] = b.aggregate.Commands
	res["Package"] = fmt.Sprintf("%s_queryhandler", utils.SnakeString(b.Aggregate.Name))
	res["ServiceName"] = b.Config.Configuration.ServiceName
	res["Namespace"] = b.Namespace()
	return res
}
