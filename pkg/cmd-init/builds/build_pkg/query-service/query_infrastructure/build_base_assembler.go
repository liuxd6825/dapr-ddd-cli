package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildAssemblerBase struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildAssemblerBase(base builds.BaseBuild, outFile string) *BuildAssemblerBase {
	res := &BuildAssemblerBase{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/userinterface/rest/assembler/base_assembler.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAssemblerBase) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
