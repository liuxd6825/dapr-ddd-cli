{{- $AggregateName := .AggregateName}}
package assembler

import (
	"github.com/kataras/iris/v12"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/appcmd"
	"{{.Namespace}}/pkg/cmd-service/infrastructure/utils"
	"{{.Namespace}}/pkg/cmd-service/userinterface/rest/{{.aggregate_name}}/dto"
)

{{- if .IsAggregate }}
type {{$AggregateName}}Assembler struct {
}
{{- end }}

{{- range $cmdName, $cmd := .Commands}}

//
// Ass{{$cmd.Name}}Dto
// @Description: {{$cmd.Description}}
// @receiver a
// @param ictx
// @return *appcmd.{{$cmd.AppName}} {{$cmd.Description}} 应用层DTO对象
// @return error 错误
//
func (a *{{$AggregateName}}Assembler) Ass{{$cmd.AppName}}(ictx iris.Context) (*appcmd.{{$cmd.AppName}}, error) {
	var request dto.{{$cmd.Name}}Request
	var appCmd appcmd.{{$cmd.AppName}}
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

{{- end }}

