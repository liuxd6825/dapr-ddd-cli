package query_application

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
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/application/internals/assembler/assembler.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssembler) Values() map[string]interface{} {
	return b.BaseBuild.ValuesOfEntity(b.entity)
}
