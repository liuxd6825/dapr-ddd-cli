package {{.Aggregate.LowerName}}_events

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventTypes
// @Description: 获取聚合根注册事件类型
// @return []restapp.RegisterEventType
//
func GetRegisterEventTypes() []restapp.RegisterEventType {
    return []restapp.RegisterEventType{
{{- range $eventName, $event := .Events}}
        {
            EventType: {{$eventName}}.String(),
            Revision:  (&{{$eventName}}{}).GetEventRevision(),
            NewFunc:   func() interface{} { return &{{$eventName}}{} },
        },
{{- end}}
    }
}
