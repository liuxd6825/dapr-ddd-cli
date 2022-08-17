package {{.snake_name}}_impl
{{ $AggregateName := .AggregateName }}
import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/infrastructure/db/session"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// {{.name}}DeleteAllExecutor
// @Description: 删除所有
//
type {{.name}}DeleteAllExecutor struct {
	{{- if .IsAggregate }}
	{{- range $entityName, $entity := .Aggregate.Entities}}
	{{$entity.FirstLowerName}}DomainService service.{{$entity.Name}}QueryDomainService
	{{- end }}
	{{- end }}
	{{.name}}Service service.{{.Name}}QueryDomainService
}

type {{.name}}DeleteAllExecutorValidate struct {
	TenantId string
	Id       string
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}DeleteAllExecutor) Execute(ctx context.Context, tenantId string) error {
	data := &{{.name}}DeleteAllExecutorValidate{TenantId: tenantId}
	if err := e.Validate(data); err != nil {
		return err
	}

	return session.StartSession(ctx, func(ctx context.Context) error {
        {{- if .IsAggregate }}
        {{- range $entityName, $entity := .Aggregate.Entities}}
        if err:= e.{{$entity.FirstLowerName}}DomainService.DeleteAll(ctx, tenantId); err!=nil {
            return err
        }
        {{- end }}
        {{- end }}
		return e.{{.name}}Service.DeleteAll(ctx, tenantId)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}DeleteAllExecutor) Validate(v *{{.name}}DeleteAllExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// new{{.Name}}DeleteExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteExecutor
//
func new{{.Name}}DeleteAllExecutor() *{{.name}}DeleteAllExecutor {
	return &{{.name}}DeleteAllExecutor{
		{{.name}}Service: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
