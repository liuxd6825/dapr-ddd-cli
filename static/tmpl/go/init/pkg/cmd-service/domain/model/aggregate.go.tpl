package {{.Package}}_model

import (
	"context"
	"{{Namespace}}/pkg/cmd-service/domain/command/{{.CommandPackageName}}"
	"{{Namespace}}/pkg/cmd-service/domain/event/{{.EventPackageName}}"
	"{{Namespace}}/pkg/cmd-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type {{.ClassName}} struct {
    {{.Id.Name}}   {{.Id.DataType}}   `json:"{{.Id.Name}}"{{if .Id.HasValidate}}  validate:"{{.Id.Validate}}"{{- end}}`
{{- range $name, $property := .Properties}}
    # {{$property.Description}}
    {{$name}}   {{$property.DataType}}   `json:"{{$name}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`
{{- end}}
}

func New{{.ClassName}}() *{{.ClassName}} {
	return &{{.ClassName}}{
	}
}

func (a *{{.ClassName}}) UserCreateCommand(ctx context.Context, cmd *{{.CommandPackageName}}.UserCreateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewCreateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) UserUpdateCommand(ctx context.Context, cmd *{{.CommandPackageName}}.UserUpdateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewUpdateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) UserDeleteCommand(ctx context.Context, cmd *{{.CommandPackageName}}.UserDeleteCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewDeleteEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) OnUserCreateEventV1s0(ctx context.Context, event *user_events2.UserCreateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Telephone = event.Data.Telephone
	a.Email = event.Data.Email
	return nil
}

func (a *{{.ClassName}}) OnUserUpdateEventV1s0(ctx context.Context, event *user_events2.UserUpdateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Telephone = event.Data.Telephone
	a.Email = event.Data.Email
	return nil
}

func (a *{{.ClassName}}) OnUserDeleteEventV1s0(ctx context.Context, event *user_events2.UserDeleteEventV1) error {
	a.IsDelete = true
	return nil
}

func (a *{{.ClassName}}) GetAggregateRevision() string {
	return {{.Aggregate}}
}

func (a *{{.ClassName}}) GetAggregateType() string {
	return "{{.AggregateType}}"
}

func (a *{{.ClassName}}) GetAggregateId() string {
	return a.{{.Aggregate.Id.Name}}
}

func (a *{{.ClassName}}) GetTenantId() string {
	return a.TenantId
}
