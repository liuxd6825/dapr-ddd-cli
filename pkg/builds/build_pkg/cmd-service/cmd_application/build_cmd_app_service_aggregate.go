package cmd_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildCmdAppServiceAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildCmdAppServiceAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildCmdAppServiceAggregate {
	res := &BuildCmdAppServiceAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/service/cmd_app_service_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildCmdAppServiceAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Commands"] = b.aggregate.AggregateCommands
	res["ClassName"] = b.ClassName()
	return res
}

func (b *BuildCmdAppServiceAggregate) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "CommandAppService")
}

func (b *BuildCmdAppServiceAggregate) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
