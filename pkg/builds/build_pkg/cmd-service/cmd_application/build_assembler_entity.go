package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAssembler struct {
	builds.BaseBuild
	entity *config.Entity
}

func NewBuildAssembler(base builds.BaseBuild, entity *config.Entity, outFile string) *BuildAssembler {
	res := &BuildAssembler{
		BaseBuild: base,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/assembler/assembler.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssembler) Values() map[string]interface{} {
	values := b.BaseBuild.ValuesOfEntity(b.entity)
	if b.entity != nil {
		values["Commands"] = b.entity.GetCommands()
	} else {
		values["Commands"] = b.Aggregate.AggregateCommands()
	}
	return values
}
