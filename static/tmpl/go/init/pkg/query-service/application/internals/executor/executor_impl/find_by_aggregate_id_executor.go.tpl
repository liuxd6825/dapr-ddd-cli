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
// {{.name}}FindBy{{.AggregateName}}IdExecutor
// @Description: 按{{.AggregateName}}Id查询
//
type {{.name}}FindBy{{.AggregateName}}IdExecutor struct {
	domainService service.{{.Name}}QueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}FindBy{{.AggregateName}}IdExecutor) Execute(ctx context.Context, aq *appquery.{{.Name}}FindBy{{.AggregateName}}IdAppQuery) ([]*view.{{.Name}}View, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindBy{{.AggregateName}}Id(ctx, query.New{{.Name}}FindBy{{.AggregateName}}IdQuery(aq.TenantId, aq.{{.AggregateName}}Id))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}FindBy{{.AggregateName}}IdExecutor) Validate(aq *appquery.{{.Name}}FindBy{{.AggregateName}}IdAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
    if len(aq.{{.AggregateName}}Id) == 0 {
        return errors.ErrorOf("Validate(aq) error: aq.{{.AggregateName}}Id is nil")
    }
	return nil
}

//
// new{{.Name}}FindBy{{.AggregateName}}IdExecutor()
// @Description: 新建命令执行器
// @return *{{.name}}FindBy{{.AggregateName}}IdExecutor
//
func new{{.Name}}FindBy{{.AggregateName}}IdExecutor() *{{.name}}FindBy{{.AggregateName}}IdExecutor {
	return &{{.name}}FindBy{{.AggregateName}}IdExecutor{
		domainService: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
