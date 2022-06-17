package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAssemblerEntity struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	entity    *config.Entity
}

func NewBuildAssemblerEntity(base builds.BaseBuild, aggregate *config.Aggregate, entity *config.Entity, outFile string) *BuildAssemblerEntity {
	res := &BuildAssemblerEntity{
		BaseBuild: base,
		aggregate: aggregate,
		entity:    entity,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/assembler/assembler.go.tpl"
	res.OutFile = outFile
	return res
}
func (b *BuildAssemblerEntity) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Commands"] = b.entity.GetCommands()
	return res
}
