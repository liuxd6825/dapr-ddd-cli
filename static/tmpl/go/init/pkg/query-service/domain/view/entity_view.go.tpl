{{- $ClassName := .ClassName}}
package view

import (
    {{- if .HasDateTimeType }}
    "time"
    {{- end}}
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
{{- if not $property.IsArray}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}},omitempty"  bson:"{{$property.BsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }} // {{$property.Description}}{{ end }}
{{- end}}
{{- end}}
{{- range $name, $property := .DefaultProperties}}
{{- if not $property.IsArray}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}},omitempty"  bson:"{{$property.BsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }} // {{$property.Description}}{{ end }}
{{- end}}
{{- end}}
}

//
// New{{.ClassName}}
// @Description: 创建 {{.Description}} 视图对象
//
func New{{.ClassName}}()*{{.ClassName}}{
    return &{{.ClassName}}{}
}

{{- range $name, $property := .DefaultViewProperties}}
{{- if not $property.IsArray}}

//
// Get{{$property.UpperName}}
// @Description: {{$property.Description}}
//
func (v *{{$ClassName}}) Get{{$property.Name}}() {{$property.LanType}} {
    return v.{{$property.UpperName}}
}

//
// Set{{$property.UpperName}}
// @Description: {{$property.Description}}
//
func (v *{{$ClassName}}) Set{{$property.Name}}(value {{$property.LanType}})  {
    v.{{$property.UpperName}} = value
}

{{- end}}
{{- end}}

