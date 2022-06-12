package dto

import (
	domain "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
)
{{- range $i, $cmd := .Commands}}

//
// {{$cmd.Name}}Dto
// @Description:  {{$cmd.Description}}传输对象
//
type {{$cmd.Name}}Dto struct {
	domain.{{$cmd.Name}}
}
{{- end }}