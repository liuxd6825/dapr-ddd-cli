package cmd_domain

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
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
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/reg_aggregate_type.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRegisterAggregateType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
