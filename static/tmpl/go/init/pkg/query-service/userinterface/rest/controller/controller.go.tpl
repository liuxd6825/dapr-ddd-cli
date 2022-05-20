package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"{{.Namespace}}/pkg/query-service/application/internales/query_appservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}Controller struct {
	queryAppService *query_appservice.{{.Name}}AppQueryService
}

func New{{.Name}}Controller() *{{.Name}}Controller {
	return &{{.Name}}Controller{
		queryAppService: query_appservice.New{{.Name}}AppQueryService(),
	}
}

func (c *{{.Name}}Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/{id}", "GetById")
	b.Handle("GET", "/tenants/{tenantId}/users", "GetPagingData")
}

func (c *{{.Name}}Controller) GetById(ctx iris.Context, tenantId, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return m.queryAppService.FindById(ctx, tenantId, id)
	})
}

func (c *{{.Name}}Controller) GetPagingData(ctx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ctx, func(ctx context.Context) (interface{}, bool, error) {
		query, err := restapp.NewListQuery(ctx, tenantId)
		if err != nil {
			return nil, false, err
		}
		return m.queryAppService.FindPagingData(ctx, query)
	})
}
