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
// {{.name}}CreateCommandCommandExecutor
// @Description: 新建分析图命令 命令执行器实现类
//
type {{.name}}FindByIdsExecutor struct {
	domainService service.{{.Name}}QueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}FindByIdsExecutor) Execute(ctx context.Context, aq *appquery.{{.Name}}FindByIdsAppQuery) ([]*view.{{.Name}}View, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindByIds(ctx, query.New{{.Name}}FindByIdsQuery(aq.TenantId, aq.Ids))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}FindByIdsExecutor) Validate(aq *appquery.{{.Name}}FindByIdsAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	if len(aq.Ids) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.Ids is nil")
	}
	return nil
}

//
// New{{.Name}}CreateCommandExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteExecutor
//
func new{{.Name}}FindByIdsExecutor() *{{.name}}FindByIdsExecutor {
	return &{{.name}}FindByIdsExecutor{
		domainService: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
