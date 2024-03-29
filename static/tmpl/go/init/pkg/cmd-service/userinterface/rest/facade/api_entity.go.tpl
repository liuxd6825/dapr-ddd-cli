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
//
// BeforeActivation
// @Description: 注册http
// @receiver c
// @param ctx
//
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
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}(ictx iris.Context) {
	_ = restapp.DoCmd(ictx, func(ctx context.Context) error {
	    cmd, err := {{$aggregateName}}Assembler.Ass{{$cmd.AppName}}(ictx)
        if err != nil {
			return err
        }
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
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}AndGet(ictx iris.Context) {
	_ = restapp.Do(ictx, func() error {
        appCmd, err := {{$aggregateName}}Assembler.Ass{{$cmd.AppName}}(ictx)
        if err != nil {
			return err
        }

        _, _, err = restapp.DoCmdAndQueryOne(ictx, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
            return c.service.{{$cmd.ServiceFuncName}}(ctx, appCmd)
        }, func(ctx context.Context) (interface{}, bool, error) {
            {{- if $cmd.DataIsItems }}
            return c.service.QueryByIds(ctx, appCmd.GetTenantId(), appCmd.Data.GetIds())
            {{- else }}
            return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
            {{- end }}
        })

        return err
	})
}

{{- end }}
{{- end }}
{{- end }}

