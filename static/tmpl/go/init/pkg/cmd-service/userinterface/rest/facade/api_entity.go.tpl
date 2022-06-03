package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
    service "{{.Namespace}}/pkg/cmd-service/application/internals/service/{{.aggregate_name}}_service"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)
{{- $AggregateName := .AggregateName}}
{{- $ClassName := .ClassName}}
type {{$ClassName}} struct {
    appService *service.{{.AggregateName}}CommandAppService
    queryAppId string
}

func New{{$ClassName}}() *{{$ClassName}} {
    return &{{$ClassName}}{
        appService: service.New{{.AggregateName}}CommandAppService(),
        queryAppId: service.Get{{.AggregateName}}QueryAppService().AppId(),
    }
}

{{$AggregatePluralName := .AggregatePluralName}}
{{$EntityPluralName := .EntityPluralName}}
func (c *{{$ClassName}}) BeforeActivation(b mvc.BeforeActivation) {
    {{- range $cmdName, $cmd := .Commands }}
        {{- if $cmd.IsEntityDeleteByIdCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}/{{$EntityPluralName}}", "{{$cmd.ControllerMethod}}")
        {{- else if $cmd.IsEntityCreateCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}/{{$EntityPluralName}}", "{{$cmd.ControllerMethod}}")
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}/{{$EntityPluralName}}:get", "{{$cmd.ControllerMethod}}AndGet")
        {{- else if $cmd.IsEntityUpdateCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}/{{$EntityPluralName}}", "{{$cmd.ControllerMethod}}")
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}/{{$EntityPluralName}}:get", "{{$cmd.ControllerMethod}}AndGet")
        {{- else if $cmd.IsEntityCustomCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}/{{$EntityPluralName}}:{{$cmd.HttpMethod}}", "{{$cmd.ControllerMethod}}")
        {{- end }}
    {{- end }}
}


{{- range $i, $cmd := .Commands}}
{{- if $cmd.IsEntity }}
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}(ctx iris.Context) {
	cmd := &command.{{$cmd.Name}}{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.appService.{{$cmd.ServiceFuncName}}(ctx, cmd)
	})
}

{{- if $cmd.IsEntityCreateOrUpdateCommand }}
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}AndGet(ctx iris.Context) {
	cmd := &command.{{$cmd.Name}}{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.appService.{{$cmd.ServiceFuncName}}(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}
{{- end }}
{{- end }}
{{- end }}

func (c *{{.ClassName}}) getById(ctx context.Context, tenantId, id string) (data *view.{{.Name}}View, isFound bool, err error) {
	return service.Get{{.Name}}QueryAppService().GetById(ctx, tenantId, id)
}
