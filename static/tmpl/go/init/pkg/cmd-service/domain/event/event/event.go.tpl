package event

import (
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
    "time"
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
    EventId       string           `json:"eventId"`       // 领域事件ID
    CommandId     string           `json:"commandId"`     // 关联命令ID
    CreatedTime   time.Time        `json:"time"`          // 事件创建时间
    {{- if .Event.IsUpdate }}
	UpdateMask   []string         `json:"updateMask"`    // 要更新字段
    {{- end }}
{{- if .HasDataProperty }}
    Data      field.{{.FieldName}}  `json:"data"`        // 业务字段项
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

func (e *{{.ClassName}}) GetCreatedTime() time.Time {
    return e.CreatedTime
}

func (e *{{.ClassName}}) GetData() interface{} {
    return e.Data
}
