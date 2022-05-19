package {{.aggregateName}}_view

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.DataType}} `json:"{{$property.JsonName}}"  bson:"{{$property.BsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

func (v *{{.ClassName}}) GetTenantId() string {
	return v.TenantId
}
func (v *{{.ClassName}}) GetId() string {
	return v.Id
}