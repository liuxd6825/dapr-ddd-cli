package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
)

type BuildUtilsAssembler struct {
	builds.BaseBuild
}

func NewBuildUtilsAssembler(base builds.BaseBuild, outFile string) *BuildUtilsAssembler {
	res := &BuildUtilsAssembler{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/utils/assembler_util.go.tpl"
	res.OutFile = outFile

	return res
}

func (b *BuildUtilsAssembler) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
