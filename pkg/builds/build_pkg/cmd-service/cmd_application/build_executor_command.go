package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildCommandExecutor struct {
	builds.BaseBuild
	command   *config.Command
	aggregate *config.Aggregate
}

func NewBuildExecutor(base builds.BaseBuild, aggregate *config.Aggregate, command *config.Command, outFile string) *BuildCommandExecutor {
	res := &BuildCommandExecutor{
		BaseBuild: base,
		command:   command,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/executor/command_executor.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildCommandExecutor) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Command"] = b.command
	res["ClassName"] = utils.FirstUpper(b.command.Name + "Executor")
	res["Name"] = utils.FirstUpper(b.command.Name)
	res["name"] = utils.FirstLower(b.command.Name)
	res["AppName"] = b.command.AppName()
	res["Description"] = b.command.Description
	return res
}
