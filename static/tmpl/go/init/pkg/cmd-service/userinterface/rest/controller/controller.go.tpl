package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"{{.Namespace}}/pkg/cmd-service/application/internals/cmd_appservice"
	"{{.Namespace}}/pkg/cmd-service/application/internals/query_appservice"
	"{{.Namespace}}/pkg/cmd-service/domain/command/{{.AggregateCommandPackage}}"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)
{{- $AggregateName := .AggregateName}}
{{- $ClassName := .ClassName}}
{{- $CommandPackage := .CommandPackage}}
{{- $AppService := .AppService}}
type {{$AggregateName}}Controller struct {
    {{.AppService}} *cmdappservice.{{.AggregateName}}CommandAppService
    queryAppId     string
}

func New{{$AggregateName}}Controller() *UserController {
    return &{{$AggregateName}}Controller{
        {{.AppService}}: cmdappservice.New{{.AggregateName}}CommandAppService(),
        queryAppId:     queryappservice.Get{{.AggregateName}}QueryAppService().AppId(),
    }
}

{{$resource := .Aggregate.LowerName}}
func (c *{{$ClassName}}) BeforeActivation(b mvc.BeforeActivation) {
    b.Handle("GET", "/tenants/{tenantId}/{{$resource}}s/aggregate/{id}", "GetAggregateById")
    b.Handle("POST", "/tenants/{tenantId}/{{$resource}}s", "{{.AggregateName}}Create")
    b.Handle("POST", "/tenants/{tenantId}/{{$resource}}s:get", "{{.AggregateName}}CreateAndGet")
    b.Handle("PATCH", "/tenants/{tenantId}/{{$resource}}s", "{{.AggregateName}}Update")
    b.Handle("PATCH", "/tenants/{tenantId}/{{$resource}}s:get", "{{.AggregateName}}UpdateAndGet")
}

func (c *{{$ClassName}}) GetAggregateById(ctx iris.Context, tenantId string, id string) {
    _, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
        return c.{{.AppService}}.GetAggregateById(ctx, tenantId, id)
	})
}

{{- range $cmdName, $cmd := .Commands}}
func (c *{{$ClassName}}) {{$cmd.ServiceFuncName}}(ctx iris.Context) {
	cmd := &{{$CommandPackage}}.{{$cmdName}}{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.{{$AppService}}.{{$cmd.ServiceFuncName}}(ctx, cmd)
	})
}

func (c *{{$ClassName}}) {{$cmd.ServiceFuncName}}AndGet(ctx iris.Context) {
	cmd := &{{$CommandPackage}}.{{$cmdName}}{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.{{$AppService}}.{{$cmd.ServiceFuncName}}(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getUserById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}
{{- end }}

func (c *{{.ClassName}}) getUserById(ctx context.Context, tenantId, userId string) (data interface{}, isFound bool, err error) {
	return queryappservice.Get{{.AggregateName}}QueryAppService().GetById(ctx, tenantId, userId)
}
