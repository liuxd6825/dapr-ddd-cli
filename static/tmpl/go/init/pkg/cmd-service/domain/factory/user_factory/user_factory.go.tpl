package {{.aggregate_name}}_factory

import (
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

{{- range $eventName, $event := .Events }}
func New{{$eventName}}(cmd *command.AddressCreateCommand) *event.AddressCreateEventV1 {
	return &event.{{$eventName}}{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}
{{- end }}
