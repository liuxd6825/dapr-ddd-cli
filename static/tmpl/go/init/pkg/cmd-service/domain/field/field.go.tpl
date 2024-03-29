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
    {{$property.UpperName}} {{if $property.IsArray}}[]{{if $property.IsObjectType}}*{{end}}{{end}}{{$property.LanType}} `{{$property.GoTags}}` {{$property.GoDescription}}
{{- end}}
}

{{- if not .IsItems }}
func (f *{{.ClassName}}) GetId() string {
    return f.Id
}

func (f *{{.ClassName}}) GetTenantId() string {
    return f.TenantId
}
{{- else }}

func (f *{{.ClassName}}) GetIds() []string {
    var ids []string
    ids =[]string{}
    for _, item := range f.Items {
        ids = append(ids, item.GetId())
    }
    return ids
}
{{- end }}