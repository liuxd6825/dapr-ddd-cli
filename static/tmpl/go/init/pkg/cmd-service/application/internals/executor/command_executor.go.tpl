package executor

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/assembler"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/dto"
	domain_service "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

type {{.Name}}Executor interface {
	Execute(context.Context, *dto.{{.Name}}Dto) error
}

type {{.name}}Executor struct {
	domainService *domain_service.{{.AggregateName}}CommandDomainService
}

func (e *{{.name}}Executor) Execute(ctx context.Context, cmdDto *dto.{{.Name}}Dto) error {
	cmd, err := assembler.Ass{{.Name}}(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = e.domainService.{{.Command.ServiceFuncName}}(ctx, cmd)
	return err
}


//
// New{{.Name}}Executor
// @Description: 新建命令执行器
// @return service.{{.Name}}Executor
//
func new{{.Name}}Executor() *{{.name}}Executor {
	return &{{.name}}Executor{
		domainService: domain_service.Get{{.AggregateName}}CommandDomainService(),
	}
}

