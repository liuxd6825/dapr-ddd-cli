{{- $AggregatePluralName := .AggregatePluralName}}
{{- $aggregateName := .aggregateName}}
{{- $AggregateName := .AggregateName}}
{{- $ClassName := .ClassName}}
package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
    "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
)


type {{$ClassName}} struct {
    service *service.{{.AggregateName}}CommandAppService
}

func New{{$ClassName}}() *{{$ClassName}} {
    return &{{$ClassName}}{
        service: service.New{{.AggregateName}}CommandAppService(),
    }
}

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

//
// {{$cmd.ControllerMethod}}
// @Description: {{$cmd.Description}}
// @receiver c
// @param ctx
//
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}(ctx iris.Context) {
    cmd, err := {{$aggregateName}}Assembler.Ass{{$cmd.Name}}Dto(ctx)
    if err != nil {
        restapp.SetError(ctx, err)
        return
    }
	_ = restapp.DoCmd(ctx, func(ctx context.Context) error {
		return c.service.{{$cmd.ServiceFuncName}}(ctx, cmd)
	})
}

{{- if $cmd.IsEntityCreateOrUpdateCommand }}

//
// {{$cmd.ControllerMethod}}AndGet
// @Description: {{$cmd.Description}}
// @receiver c
// @param ctx
//
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}AndGet(ctx iris.Context) {
    cmd, err := {{$aggregateName}}Assembler.Ass{{$cmd.Name}}Dto(ctx)
    if err != nil {
        restapp.SetError(ctx, err)
        return
    }
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.service.QueryAppId, cmd, func(ctx context.Context) error {
		return c.service.{{$cmd.ServiceFuncName}}(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.service.QueryById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

{{- end }}
{{- end }}
{{- end }}

