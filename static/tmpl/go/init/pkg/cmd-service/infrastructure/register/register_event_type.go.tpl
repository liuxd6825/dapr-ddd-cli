package register
{{$namespace := .Namespace}}
import (
{{- range $name, $agg := .Aggregates}}
    {{$agg.SnakeName}}_event "{{$namespace}}/pkg/cmd-service/domain/{{$agg.SnakeName}}/event"
{{- end}}
    "github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventType
// @Description: 获取所有注册事件类型
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventType() []restapp.RegisterEventType {
    var list []restapp.RegisterEventType
{{- range $name, $agg := .Aggregates}}
    list = append(list, {{$agg.SnakeName}}_event.GetRegisterEventTypes()...)
{{- end}}
    return list
}

