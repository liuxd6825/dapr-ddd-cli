package model

import (
    {{.time}}
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.GoLanType}} `json:"{{$property.JsonName}}"  validate:"{{$property.GetValidate}}" ` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{}
}
