package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildRegisterAggregateType struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRegisterAggregateType(base builds.BaseBuild, outFile string) *BuildRegisterAggregateType {
	res := &BuildRegisterAggregateType{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/register/register_aggregate_type.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRegisterAggregateType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
