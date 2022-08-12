package service
{{- $commandPackage := .CommandPackage }}
{{- $modelPackage := .ModelPackage }}

import (
    "context"
    "sync"
    base_service "{{.Namespace}}/pkg/cmd-service/infrastructure/base/domain/service"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/model"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description:  {{.Description}} 命令领域服务
//
type {{.ClassName}} struct {
    base_service.BaseCommandDomainService
}

// 单例应用服务
var {{.name}}CommandDomainService *{{.Name}}CommandDomainService

// 并发安全
var once{{.Name}} sync.Once

//
// Get{{.Name}}CommandDomainService
// @Description: 获取单例领域服务
// @return service.{{.Name}}QueryDomainService
//
func Get{{.Name}}CommandDomainService() *{{.Name}}CommandDomainService {
    once{{.Name}}.Do(func() {
        {{.name}}CommandDomainService = new{{.Name}}CommandDomainService()
    })
	return {{.name}}CommandDomainService
}

//
// New{{.ClassName}}
// @Description: 创建领域服务
// @return *{{.ClassName}}
//
func new{{.ClassName}}() *{{.ClassName}} {
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
func (s *{{$ClassName}}) {{$command.ServiceFuncName}}(ctx context.Context, cmd *command.{{$commandName}}, opts ...ddd.DoCommandOption) (*model.{{$AggregateName}}Aggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

{{- end}}

//
//  doCommand
//  @Description:
//  @receiver s
//  @param ctx
//  @param cmd
//  @return *model.{{$AggregateName}}Aggregate
//  @return error
//
func (s *{{$ClassName}}) doCommand(ctx context.Context, cmd ddd.Command, validateFunc func() error, opts ...ddd.DoCommandOption) (*model.{{$AggregateName}}Aggregate, error) {
	option := ddd.NewDoCommandOptionMerges(opts...)

	// 进行业务检查
	if validateFunc != nil {
		if err := validateFunc(); err != nil {
			return nil, err
		}
	} else if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// 如果只是业务检查，则不执行领域命令，
	validOnly := option.GetIsValidOnly()
	if (validOnly == nil && cmd.GetIsValidOnly()) || (validOnly != nil && *validOnly == true) {
		return nil, nil
	}

	// 执行领域命令
	var err error
	agg := s.NewAggregate()
	if _, ok := cmd.(*command.{{$AggregateName}}CreateCommand); ok {
		err = ddd.CreateAggregate(ctx, agg, cmd)
	} else {
		err = ddd.CommandAggregate(ctx, agg, cmd)
	}

	// 如果领域命令执行时出错，则返回错误
	if err != nil {
		return nil, err
	}

	return agg, nil
}


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
func (s *{{.ClassName}}) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.{{.AggregateName}}Aggregate, bool, error) {
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
func (s *{{.ClassName}}) NewAggregate() *model.{{.AggregateName}}Aggregate {
	return model.New{{.AggregateName}}Aggregate()
}
