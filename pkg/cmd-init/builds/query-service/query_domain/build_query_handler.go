package query_domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildQueryHandler struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildQueryHandler(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildQueryHandler {
	res := &BuildQueryHandler{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/queryhandler/query_handler.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildQueryHandler) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = fmt.Sprintf("%sQueryHandler", utils.FirstUpper(b.AggregateName()))
	res["AggregateName"] = utils.FirstUpper(b.Aggregate.Name)
	res["aggregateName"] = utils.FirstLower(b.Aggregate.Name)
	res["Entities"] = b.Aggregate.Entities
	res["Events"] = b.Aggregate.Events
	res["Commands"] = b.aggregate.Commands
	res["Package"] = "queryhandler"
	res["CommandPackage"] = fmt.Sprintf("%s_commands", utils.ToLower(b.aggregate.Name))
	res["EventPackage"] = fmt.Sprintf("%s_events", utils.ToLower(b.aggregate.Name))
	res["ModelPackage"] = fmt.Sprintf("%s_model", utils.ToLower(b.aggregate.Name))
	res["FactoryPackage"] = fmt.Sprintf("%s_factory", utils.ToLower(b.aggregate.Name))
	res["ServiceName"] = b.Config.Configuration.ServiceName
	res["Namespace"] = b.Namespace()
	return res
}
