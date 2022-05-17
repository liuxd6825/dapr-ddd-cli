package event
{{$namespace := .Namespace}}
import (
{{- range $name, $agg := .Aggregates}}
    "{{$namespace}}/pkg/cmd-service/domain/event/{{$agg.LowerName}}_events"
{{- end}}
    "github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventTypes
// @Description: 获取所有注册事件类型
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventTypes() *[]restapp.RegisterEventType {
    list := []restapp.RegisterEventType{}
{{- range $name, $agg := .Aggregates}}
    list = append(list, {{$agg.LowerName}}_events.GetRegisterEventTypes()...)
{{- end}}
    return &list
}

