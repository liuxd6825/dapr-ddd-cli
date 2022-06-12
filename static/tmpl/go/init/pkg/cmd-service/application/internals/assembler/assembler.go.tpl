package assembler

import (
	"context"
	"errors"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/dto"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
)

{{- range $i, $cmd := .Commands}}

//
// Ass{{$cmd.Name}}
// @Description:  {{$cmd.Name}}Dto {{$cmd.Description}}转换器
// @param ctx
// @param cmdDto
// @return *command.{{$cmd.Name}}
// @return error
//
func Ass{{$cmd.Name}}(ctx context.Context, cmdDto *dto.{{$cmd.Name}}Dto) (*command.{{$cmd.Name}}, error) {
    if cmdDto == nil {
        return nil, errors.New("cmdDto is nil")
    }
	o := *cmdDto
	cmd := o.{{$cmd.Name}}
	return &cmd, nil
}

{{- end }}

