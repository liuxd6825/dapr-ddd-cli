package field

import (
    "time"
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
{{- if .IsEntity}}
    {{.AggregateName}}Id string `json:"{{.aggregateName}}Id" validate:"gt=0"` // 聚合根Id
{{- end }}
}

func (f *{{.ClassName}}) GetId() string {
    return f.Id
}

func (f *{{.ClassName}}) GetTenantId() string {
    return f.TenantId
}