package register
{{$Namespace := .Namespace}}
import (
    {{range $aggName, $agg := .Aggregates}}
    {{$agg.FileName}}_api "{{$Namespace}}/pkg/cmd-service/userinterface/rest/{{$agg.FileName}}/facade"
    {{- end }}
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterController
// @Description: 注册的控制器
// @return *[]restapp.Controller
//
func GetRegisterController() []restapp.Controller {
	var list []restapp.Controller
	{{- range $aggName, $agg := .Aggregates}}
	list = append(list, {{$agg.FileName}}_api.New{{$agg.Name}}CommandApi())
    {{- range $entityName, $entity := $agg.Entities}}
    list = append(list, {{$agg.FileName}}_api.New{{$entity.Name}}CommandApi())
    {{- end }}
	{{- end }}
	return list
}
