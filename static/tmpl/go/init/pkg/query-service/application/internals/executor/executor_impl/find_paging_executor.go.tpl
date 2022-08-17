package {{.snake_name}}_impl
{{ $AggregateName := .AggregateName }}
import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/assembler"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/service_impl"
)

//
// {{.name}}CreateCommandCommandExecutor
// @Description: 分页查询
//
type {{.name}}FindPagingExecutor struct {
	domainService service.{{.Name}}QueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}FindPagingExecutor) Execute(ctx context.Context, aq *appquery.{{.Name}}FindPagingAppQuery) (*appquery.{{.Name}}FindPagingResult, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}
	qry := query.New{{.Name}}FindPagingQuery(aq.TenantId, aq.Fields, aq.Filter, aq.Sort, aq.PageNum, aq.PageSize)
	fpr, ok, err := e.domainService.FindPaging(ctx, qry)
	if err != nil {
		return nil, false, err
	}
	res := assembler.Ass{{.Name}}FindPagingResult(fpr)
	return res, ok, nil
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}FindPagingExecutor) Validate(aq *appquery.{{.Name}}FindPagingAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// new{{.Name}}{{.Name}}FindPagingExecutor
// @Description: 新建命令执行器
// @return *{{.name}}FindPagingExecutor
//
func new{{.Name}}FindPagingExecutor() *{{.name}}FindPagingExecutor {
	return &{{.name}}FindPagingExecutor{
		domainService: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
