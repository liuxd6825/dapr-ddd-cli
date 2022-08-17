package {{.snake_name}}_impl
{{ $AggregateName := .AggregateName }}
import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/infrastructure/db/session"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/service_impl"
)

//
// {{.name}}DeleteAllExecutor
// @Description: 按{{.AggregateName}}Id删除
//
type {{.name}}DeleteBy{{.AggregateName}}IdExecutor struct {
	{{.name}}Service service.{{.Name}}QueryDomainService
}

type {{.name}}DeleteBy{{.AggregateName}}IdExecutorValidate struct {
	TenantId string
	{{.AggregateName}}Id       string
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}DeleteBy{{.AggregateName}}IdExecutor) Execute(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error {
	data := &{{.name}}DeleteBy{{.AggregateName}}IdExecutorValidate{TenantId: tenantId, {{.AggregateName}}Id: {{.aggregateName}}Id}
	if err := e.Validate(data); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.{{.name}}Service.DeleteBy{{$AggregateName}}Id(ctx, data.TenantId, data.{{.AggregateName}}Id)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}DeleteBy{{.AggregateName}}IdExecutor) Validate(v *{{.name}}DeleteBy{{.AggregateName}}IdExecutorValidate) error {
	if v == nil {
		return errors.New("Validate(v) error: v is nil")
	}
    if len(v.TenantId) == 0 {
        return errors.New("Validate(v) error: v.TenantId is nil")
    }
    if len(v.{{.AggregateName}}Id) == 0 {
        return errors.New("Validate(v) error: v.{{.AggregateName}}Id is nil")
    }
	return nil
}

//
// new{{.Name}}DeleteBy{{.AggregateName}}IdExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteBy{{.AggregateName}}IdExecutor
//
func new{{.Name}}DeleteBy{{.AggregateName}}IdExecutor() *{{.name}}DeleteBy{{.AggregateName}}IdExecutor {
	return &{{.name}}DeleteBy{{.AggregateName}}IdExecutor{
		{{.name}}Service: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
