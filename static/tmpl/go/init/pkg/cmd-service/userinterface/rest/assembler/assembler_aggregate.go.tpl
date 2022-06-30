package assembler

import (
	"github.com/kataras/iris/v12"
	adto "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/dto"
	"{{.Namespace}}/pkg/cmd-service/infrastructure/utils"
	"{{.Namespace}}/pkg/cmd-service/userinterface/rest/{{.aggregate_name}}/dto"
)

//
// {{.AggregateName}}Assembler
// @Description: {{.Aggregate.Description}}
//
type {{.AggregateName}}Assembler struct {
}

{{- $AggregateName := .AggregateName}}
var {{.AggregateName}} = &{{.AggregateName}}Assembler{}
{{- range $cmdName, $cmd := .Commands}}

//
// Ass{{$cmd.Name}}Dto
// @Description: {{$cmd.Description}}
// @receiver a
// @param ictx
// @return *adto.{{$cmd.Name}}Dto {{$cmd.Description}} 应用层DTO对象
// @return error 错误
//
func (a *{{$AggregateName}}Assembler) Ass{{$cmd.Name}}Dto(ictx iris.Context) (*adto.{{$cmd.Name}}Dto, error) {
	var request dto.{{$cmd.Name}}Request
	var cmd adto.{{$cmd.Name}}Dto
	if err := utils.AssemblerRequestBody(ictx, &request, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}
{{- end }}

