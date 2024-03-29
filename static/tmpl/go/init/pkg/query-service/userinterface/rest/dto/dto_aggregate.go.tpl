package dto

import (
    "github.com/kataras/iris/v12"
    "github.com/liuxd6825/dapr-go-ddd-sdk/assert"
    "github.com/liuxd6825/dapr-go-ddd-sdk/types"
    "{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	base "{{.Namespace}}/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

// 按ID查询

//
// {{.Name}}FindByIdResponse
// @Description:  按ID查询响应体
//
type {{.Name}}FindByIdResponse struct {
    {{.Name}}Dto
}

func New{{.Name}}FindByIdResponse() *{{.Name}}FindByIdResponse {
	return &{{.Name}}FindByIdResponse{}
}

// 按多个ID查询

//
// {{.Name}}FindByIdsResponse
// @Description: {{.Description}}  查询所有响应体
//
type {{.Name}}FindByIdsResponse []*{{.Name}}FindByIdsResponse

func New{{.Name}}FindByIdsResponse() *{{.Name}}FindByIdsResponse {
	return &{{.Name}}FindByIdsResponse{}
}

//
// {{.Name}}FindByIdsResponseItem
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}FindByIdsResponseItem struct {
    {{.Name}}Dto
}

func New{{.Name}}FindByIdsResponseItem() *{{.Name}}FindByIdsResponseItem {
	return &{{.Name}}FindByIdsResponseItem{}
}

// 分页查询

//
// {{.Name}}FindPagingRequest
// @Description: 分页请求数据
//
type {{.Name}}FindPagingRequest struct {
	base.FindPagingRequest
}

func New{{.Name}}FindPagingRequest() *{{.Name}}FindPagingRequest {
	return &{{.Name}}FindPagingRequest{}
}

//
// {{.Name}}FindPagingResponse
// @Description: {{.Name}} 分页响应数据
//
type {{.Name}}FindPagingResponse struct {
	base.FindPagingResponse[*{{.Name}}FindPagingResponseItem]
}

func New{{.Name}}FindPagingResponse() *{{.Name}}FindPagingResponse {
	resp := &{{.Name}}FindPagingResponse{}
	resp.Init()
	return resp
}

//
// {{.Name}}FindPagingResponseItem
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}FindPagingResponseItem struct {
    {{.Name}}Dto
}

func New{{.Name}}FindPagingResponseItem() *{{.Name}}FindPagingResponseItem {
	return &{{.Name}}FindPagingResponseItem{}
}

// 查询所有

//
// {{.Name}}FindAllResponse
// @Description: {{.Description}}  查询所有响应体
//
type {{.Name}}FindAllResponse []*{{.Name}}FindAllResponseItem

func New{{.Name}}FindAllResponse() *{{.Name}}FindAllResponse {
	return &{{.Name}}FindAllResponse{}
}

//
// {{.Name}}FindAllResponseItem
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}FindAllResponseItem struct {
    {{.Name}}Dto
}

func New{{.Name}}FindAllResponseItem() *{{.Name}}FindAllResponseItem {
	return &{{.Name}}FindAllResponseItem{}
}

//
// {{.Name}}Dto
// @Description: {{.Description}} 请求或响应业务数据
//
type {{.Name}}Dto struct {
    base.BaseDto
{{- range $name, $property := .Properties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} []*{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}""`  // {{$property.Description}}
    {{- else if $property.IsEntityType}}
    {{$property.UpperName}} *{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}""`  // {{$property.Description}}
    {{- else if $property.IsDates }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsTimes }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEnumType }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.GoLanType}}`json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- end}}
{{- end}}
}

func New{{.Name}}Dto() *{{.Name}}Dto {
	return &{{.Name}}Dto{}
}

