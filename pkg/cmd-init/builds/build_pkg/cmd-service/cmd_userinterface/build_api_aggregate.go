package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
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
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/facade/api_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestApiAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["Events"] = b.aggregate.Events
	res["Commands"] = b.aggregate.Commands
	res["AggregatePluralName"] = b.aggregate.PluralName()
	return res
}

func (b *BuildRestApiAggregate) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "CommandApi")
}
