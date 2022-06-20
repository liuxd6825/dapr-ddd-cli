{{- $ClassName := .ClassName}}
package view

import (
    "time"
    base "{{.Namespace}}/pkg/query-service/infrastructure/base/domain/view"
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
    base.BaseView `bson:",inline"`
{{- range $name, $property := .Properties}}
{{- if not $property.IsArray}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}},omitempty"  bson:"{{$property.BsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }} // {{$property.Description}}{{ end }}
{{- end}}
{{- end}}
}


//
// New{{.ClassName}}
// @Description: {{.Description}}
//
func New{{.ClassName}}()*{{.ClassName}}{
    return &{{.ClassName}}{}
}



