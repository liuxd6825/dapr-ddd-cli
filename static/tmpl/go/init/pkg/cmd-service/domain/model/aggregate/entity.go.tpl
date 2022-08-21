package model
{{- $ClassName := .ClassName }}
{{- $EventPackage := .EventPackage}}
{{- $CommandPackage := .CommandPackage}}

import (
    {{.time}}
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
)

//
// {{.ClassName}}
// @Description: {{.Description}} 实体类型
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.GoLanType}} `json:"{{$property.JsonName}}"  validate:"{{$property.GetValidate}}"` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

//
// New{{.ClassName}}
// @Description: 新建{{.Description}}对象
//
func New{{.ClassName}}()*{{.ClassName}}{
    return &{{.ClassName}}{}
}

//
// GetId
// @Description: 取ID值
//
func (e *{{.ClassName}}) GetId() string {
    return e.Id
}