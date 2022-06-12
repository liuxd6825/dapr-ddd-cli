package dto

import (
    "github.com/kataras/iris/v12"
    "github.com/liuxd6825/dapr-go-ddd-sdk/assert"
    "{{.Namespace}}/pkg/query-service/infrastructure/types"
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


// 按{{.Name}}Id查询

//
// {{.Name}}FindBy{{.AggregateName}}IdRequest
// @Description: {{.Description}}  查询{{.Name}}Id请求体
//
type {{.Name}}FindBy{{.AggregateName}}IdRequest struct {
    TenantId string
    {{.AggregateName}}Id string
}

//
// {{.Name}}FindBy{{.AggregateName}}IdResponse
// @Description: {{.Description}}  查询{{.Name}}Id响应体
//
type {{.Name}}FindBy{{.AggregateName}}IdResponse []*{{.Name}}FindBy{{.AggregateName}}IdResponseItem

//
// {{.Name}}FindBy{{.AggregateName}}IdResponseItem
// @Description: {{.Description}}  请求{{.Name}}Id响应体
//
type {{.Name}}FindBy{{.AggregateName}}IdResponseItem struct {
    {{.Name}}Dto
}

// 分页查询

//
// {{.Name}}FindPagingResponse
// @Description: {{.Name}} 分页请求数据
//
type {{.Name}}FindPagingResponse struct {
	base.FindPagingResponse[*{{.Name}}FindPagingResponseItem]
}


//
// {{.Name}}FindPagingResponseItem
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}FindPagingResponseItem struct {
    {{.Name}}Dto
}


// 查询所有

//
// {{.Name}}FindAllResponse
// @Description: {{.Description}}  查询所有响应体
//
type {{.Name}}FindAllResponse []*{{.Name}}FindAllResponseItem

//
// {{.Name}}FindAllResponseItem
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}FindAllResponseItem struct {
    {{.Name}}Dto
}


//
// {{.Name}}Dto
// @Description: {{.Description}}  响应业务数据
//
type {{.Name}}Dto struct {
{{- range $name, $property := .DataFieldsProperties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}},omitempty"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}

