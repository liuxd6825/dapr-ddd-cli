package register
{{$Namespace := .Namespace}}
import (
    {{range $aggName, $agg := .Aggregates}}
    "{{$Namespace}}/pkg/query-service/application/internals/handler/{{$agg.FileName}}_handler"
    {{- end }}
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterSubscribe
// @Description: 注册领域事件监听器
// @return *[]restapp.RegisterSubscribe
//
func GetRegisterSubscribe() *[]restapp.RegisterSubscribe {
	var list []restapp.RegisterSubscribe

	{{- range $aggName, $agg := .Aggregates}}
	list = append(list, {{$agg.FileName}}_handler.New{{$agg.Name}}Subscribe())
	{{- range $entityName, $entity := $agg.Entities}}
    list = append(list, {{$agg.FileName}}_handler.New{{$entity.Name}}Subscribe())
	{{- end }}
	{{- end }}

	return &list
}
