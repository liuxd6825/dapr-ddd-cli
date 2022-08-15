package assembler

import (
	"context"
	"errors"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/appcmd"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
)

{{- range $i, $cmd := .Commands}}

//
// Ass{{$cmd.Name}}
// @Description:  {{$cmd.AppName}} {{$cmd.Description}}转换器
// @param ctx
// @param cmdDto
// @return *command.{{$cmd.Name}}
// @return error
//
func Ass{{$cmd.Name}}(ctx context.Context, appCmd *appcmd.{{$cmd.AppName}}) (*command.{{$cmd.Name}}, error) {
    if appCmd == nil {
        return nil, errors.New("Ass{{$cmd.Name}}() appCmd is nil")
    }
    cmd := (*appCmd).{{$cmd.Name}}
    return &cmd, nil
}

{{- end }}

