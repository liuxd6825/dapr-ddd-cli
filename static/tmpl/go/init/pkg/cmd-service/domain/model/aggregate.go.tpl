package model

import (
	"context"
	user_commands2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/factory/user_factory"
	user_events2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
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

func (a *{{.ClassName}}) AddressCreateCommand(ctx context.Context, cmd *user_commands2.AddressCreateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewAddressCreateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) AddressUpdateCommand(ctx context.Context, cmd *user_commands2.AddressUpdateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewAddressUpdateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) AddressDeleteCommand(ctx context.Context, cmd *user_commands2.AddressDeleteCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewAddressDeleteEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) UserCreateCommand(ctx context.Context, cmd *user_commands2.UserCreateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewCreateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) UserUpdateCommand(ctx context.Context, cmd *user_commands2.UserUpdateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewUpdateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) UserDeleteCommand(ctx context.Context, cmd *user_commands2.UserDeleteCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewDeleteEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *{{.ClassName}}) OnUserAddressCreateEventV1s0(ctx context.Context, event *user_events2.AddressCreateEventV1) (err error) {
	fields := &event.Data
	a.Addresses[fields.Id] = NewAddressValue(fields)
	return nil
}

func (a *{{.ClassName}}) OnUserAddressUpdateEventV1s0(ctx context.Context, event *user_events2.AddressUpdateEventV1) error {
	fields := &event.Data
	a.Addresses[fields.Id] = NewAddressValue(fields)
	return nil
}

func (a *{{.ClassName}}) OnUserAddressDeleteEventV1s0(ctx context.Context, event *user_events2.AddressDeleteEventV1) error {
	delete(a.Addresses, event.AddressId)
	return nil
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
	return "1.0"
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
