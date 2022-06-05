package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/assembler"
)

var {{.Name}}Assembler = assembler.{{.Name}}


type {{.Name}}QueryApi struct {
	queryService *service.{{.Name}}QueryAppService
}


func New{{.Name}}QueryApi() *{{.Name}}QueryApi {
	return &{{.Name}}QueryApi{
		queryService: service.New{{.Name}}QueryAppService(),
	}
}


func (a *{{.Name}}QueryApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{{.EntityPluralName}}/{id}", "FindById")
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{{.ParentId}}/{{.EntityPluralName}}", "FindBy{{.AggregateName}}Id")
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{{.EntityPluralName}}:{{.aggregateMidlineName}}-id/{{.ParentId}}", "FindBy{{.Name}}Id")
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{{.EntityPluralName}}:all", "FindAll")
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{{.EntityPluralName}}", "FindPaging")
}


func (a *{{.Name}}QueryApi) FindById(ctx iris.Context, tenantId, id string) {
	req, err := {{.Name}}Assembler.AssFindByIdRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQueryOne(ctx, func(c context.Context) (interface{}, bool, error) {
		v, b, e := a.queryService.FindById(c, req.TenantId(), req.Id())
		return {{.Name}}Assembler.AssOneResponse(ctx, v, b, e)
	})
}


func (a *{{.Name}}QueryApi) FindBy{{.AggregateName}}Id(ctx iris.Context, tenantId, {{.aggregateName}}Id string) {
	req, err := {{.Name}}Assembler.AssFindBy{{.AggregateName}}IdRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		v, b, e := a.queryService.FindBy{{.AggregateName}}Id(c, req.TenantId(), req.{{.AggregateName}}Id)
		return {{.Name}}Assembler.AssListResponse(ctx, v, b, e)
	})
}


func (a *{{.Name}}QueryApi) FindAll(ctx iris.Context, tenantId string) {
	req, err := {{.Name}}Assembler.AssFindAllRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		fpr, b, e := a.queryService.FindAll(c, req.TenantId())
		return {{.Name}}Assembler.AssListResponse(ctx, fpr, b, e)
	})
}


func (a *{{.Name}}QueryApi) FindPaging(ctx iris.Context, tenantId string) {
	req, err := {{.Name}}Assembler.AssGetPagingRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		fpr, b, e := a.queryService.FindPaging(c, req)
		return {{.Name}}Assembler.AssFindPagingResponse(ctx, fpr, b, e)
	})
}

