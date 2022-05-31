package {{.aggregate_name}}_event

import (
    "github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)
{{- $AggregateName := .AggregateName}}

type {{$AggregateName}}EventType uint32

const (
{{ $zero:=0 }}
{{- range $typeName := .EventTypes}}
	{{$typeName}}Type {{if eq $zero 0 }} {{$AggregateName}}EventType = iota{{ end }}
	{{- $zero = add $zero 1 }}
{{- end }}
)

func (p {{$AggregateName}}EventType) String() string {
	switch p {
{{- $ServiceName := .ServiceName}}
{{- range $typeName := .EventTypes}}
    case {{$typeName}}Type:
        return "{{$ServiceName}}.{{$typeName}}"
{{- end }}
    default:
        return "UNKNOWN"
	}
}

//
// GetRegisterEventTypes
// @Description: 获取聚合根注册事件类型
// @return []restapp.RegisterEventType
//
func GetRegisterEventTypes() []restapp.RegisterEventType {
    return []restapp.RegisterEventType{
{{- range $eventName, $event := .Events}}
        {
            EventType: {{$event.EventType}}Type.String(),
            Version:  (&{{$eventName}}{}).GetEventVersion(),
            NewFunc:   func() interface{} { return &{{$eventName}}{} },
        },
{{- end}}
    }
}
