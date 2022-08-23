package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
)

type BuildLogs struct {
	builds.BaseBuild
}

func NewBuildLogs(base builds.BaseBuild, outFile string) *BuildLogs {
	res := &BuildLogs{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/logs/logs.go.tpl"
	res.OutFile = outFile
	return res
}
