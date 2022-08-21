package event

import (
    "time"
    "fmt"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
    "github.com/liuxd6825/dapr-go-ddd-sdk/utils/stringutils"
)

//
// {{.Name}}
// @Description: {{.Description}}
//
type {{.Name}} struct {
    EventId       string           `json:"eventId" validate:"required"`       // 领域事件ID
    CommandId     string           `json:"commandId" validate:"required"`     // 关联命令ID
    CreatedTime   time.Time        `json:"time" validate:"required"`          // 事件创建时间
    Version       string           `json:"version" validate:"required"`       // 事件版本
    EventType     string           `json:"eventType" validate:"required"`     // 事件类型
    {{- if .Event.IsUpdate }}
	UpdateMask    []string          `json:"updateMask" validate:"-"`           // 更新字段
    {{- end }}
{{- if .HasDataProperty }}
    Data          field.{{.FieldName}}  `json:"data" validate:"required"`      // 业务字段项
{{- else }}
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]{{if $property.IsObjectType}}*{{end}}{{end}}{{$property.GoLanType}} `json:"{{$property.JsonName}}" validate:"{{$property.GetValidate}}"` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end }}
{{- end }}
}

func New{{.Name}}(commandId string) *{{.Name}} {
    return &{{.Name}}{
        EventId    : fmt.Sprintf("%s(%s)", commandId, {{.EventType}}Type.String()),
        CommandId  : commandId,
        Version    : "{{.Version}}",
        EventType  : {{.EventType}}Type.String(),
        CreatedTime: time.Now(),
    }
}

func (e *{{.Name}}) GetEventId() string {
    return e.EventId
}

func (e *{{.Name}}) GetEventType() string {
    return e.EventType
}

func (e *{{.Name}}) GetEventVersion() string {
    return e.Version
}

func (e *{{.Name}}) GetCommandId() string {
    return e.CommandId
}

func (e *{{.Name}}) GetTenantId() string {
    return e.Data.TenantId
}

func (e *{{.Name}}) GetAggregateId() string {
    {{- if .IsEntity }}
    return e.Data.{{.AggregateName}}Id
    {{- else }}
    return e.Data.Id
    {{- end }}
}

func (e *{{.Name}}) GetCreatedTime() time.Time {
    return e.CreatedTime
}

func (e *{{.Name}}) GetData() interface{} {
    return e.Data
}
