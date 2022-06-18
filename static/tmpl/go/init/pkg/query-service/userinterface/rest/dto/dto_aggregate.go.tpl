package dto

import (
    "github.com/kataras/iris/v12"
    "github.com/liuxd6825/dapr-go-ddd-sdk/assert"
    "github.com/liuxd6825/dapr-go-ddd-sdk/types"
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

// 分页查询

//
// {{.Name}}FindPagingResponse
// @Description: {{.Name}} 分页请求数据
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
// @Description: {{.Description}}  响应业务数据
//
type {{.Name}}Dto struct {
    base.BaseDto
{{- range $name, $property := .Properties}}
    {{- if $property.TypeIsDateTime }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.LanType}}`json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
    {{- end}}
{{- end}}
}

func New{{.Name}}Dto() *{{.Name}}Dto {
	return &{{.Name}}Dto{}
}

