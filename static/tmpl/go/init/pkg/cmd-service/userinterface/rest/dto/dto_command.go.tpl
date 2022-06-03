package dto

//
// {{.Name}}RequestDto
// @Description: {{.Description}} 请求内容
//
type {{.Name}}RequestDto struct {
    CommandId   string `json:"commandId"  validate:"gt=0"`
    IsValidOnly bool   `json:"isValidOnly"`
    Data  {{.Name}}RequestDataDto `json:"data"`
}


//
// {{.Name}}RequestDataDto
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}RequestDataDto struct {
{{- range $name, $property := .DataFieldsProperties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}



//
// {{.Name}}ResponseDto
// @Description: {{.Description}} 响应内容
//
type {{.Name}}ResponseDto struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}