package query_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
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
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/assembler/assembler_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssemblerAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.aggregate.Name
	res["name"] = b.aggregate.LowerName()
	res["Properties"] = b.aggregate.Properties
	res["Description"] = b.aggregate.Description
	return res
}
