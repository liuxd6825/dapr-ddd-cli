package executor

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/model"
	domain_service "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/service"
)

type FindAggregateByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string ) (*model.{{.Aggregate.Name}}Aggregate, bool, error)
}

type findAggregateByIdExecutor struct {
	domainService *domain_service.{{.AggregateName}}CommandDomainService
}

//
// Execute
// @Description:
// @receiver s
// @param ctx 上下文
// @param tenantId 租户Id
// @param id 聚合根Id
// @return error
//
func (s *findAggregateByIdExecutor) Execute(ctx context.Context, tenantId string, id string) (*model.{{.Aggregate.Name}}Aggregate, bool, error) {
	return s.domainService.GetAggregateById(ctx, tenantId, id)
}

//
// new{{.Name}}Executor
// @Description: 新建命令执行器
// @return service.{{.Name}}Executor
//
func new{{.Name}}Executor() *findAggregateByIdExecutor {
	return &{{.name}}Executor{
		domainService: domain_service.Get{{.AggregateName}}CommandDomainService(),
	}
}

