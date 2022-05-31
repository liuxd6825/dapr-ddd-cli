package {{.Package}}
{{- $commandPackage := .CommandPackage }}
{{- $modelPackage := .ModelPackage }}

import (
    "context"
    "{{.Namespace}}/pkg/cmd-service/domain/service"
    "{{.Namespace}}/pkg/cmd-service/domain/command/{{$commandPackage}}"
    "{{.Namespace}}/pkg/cmd-service/domain/model/{{$modelPackage}}"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type {{.ClassName}} struct {
    service.BaseCommandDomainService
}

//
// New{{.ClassName}}
// @Description: 创建领域服务
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{}
}

{{- $ClassName := .ClassName}}
{{- $AggregateName:= .AggregateName }}
{{- range $commandName, $command := .Commands}}

//
// {{$command.ServiceFuncName}}
// @Description: {{$command.Description}}
// @receiver s
// @param ctx 上下文
// @param cmd {{$command.Description}}
// @return *model.{{$ClassName}}
// @return error
//
func (s *{{$ClassName}}) {{$command.ServiceFuncName}}(ctx context.Context, cmd *{{$commandPackage}}.{{$commandName}}) (*{{$modelPackage}}.{{$AggregateName}}Aggregate, error) {
    if err := s.ValidateCommand(cmd); err != nil {
        return nil, err
    }
    agg := s.NewAggregate()
    {{- if $command.IsCreate }}
    err := ddd.CreateAggregate(ctx, agg, cmd)
    {{- else }}
    err := ddd.CommandAggregate(ctx, agg, cmd)
    {{- end }}
    if err != nil {
        return nil, err
    }
    return agg, nil
}

{{- end}}

//
// GetAggregateById
// @Description: 获取聚合对象
// @receiver s
// @param ctx 上下文
// @param tenantId 租户id
// @param id 主键id
// @return *{{.ModelPackage}}.{{.ClassName}}  聚合对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
//
func (s *{{.ClassName}}) GetAggregateById(ctx context.Context, tenantId string, id string) (*{{$modelPackage}}.{{.AggregateName}}Aggregate, bool, error) {
    agg := s.NewAggregate()
    _, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
    return agg, ok, err
}

//
// NewAggregate
// @Description: 新建聚合对象
// @receiver s
// @return *{{.ModelPackage}}.{{.ClassName}} 聚合对象
//
func (s *{{.ClassName}}) NewAggregate() *{{$modelPackage}}.{{.AggregateName}}Aggregate {
	return {{$modelPackage}}.New{{.AggregateName}}Aggregate()
}
