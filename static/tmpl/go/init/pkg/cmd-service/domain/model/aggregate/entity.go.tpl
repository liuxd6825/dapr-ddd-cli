package model

import (
    {{- if .Properties.HasDateTimeType }}
    "time"
    {{- end}}
)

//
// {{.ClassName}}
// @Description: {{.Description}} 实体类型
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
    {{.AggregateName}}Id string `json:"{{.aggregateName}}Id" bson:"{{.aggregate_name}}_id" `   // 聚合根Id
}

//
// New{{.ClassName}}
// @Description: 新建{{.Description}}对象
//
func New{{.ClassName}}()*{{.ClassName}}{
    return &{{.ClassName}}{}
}
