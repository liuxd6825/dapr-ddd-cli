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

func NewAddressUpdateEvent(cmd *user_commands2.AddressUpdateCommand) *event.AddressUpdateEventV1 {
	return &event.AddressUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressDeleteEvent(cmd *user_commands2.AddressDeleteCommand) *event.AddressDeleteEventV1 {
	return &event.AddressDeleteEventV1{
		TenantId:  cmd.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		UserId:    cmd.UserId,
		AddressId: cmd.AddressId,
	}
}

func NewCreateEvent(cmd *user_commands2.UserCreateCommand) *event.UserCreateEventV1 {
	return &event.UserCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewUpdateEvent(cmd *user_commands2.UserUpdateCommand) *event.UserUpdateEventV1 {
	return &event.UserUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewDeleteEvent(cmd *user_commands2.UserDeleteCommand) *event.UserDeleteEventV1 {
	return &event.UserDeleteEventV1{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Id,
		TenantId:  cmd.TenantId,
	}
}

func NewEvent(eventType string) ddd.DomainEvent {
	switch eventType {
	case event.UserCreateEventType.String():
		return &event.UserCreateEventV1{}
	case event.UserUpdateEventType.String():
		return &event.UserUpdateEventV1{}
	case event.UserDeleteEventType.String():
		return &event.UserDeleteEventV1{}
	}
	return nil
}
