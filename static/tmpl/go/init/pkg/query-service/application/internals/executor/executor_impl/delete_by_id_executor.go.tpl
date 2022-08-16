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
// @Description: 新建分析图命令 命令执行器实现类
//
type {{.name}}DeleteByIdExecutor struct {
	{{- if .IsAggregate }}
	{{- range $entityName, $entity := .Aggregate.Entities}}
	{{$entity.FirstLowerName}}DomainService service.{{$entity.Name}}QueryDomainService
	{{- end }}
	{{- end }}
	{{.name}}Service service.{{.Name}}QueryDomainService
}

type {{.name}}DeleteByIdExecutorValidate struct {
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
func (e *{{.name}}DeleteByIdExecutor) Execute(ctx context.Context, tenantId string, id string) error {
	data := &{{.name}}DeleteByIdExecutorValidate{TenantId: tenantId, Id: id}
	if err := e.Validate(data); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
        {{- if .IsAggregate }}
        {{- range $entityName, $entity := .Aggregate.Entities}}
        if err:= e.{{$entity.FirstLowerName}}DomainService.DeleteBy{{$AggregateName}}Id(ctx, tenantId, id); err!=nil {
            return err
        }
        {{- end }}
        {{- end }}
		return e.{{.name}}Service.DeleteById(ctx, tenantId, id)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}DeleteByIdExecutor) Validate(v *{{.name}}DeleteByIdExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
    if len(v.TenantId) == 0 {
        return errors.New("Validate() error: tenantId is nil")
    }
    if len(v.Id) == 0 {
        return errors.New("Validate() error: id is nil")
    }
	return nil
}

//
// new{{.Name}}DeleteExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteExecutor
//
func new{{.Name}}DeleteByIdExecutor() *{{.name}}DeleteByIdExecutor {
	return &{{.name}}DeleteByIdExecutor{
		{{.name}}Service: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
