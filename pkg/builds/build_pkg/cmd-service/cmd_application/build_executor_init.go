package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildExecutorInit struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildExecutorInit(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildExecutorInit {
	res := &BuildExecutorInit{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/executor/init.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildExecutorInit) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Commands"] = b.aggregate.Commands
	return res
}
