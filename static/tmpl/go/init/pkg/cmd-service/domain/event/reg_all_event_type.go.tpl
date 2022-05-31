package event
{{$namespace := .Namespace}}
import (
{{- range $name, $agg := .Aggregates}}
    "{{$namespace}}/pkg/cmd-service/domain/event/{{$agg.SnakeName}}_event"
{{- end}}
    "github.com/dapr/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventTypes
// @Description: 获取所有注册事件类型
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventTypes() *[]restapp.RegisterEventType {
    var list []restapp.RegisterEventType
{{- range $name, $agg := .Aggregates}}
    list = append(list, {{$agg.SnakeName}}_event.GetRegisterEventTypes()...)
{{- end}}
    return &list
}

