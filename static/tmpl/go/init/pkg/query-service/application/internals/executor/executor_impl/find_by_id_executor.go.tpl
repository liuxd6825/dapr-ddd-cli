package {{.snake_name}}_impl
{{ $AggregateName := .AggregateName }}
import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/service_impl"
)

//
// {{.name}}FindByIdExecutor
// @Description: 按Id查询
//
type {{.name}}FindByIdExecutor struct {
	domainService service.{{.Name}}QueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}FindByIdExecutor) Execute(ctx context.Context, aq *appquery.{{.Name}}FindByIdAppQuery) (*view.{{.Name}}View, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindById(ctx, query.New{{.Name}}FindByIdQuery(aq.TenantId, aq.Id))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}FindByIdExecutor) Validate(aq *appquery.{{.Name}}FindByIdAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// New{{.Name}}CreateCommandExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteExecutor
//
func new{{.Name}}FindByIdExecutor() *{{.name}}FindByIdExecutor {
	return &{{.name}}FindByIdExecutor{
		domainService: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
