package {{.Package}}

# {{.Description}}
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    # {{$property.Description}}
    {{$property.UpperName}}   {{$property.DataType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`
{{- end}}
}

