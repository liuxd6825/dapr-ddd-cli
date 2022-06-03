package dto

import (
    base "{{.Namespace}}/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

//
// {{.Name}}GetByIdRequest
// @Description: {{.Description}} 请求内容
//
type {{.Name}}GetByIdRequest struct {
    base.GetByIdRequest
    {{.Name}}ViewDto
}

//
// {{.Name}}GetByIdResponse
// @Description: {{.Description}} 请求内容
//
type {{.Name}}GetByIdResponse struct {
    base.GetByIdResponse
}


//
// {{.Name}}GetAllRequest
// @Description: {{.Description}}
//
type {{.Name}}GetAllRequest struct {
    base.GetAllRequest
}


//
// {{.Name}}GetPagingRequest
// @Description: {{.Description}}
//
type {{.Name}}GetPagingRequest struct {
    base.GetPagingRequest
}


//
// {{.Name}}GetPagingResponse
// @Description: {{.Description}}
//
type {{.Name}}GetPagingResponse struct {
    base.GetPagingResponse[*{{.Name}}ViewDto]
}



//
// {{.Name}}ViewDto
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}ViewDto struct {
{{- range $name, $property := .DataFieldsProperties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}


