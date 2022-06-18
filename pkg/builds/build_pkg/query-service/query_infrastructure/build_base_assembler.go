package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseAssembler struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseAssembler(base builds.BaseBuild, outFile string) *BuildBaseAssembler {
	res := &BuildBaseAssembler{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/userinterface/rest/assembler/base_assembler.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseAssembler) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
