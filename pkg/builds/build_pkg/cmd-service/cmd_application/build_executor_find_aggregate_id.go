package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildFindAggregateByIdExecutor struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildFindAggregateByIdExecutor(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildFindAggregateByIdExecutor {
	res := &BuildFindAggregateByIdExecutor{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/executor/find_aggregate_executor.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildFindAggregateByIdExecutor) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	Name := "FindAggregateById"
	res["ClassName"] = utils.FirstUpper(Name + "Executor")
	res["Name"] = utils.FirstUpper(Name)
	res["name"] = utils.FirstLower(Name)
	return res
}
