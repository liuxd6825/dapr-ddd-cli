package model

import (
    "time"
)

//
// {{.ClassName}}
// @Description: {{.Description}} 实体类型
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

//
// New{{.ClassName}}
// @Description: 新建{{.Description}}对象
//
func New{{.ClassName}}()*{{.ClassName}}{
    return &{{.ClassName}}{}
}
