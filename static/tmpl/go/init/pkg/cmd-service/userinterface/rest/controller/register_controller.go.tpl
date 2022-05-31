package controller
{{$Namespace := .Namespace}}
import (
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
	list = append(list, New{{$agg.Name}}Controller())
	{{- end }}
	return &list
}
