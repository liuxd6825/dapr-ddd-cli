package event

import (
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
)

type {{.ClassName}} struct {
    CommandId string               `json:"commandId"`
    EventId   string               `json:"eventId"`
{{- if .HasDataProperty }}
    Data      field.{{.FieldName}}  `json:"data"`
{{- else }}
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
{{- end }}
}

func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{}
}

func (e *{{.ClassName}}) GetEventId() string {
    return e.EventId
}

func (e *{{.ClassName}}) GetEventType() string {
    return {{.EventType}}Type.String()
}

func (e *{{.ClassName}}) GetEventVersion() string {
    return "{{.Version}}"
}

func (e *{{.ClassName}}) GetCommandId() string {
    return e.CommandId
}

func (e *{{.ClassName}}) GetTenantId() string {
    return e.Data.TenantId
}

func (e *{{.ClassName}}) GetAggregateId() string {
    return e.Data.Id
}
