package view

import (
	"time"
)

//
// BaseView
// @Description: 视图基类
//
type BaseView struct {
{{- range $name, $property := .Properties}}
{{- if not $property.IsArray}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}},omitempty"  bson:"{{$property.BsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }} // {{$property.Description}}{{ end }}
{{- end}}
{{- end}}
}

{{- range $name, $property := .Properties}}
{{- if not $property.IsArray}}

//
// Get{{$property.UpperName}}
// @Description: 获取 {{$property.Description}}
//
func (v *BaseView) Get{{$property.UpperName}}(){{$property.LanType}} {
    return v.{{$property.UpperName}}
}

//
// Set{{$property.UpperName}}
// @Description: 设置 {{$property.Description}}
//
func (v *BaseView) Set{{$property.UpperName}}(value {{$property.LanType}}) {
    v.{{$property.UpperName}} = value
}
{{- end}}
{{- end}}