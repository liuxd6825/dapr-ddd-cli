package service

import (
	"context"
    "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/dto"
    "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/assembler"
)
{{- $ClassName := .ClassName}}
{{- $CommandPackage := .CommandPackage}}
{{- range $i, $cmd := .Commands}}

//
// {{$cmd.ServiceFuncName}}
// @Description: {{$cmd.Description}}
// @receiver s
// @param ctx 上下文
// @param cmd {{$cmd.Description}}
// @return error
//
func (s *{{$ClassName}}) {{$cmd.ServiceFuncName}}(ctx context.Context, cmdDto *dto.{{$cmd.Name}}Dto) error {
	cmd, err := assembler.Ass{{$cmd.Name}}(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = s.domainService.{{$cmd.ServiceFuncName}}(ctx, cmd)
	return err
}
{{- end }}

