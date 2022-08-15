package executor

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/assembler"
	"{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/appcmd"
	domain_service "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/service"
    "github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// {{.Name}}Executor
// @Description: {{.Description}} 命令执行器接口
//
type {{.Name}}Executor interface {
	Execute(context.Context, *appcmd.{{.AppName}}) error
}

//
// {{.name}}CommandExecutor
// @Description: {{.Description}} 命令执行器实现类
//
type {{.name}}Executor struct {
	domainService *domain_service.{{.AggregateName}}CommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}Executor) Execute(ctx context.Context, appCmd *appcmd.{{.AppName}}) error {
	if err := e.Validate(appCmd); err!=nil {
		return err
	}

	cmd, err := assembler.Ass{{.Name}}(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.{{.Command.ServiceFuncName}}(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}Executor) Validate(appCmd *appcmd.{{.AppName}}) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
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

