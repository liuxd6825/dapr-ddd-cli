package controller
{{$Namespace := .Namespace}}
import (
    {{range $aggName, $agg := .Aggregates}}
    "{{$Namespace}}/pkg/query-service/userinterface/rest/controller/{{$agg.FileName}}_controller"
    {{- end }}
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterController
// @Description: 注册的控制器
// @return *[]restapp.Controller
//
func GetRegisterController() *[]restapp.Controller {
	var list []restapp.Controller

	{{- range $aggName, $agg := .Aggregates}}
	list = append(list, {{$agg.FileName}}_controller.New{{$agg.Name}}Controller())
	{{- range $entityName, $entity := $agg.Entities}}
    list = append(list, {{$agg.FileName}}_controller.New{{$entity.Name}}Controller())
	{{- end }}
	{{- end }}

	return &list
}
