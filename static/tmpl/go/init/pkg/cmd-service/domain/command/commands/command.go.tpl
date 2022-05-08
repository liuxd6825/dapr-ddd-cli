package {{.Package}}

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
)

type {{.ClassName}} struct {
	ddd.BaseCommand
{{- range $name, $property := .Properties}}
    # {{$property.Description}}
    {{$property.UpperName}}   {{$property.DataType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`
{{- end}}
}

func (c *{{.ClassName}}) NewDomainEvent() ddd.DomainEvent {
	return &user_events.UserUpdateEventV1{
		EventId: c.CommandId,
		Data:    c.Data,
	}
}

func (c *{{.ClassName}}) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.{{.AggregateId}})
}

func (c *{{.ClassName}}) GetCommandId() string {
	return c.CommandId
}

func (c *{{.ClassName}}) GetTenantId() string {
	return c.Data.TenantId
}

func (c *{{.ClassName}}) Validate() error {
	errs := ddd_errors.NewVerifyError()
	c.BaseCommand.ValidateError(errs)
	if c.Data.TenantId == "" {
		errs.AppendField("data.tenantId", "不能为空")
	}
	return errs.GetError()
}
