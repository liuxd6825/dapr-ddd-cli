package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"strings"
)

type BuildDtoCommand struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	command   *config.Command
}

func NewBuildDtoCommand(base builds.BaseBuild, aggregate *config.Aggregate, command *config.Command, outFile string) *BuildDtoCommand {
	res := &BuildDtoCommand{
		BaseBuild: base,
		aggregate: aggregate,
		command:   command,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/dto/dto_command.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildDtoCommand) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = strings.ReplaceAll(b.command.Name, "Command", "")
	event := b.command.Event()
	if event != nil && event.DataFields != nil && event.DataFields.Properties != nil {
		res["DataFieldsProperties"] = event.DataFields.Properties
	} else {
		res["DataFieldsProperties"] = config.Properties{}
	}
	res["Description"] = b.command.Description
	return res
}
