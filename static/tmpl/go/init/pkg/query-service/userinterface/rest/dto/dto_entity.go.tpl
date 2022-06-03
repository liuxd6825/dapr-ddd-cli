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
}


//
// {{.Name}}GetByIdResponse
// @Description: {{.Description}} 请求内容
//
type {{.Name}}GetByIdResponse {{.Name}}ViewDto



//
// {{.Name}}GetAllRequest
// @Description: {{.Description}} List
//
type {{.Name}}GetAllRequest struct {
    base.GetAllRequest
}


//
// {{.Name}}GetBy{{.AggregateName}}IdRequest
// @Description: {{.Description}} List
//
type {{.Name}}GetBy{{.AggregateName}}IdRequest struct {
    base.GetRequest
    {{.AggregateName}}Id string `json:"{{.aggregateName}}Id"`
}


//
// {{.Name}}GetBy{{.AggregateName}}IdResponse
// @Description: {{.Description}} List
//
type {{.Name}}GetBy{{.AggregateName}}IdResponse []*{{.Name}}ViewDto

//
// {{.Name}}GetPagingRequest
// @Description: {{.Description}} List
//
type {{.Name}}GetPagingRequest struct {
    base.GetPagingRequest
}


//
// {{.Name}}GetPagingResponse
// @Description: {{.Description}} List
//
type {{.Name}}GetPagingResponse struct {
    base.GetPagingResponse[*{{.Name}}ViewDto]
}


//
// {{.Name}}ListDto
// @Description: {{.Description}} List
//
type {{.Name}}ListDto []*{{.Name}}ViewDto


//
// {{.Name}}ViewDto
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}ViewDto struct {
{{- range $name, $property := .DataFieldsProperties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}


