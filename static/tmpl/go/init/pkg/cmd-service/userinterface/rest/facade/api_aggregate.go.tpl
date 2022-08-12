{{- $AggregateName := .AggregateName}}
{{- $ClassName := .ClassName}}
{{- $aggregateName := .aggregateName}}
{{- $AggregatePluralName := .AggregatePluralName}}
package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"{{.Namespace}}/pkg/cmd-service/userinterface/rest/{{.aggregate_name}}/assembler"
    "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
)

var {{.aggregateName}}Assembler = assembler.{{$AggregateName}}Assembler{}

type {{$ClassName}} struct {
    service *service.{{$AggregateName}}CommandAppService
}

func New{{$ClassName}}() *{{$ClassName}} {
    return &{{$ClassName}}{
        service: service.New{{$AggregateName}}CommandAppService(),
    }
}

func (c *{{$ClassName}}) BeforeActivation(b mvc.BeforeActivation) {
    b.Handle("GET", "/tenants/{tenantId}/{{$AggregatePluralName}}/aggregate/{id}", "FindAggregateById")
    {{- range $cmdName, $cmd := .Commands }}
        {{- if $cmd.IsAggregateDeleteByIdCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}", "{{$cmd.ControllerMethod}}")
        {{- else if $cmd.IsAggregateCreateCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}", "{{$cmd.ControllerMethod}}")
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}:get", "{{$cmd.ControllerMethod}}AndGet")
        {{- else if $cmd.IsAggregateUpdateCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}", "{{$cmd.ControllerMethod}}")
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}:get", "{{$cmd.ControllerMethod}}AndGet")
        {{- else if $cmd.IsAggregateCustomCommand }}
    b.Handle("{{$cmd.HttpType}}", "/tenants/{tenantId}/{{$AggregatePluralName}}:{{$cmd.HttpMethod}}", "{{$cmd.ControllerMethod}}")
        {{- end }}
    {{- end }}
}

//
// FindAggregateById godoc
// @Summary      按聚合根ID查找聚合对象
// @Description  按聚合根ID查找聚合对象
// @Tags         {{.AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/{{.AggregatePluralName}}/aggregate/{id} [get]
//
func (c *{{$ClassName}}) FindAggregateById(ictx iris.Context, tenantId string, id string) {
    _, _, _ = restapp.DoQueryOne(ictx, func(ctx context.Context) (interface{}, bool, error) {
        return c.service.FindAggregateById(ctx, tenantId, id)
	})
}


{{- range $i, $cmd := .Commands}}
{{- if $cmd.IsAggregate }}

//
// {{$cmd.ControllerMethod}} godoc
// @Summary      {{$cmd.Description}}
// @Description  {{$cmd.Description}}
// @Tags         {{$AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/{{$AggregatePluralName}} [{{$cmd.HttpType}}]
//
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}(ictx iris.Context) {
	_ = restapp.DoCmd(ictx, func(ctx context.Context) error {
	    cmd, err := {{$aggregateName}}Assembler.Ass{{$cmd.Name}}Dto(ictx)
        if err != nil {
            return err
        }
		return c.service.{{$cmd.ServiceFuncName}}(ctx, cmd)
	})
}

{{- if $cmd.IsAggregateCreateOrUpdate}}

//
// {{$cmd.ControllerMethod}}AndGet godoc
// @Summary      {{$cmd.Description}}
// @Description  {{$cmd.Description}}
// @Tags         {{$AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/{{$AggregatePluralName}}:get [{{$cmd.HttpType}}]
//
func (c *{{$ClassName}}) {{$cmd.ControllerMethod}}AndGet(ictx iris.Context) {
	_ = restapp.Do(ictx, func() error {
		cmd, err := {{$aggregateName}}Assembler.Ass{{$cmd.Name}}Dto(ictx)
        if err != nil {
            return err
    	}

        _, _, err = restapp.DoCmdAndQueryOne(ictx, c.service.QueryAppId, cmd, func(ctx context.Context) error {
            return c.service.{{$cmd.ServiceFuncName}}(ctx, cmd)
        }, func(ctx context.Context) (interface{}, bool, error) {
            return c.service.QueryById(ctx, cmd.GetTenantId(), cmd.Data.Id)
        })
        return err
    })
}
{{- end }}
{{- end }}
{{- end }}
