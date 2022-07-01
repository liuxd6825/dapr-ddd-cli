package model

import (
    {{- if .Properties.HasDateTimeType }}
    "time"
    {{- end}}
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"  validate:"{{$property.GetValidate}}" ` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{}
}
