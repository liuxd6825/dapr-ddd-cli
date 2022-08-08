package field

import (
    {{.time}}
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"  validate:"{{$property.GetValidate}}"` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

func (f *{{.ClassName}}) GetId() string {
    return f.Id
}

func (f *{{.ClassName}}) GetTenantId() string {
    return f.TenantId
}