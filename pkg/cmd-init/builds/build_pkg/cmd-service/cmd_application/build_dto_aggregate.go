package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildDtoAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildDtoAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildDtoAggregate {
	res := &BuildDtoAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/dto/dto.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDtoAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Commands"] = b.aggregate.AggregateCommands
	return res
}
