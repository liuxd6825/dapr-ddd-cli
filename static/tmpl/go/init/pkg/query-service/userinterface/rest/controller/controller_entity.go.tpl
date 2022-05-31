package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	app_service "{{.Namespace}}/pkg/query-service/application/internales/service/{{.aggregate_name}}_service"
	"github.com/dapr/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}Controller struct {
	appService *app_service.{{.Name}}QueryAppService
}

func New{{.Name}}Controller() *{{.Name}}Controller {
	return &{{.Name}}Controller{
		appService: app_service.New{{.Name}}QueryAppService(),
	}
}

func (c *{{.Name}}Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/{{.ResourceName}}/{id}", "GetById")
	b.Handle("GET", "/tenants/{tenantId}/{{.ResourceName}}/{{.ParentResourceName}}/{{.ParentId}}", "GetBy{{.AggregateName}}Id")
	b.Handle("GET", "/tenants/{tenantId}/{{.ResourceName}}", "GetPagingData")

}

func (c *{{.Name}}Controller) GetById(ctx iris.Context, tenantId, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return c.appService.FindById(ctx, tenantId, id)
	})
}

func (c *{{.Name}}Controller) GetBy{{.AggregateName}}Id(ctx iris.Context, tenantId, {{.aggregateName}}Id string) {
	_, _, _ = restapp.DoQuery(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return c.appService.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
	})
}

func (c *{{.Name}}Controller) GetPagingData(ctx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ctx, func(ctx context.Context) (interface{}, bool, error) {
		query, err := restapp.NewListQuery(ctx, tenantId)
		if err != nil {
			return nil, false, err
		}
		return c.appService.FindPagingData(ctx, query)
	})
}
