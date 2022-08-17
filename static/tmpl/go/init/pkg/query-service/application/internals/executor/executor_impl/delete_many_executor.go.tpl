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
// {{.Name}}DeleteManyExecutor
// @Description: 删除多个
//
type {{.name}}DeleteManyExecutor struct {
	{{- if .IsAggregate }}
	{{- range $entityName, $entity := .Aggregate.Entities}}
	{{$entity.FirstLowerName}}DomainService service.{{$entity.Name}}QueryDomainService
	{{- end }}
	{{- end }}
	{{.name}}Service  service.{{.Name}}QueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *{{.name}}DeleteManyExecutor) Execute(ctx context.Context, tenantId string, vList []*view.{{.Name}}View) error {
	if err := e.Validate(vList); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
        {{- if .IsAggregate }}
        {{- if not .Aggregate.Entities.Empty }}
        for _, item := range vList {
            {{- range $entityName, $entity := .Aggregate.Entities}}
            if err:= e.{{$entity.FirstLowerName}}DomainService.DeleteBy{{$AggregateName}}Id(ctx, tenantId, item.Id); err!=nil {
                return err
            }
            {{- end }}
        }
        {{- end }}
        {{- end }}
		return e.{{.name}}Service.DeleteMany(ctx, tenantId, vList)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *{{.name}}DeleteManyExecutor) Validate(view []*view.{{.Name}}View) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// new{{.Name}}DeleteManyExecutor
// @Description: 新建命令执行器
// @return service.{{.Name}}DeleteExecutor
//
func new{{.Name}}DeleteManyExecutor() *{{.name}}DeleteManyExecutor {
	return &{{.name}}DeleteManyExecutor{
		{{.name}}Service: service_impl.Get{{.Name}}QueryDomainService(),
	}
}
