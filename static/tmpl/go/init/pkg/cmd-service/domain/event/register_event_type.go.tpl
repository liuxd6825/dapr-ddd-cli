package event
{{$namespace := .Namespace}}
import (
{{- range $name, $agg := .Aggregates}}
    "{{$namespace}}/pkg/cmd-service/domain/event/{{$agg.SnakeName}}_event"
{{- end}}
    "github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventType
// @Description: 获取所有注册事件类型
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventType() *[]restapp.RegisterEventType {
    var list []restapp.RegisterEventType
{{- range $name, $agg := .Aggregates}}
    list = append(list, {{$agg.SnakeName}}_event.GetRegisterEventTypes()...)
{{- end}}
    return &list
}

