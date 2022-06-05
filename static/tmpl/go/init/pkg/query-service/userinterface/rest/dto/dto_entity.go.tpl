package dto


import (
    "github.com/kataras/iris/v12"
    "github.com/liuxd6825/dapr-go-ddd-sdk/assert"
    "{{.Namespace}}/pkg/query-service/infrastructure/types"
	base "{{.Namespace}}/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

//
// {{.Name}}FindByIdRequest
// @Description:  请求内容
//
type {{.Name}}FindByIdRequest struct {
	base.FindByIdRequest
}

//
// {{.Name}}FindByIdResponse
// @Description:  请求内容
//
type {{.Name}}FindByIdResponse struct {
	base.FindByIdResponse
	{{.Name}}ViewDto
}

//
// {{.Name}}FindBy{{.AggregateName}}IdRequest
// @Description:  请求内容
//
type {{.Name}}FindBy{{.AggregateName}}IdRequest struct {
	base.FindRequest
	{{.AggregateName}}Id string
}

func (r *{{.Name}}FindBy{{.AggregateName}}IdRequest) Init(ctx iris.Context)error{
    if err:= r.FindRequest.Init(ctx);err!=nil{
        return err
    }
	r.{{.AggregateName}}Id = ctx.Params().Get("{{.aggregateName}}Id")
	if err := assert.NotEmpty(r.{{.AggregateName}}Id, assert.NewOptions("url \"{{.aggregateName}}Id\" cannot be empty")); err != nil {
		return err
	}
	return nil
}


//
// {{.Name}}FindBy{{.AggregateName}}IdResponse
// @Description:  请求内容
//
type {{.Name}}FindBy{{.AggregateName}}IdResponse struct {
	base.FindResponse
	{{.Name}}ViewDto
}

//
// {{.Name}}FindAllRequest
// @Description:
//
type {{.Name}}FindAllRequest struct {
	base.FindAllRequest
}

//
// {{.Name}}FindPagingRequest
// @Description:
//
type {{.Name}}FindPagingRequest struct {
	base.FindPagingRequest[*{{.Name}}ViewDto]
}

//
// {{.Name}}FindPagingResponse
// @Description:
//
type {{.Name}}FindPagingResponse struct {
	base.FindPagingResponse
}

//
// {{.Name}}ViewList
// @Description: {{.Description}}  请求业务数据列表
//
type {{.Name}}ViewList *[]*{{.Name}}ViewDto


//
// {{.Name}}ViewDto
// @Description: {{.Description}}  请求业务数据
//
type {{.Name}}ViewDto struct {
{{- range $name, $property := .DataFieldsProperties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}
