package cmd_domain

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

func NewBuildCommand(base builds.BaseBuild, command *config.Command, outFile string) *BuildCommand {
	res := &BuildCommand{
		BaseBuild: base,
		command:   command,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/command/command/command.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildCommand) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = utils.FirstUpper(b.command.Name)
	res["Properties"] = b.command.Properties
	res["Description"] = b.command.Description
	res["Name"] = b.command.Name
	res["IsHandler"] = b.command.IsHandler
	res["AggregateId"] = b.command.AggregateId
	res["Package"] = fmt.Sprintf("%s_command", b.command.SnakeName())
	res["Action"] = b.command.Action
	res["IsCreate"] = b.command.IsCreate()
	res["IsUpdate"] = b.command.IsUpdate()
	res["IsDelete"] = b.command.IsDelete()
	res["EventName"] = b.command.Event
	return res
}

func (b *BuildCommand) ClassName() string {
	return utils.FirstUpper(b.command.Name)
}

func (b *BuildCommand) Package() string {
	return fmt.Sprintf("%s_commands", utils.ToLower(b.AggregateName()))
}
