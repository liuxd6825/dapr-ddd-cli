package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	services "{{.Namespace}}/pkg/cmd-service/application/internals/service/{{.aggregate_name}}_service"
	commands "{{.Namespace}}/pkg/cmd-service/domain/command/{{.aggregate_name}}_command"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)
{{- $AggregateName := .AggregateName}}
{{- $ClassName := .ClassName}}
type {{$AggregateName}}Controller struct {
    appService *services.{{.AggregateName}}CommandAppService
    queryAppId     string
}

func New{{$AggregateName}}Controller() *{{$AggregateName}}Controller {
    return &{{$AggregateName}}Controller{
        appService: services.New{{.AggregateName}}CommandAppService(),
        queryAppId:     services.Get{{.AggregateName}}QueryAppService().AppId(),
    }
}

{{$resource := .Aggregate.LowerName}}
func (c *{{$ClassName}}) BeforeActivation(b mvc.BeforeActivation) {
    b.Handle("GET", "/tenants/{tenantId}/{{$resource}}/aggregate/{id}", "GetAggregateById")
    {{- range $cmdName, $cmd := .Commands }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$cmd.HttpPath}}", "{{$cmd.Name}}")
    {{- end}}
}

func (c *{{$ClassName}}) GetAggregateById(ctx iris.Context, tenantId string, id string) {
    _, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
        return c.appService.GetAggregateById(ctx, tenantId, id)
	})
}

{{- range $cmdName, $cmd := .Commands}}

func (c *{{$ClassName}}) {{$cmd.Name}}(ctx iris.Context) {
	cmd := &commands.{{$cmdName}}{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.appService.{{$cmd.ServiceFuncName}}(ctx, cmd)
	})
}
{{- if $cmd.IsAggregateCreateOrUpdate}}
func (c *{{$ClassName}}) {{$cmd.Name}}AndGet(ctx iris.Context) {
	cmd := &commands.{{$cmdName}}{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.appService.{{$cmd.ServiceFuncName}}(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getUserById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}
{{- end }}
{{- end }}

func (c *{{.ClassName}}) getUserById(ctx context.Context, tenantId, userId string) (data interface{}, isFound bool, err error) {
	return services.Get{{.AggregateName}}QueryAppService().GetById(ctx, tenantId, userId)
}
