package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	app_service "{{.Namespace}}/pkg/query-service/application/internals/service/{{.aggregate_name}}_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}QueryApi struct {
	appService *app_service.{{.Name}}QueryAppService
}

func New{{.Name}}QueryApi() *{{.Name}}QueryApi {
	return &{{.Name}}QueryApi{
		appService: app_service.New{{.Name}}QueryAppService(),
	}
}

func (c *{{.Name}}QueryApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{id}", "GetById")
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}:all", "GetAll")
	b.Handle("GET", "/tenants/{tenantId}/{{.AggregatePluralName}}", "GetPagingData")
}

func (c *{{.Name}}QueryApi) GetById(ctx iris.Context, tenantId, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return c.appService.FindById(ctx, tenantId, id)
	})
}

func (c *{{.Name}}QueryApi) GetAll(ctx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return c.appService.FindAll(ctx, tenantId)
	})
}

func (c *{{.Name}}QueryApi) GetPagingData(ctx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ctx, func(ctx context.Context) (interface{}, bool, error) {
		query, err := restapp.NewListQuery(ctx, tenantId)
		if err != nil {
			return nil, false, err
		}
		return c.appService.FindPagingData(ctx, query)
	})
}
