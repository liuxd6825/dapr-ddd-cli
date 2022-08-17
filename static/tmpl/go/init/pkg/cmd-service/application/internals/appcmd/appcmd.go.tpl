package appcmd

import (
	domain "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
)
{{- range $i, $cmd := .Commands}}
//
// {{$cmd.AppName}}
// @Description: 应用服务层命令, {{$cmd.Description}}
//
type {{$cmd.AppName}} struct {
	domain.{{$cmd.Name}}
}
{{- end }}