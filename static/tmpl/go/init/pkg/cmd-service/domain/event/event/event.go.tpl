package {{.Package}}

import (
    "{{.Namespace}}/pkg/cmd-service/domain/fields/{{.AggregateFieldPackage}}"
)

type {{.ClassName}} struct {
    CommandId string    `json:"commandId"`
    EventId   string    `json:"eventId"`
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.DataType}} `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}//{{$property.Description}}{{ end }}
{{- end}}
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
    return "1.0"
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
