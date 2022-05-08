package {{.Package}}

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
)

type {{.ClassName}} struct {
	TenantId  string                 `json:"tenantId"`
	CommandId string                 `json:"commandId"`
	EventId   string                 `json:"eventId"`
{{- range $name, $property := .Properties}}
    # {{$property.Description}}
    {{$property.UpperName}}   {{$property.DataType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`
{{- end}}
}

func New{{.ClassName}}() *{{.ClassName}} {
	return &{{.ClassName}}{}
}

func (e *{{.ClassName}}) GetEventId() string {
	return e.EventId
}

func (e *{{.ClassName}}) GetEventType() string {
	return UserCreateEventType.String()
}

func (e *{{.ClassName}}) GetEventRevision() string {
	return "1.0"
}

func (e *{{.ClassName}}) GetCommandId() string {
	return e.CommandId
}

func (e *{{.ClassName}}) GetTenantId() string {
	return e.TenantId
}

func (e *{{.ClassName}}) GetAggregateId() string {
	return e.Data.Id
}
