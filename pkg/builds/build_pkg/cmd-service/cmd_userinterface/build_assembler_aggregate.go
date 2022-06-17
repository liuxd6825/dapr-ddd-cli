package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildAssemblerAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildAssemblerAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildAssemblerAggregate {
	res := &BuildAssemblerAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/assembler/assembler_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssemblerAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Commands"] = b.aggregate.AggregateCommands
	return res
}

func (b *BuildAssemblerAggregate) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "CommandApi")
}
