package cmd_service

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildCommand struct {
	builds.BaseBuild
	command *config.Command
}

func NewBuildCommand(base builds.BaseBuild, command *config.Command) *BuildCommand {
	res := &BuildCommand{
		BaseBuild: base,
		command:   command,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/command/commands/command.go.tpl"
	res.OutFile = ""
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildCommand) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["Properties"] = b.command.Properties
	res["Description"] = b.command.Description
	res["Name"] = b.command.Name
	res["IsHandler"] = b.command.IsHandler
	res["AggregateId"] = b.command.AggregateId
	res["Package"] = b.Package()
	return res
}

func (b *BuildCommand) ClassName() string {
	return utils.FirstUpper(b.command.Name)
}

func (b *BuildCommand) Package() string {
	return fmt.Sprintf("%s_command", utils.FirstLower(b.AggregateName()))
}
