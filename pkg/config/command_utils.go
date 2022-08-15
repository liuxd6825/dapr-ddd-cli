package config

import "strings"

var (
	commandNames = []string{"CreateCommand", "UpdateCommand", "DeleteCommand", "ManyCreateCommand", "ManyUpdateCommand", "ManyDeleteCommand"}
)

type CommandUtils struct {
}

func (c *CommandUtils) getEventName(commandName string) string {
	return strings.ReplaceAll(commandName, "Command", "Event")
}

func (c *CommandUtils) getFieldsName(commandName string) string {
	return strings.ReplaceAll(commandName, "Command", "Fields")
}

func (c *CommandUtils) getAction(commandName string) string {
	if strings.HasSuffix(commandName, "UpdateCommand") {
		return "update"
	} else if strings.HasSuffix(commandName, "CreateCommand") {
		return "create"
	} else if strings.HasSuffix(commandName, "DeleteCommand") {
		return "delete"
	}
	return ""
}

func (c *CommandUtils) getTo(command *Command) string {
	if command == nil {
		return ""
	}
	var toNames []string
	cmdName := command.Name
	for _, name := range commandNames {
		if strings.Contains(cmdName, name) {
			toName := strings.ReplaceAll(cmdName, name, "")
			toNames = append(toNames, toName)
		}
	}
	if len(toNames) == 1 {
		return toNames[0]
	}
	return command.Aggregate.Name
}
