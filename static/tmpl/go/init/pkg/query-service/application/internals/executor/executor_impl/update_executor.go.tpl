package {{.snake_name}}_impl
{{ $AggregateName := .AggregateName }}
import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/infrastructure/db/session"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/service_impl"
)

//
// {{.name}}UpdateExecutor
// @Description: 更新多个
//
type {{.name}}UpdateExecutor struct {
	{{.name}}Service service.{{.Name}}QueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}UpdateExecutor) Execute(ctx context.Context, view *view.{{.Name}}View) error {
	if err := e.Validate(view); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.{{.name}}Service.Update(ctx, view)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}UpdateExecutor) Validate(view *view.{{.Name}}View) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// new{{.Name}}CreateExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteExecutor
//
func new{{.Name}}UpdateExecutor() *{{.name}}UpdateExecutor {
	return &{{.name}}UpdateExecutor{
		{{.name}}Service: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
