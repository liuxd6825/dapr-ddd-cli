package service
{{$AggregateName := .AggregateName}}
import (
    "sync"
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
    "{{.Namespace}}/pkg/query-service/infrastructure/domain/{{.aggregate_name}}/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

//
// {{.Name}}QueryAppService
// @Description: {{.Description}}查询应用服务类
//
type {{.Name}}QueryAppService struct {
	{{.name}}DomainService service.{{.Name}}QueryDomainService
    {{- range $entityName, $entity := .Aggregate.Entities}}
    {{$entity.FirstLowerName}}DomainService service.{{$entity.Name}}QueryDomainService
    {{- end }}
}

// 单例应用服务
var {{.name}}QueryAppService *{{.Name}}QueryAppService

// 并发安全
var once{{.Name}} sync.Once

//
// Get{{.Name}}QueryAppService
// @Description: 获取单例应用服务
// @return *{{.Name}}QueryAppService
//
func Get{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
    once{{.Name}}.Do(func() {
        {{.name}}QueryAppService = new{{.Name}}QueryAppService()
    })
	return {{.name}}QueryAppService
}


//
// New{{.Name}}QueryAppService
// @Description: 创建{{.Name}}查询应用服务
// @return *{{.Name}}QueryAppService
//
func new{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	return &{{.Name}}QueryAppService{
		{{.name}}DomainService: service_impl.Get{{.Name}}QueryDomainService(),
        {{- range $entityName, $entity := .Aggregate.Entities}}
        {{$entity.FirstLowerName}}DomainService : service_impl.Get{{$entity.Name}}QueryDomainService(),
        {{- end }}
	}
}

//
// Create
// @Description:  创建{{.Name}}View
// @receiver a
// @param ctx
// @return *view.{{.Name}}View
// @return error
//
func (a *{{.Name}}QueryAppService) Create(ctx context.Context, v *view.{{.Name}}View) error {
	return a.{{.name}}DomainService.Create(ctx, v)
}


//
// Update
// @Description:  按ID更新{{.Name}}View
// @receiver a
// @param ctx
// @param v  *view.{{.Name}}View
// @return error
//
func (a *{{.Name}}QueryAppService) Update(ctx context.Context, v *view.{{.Name}}View) error {
	return a.{{.name}}DomainService.Update(ctx, v)
}


//
// DeleteById
// @Description:  按ID删除{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @param v *view.{{.Name}}View
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
    {{- range $entityName, $entity := .Aggregate.Entities}}
    if err:= a.{{$entity.FirstLowerName}}DomainService.DeleteBy{{$AggregateName}}Id(ctx, tenantId, id); err!=nil {
        return err
    }
    {{- end }}
	return a.{{.name}}DomainService.DeleteById(ctx, tenantId, id)
}



//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @param v *view.{{.Name}}View
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteAll(ctx context.Context, tenantId, id string) error {
    {{- range $entityName, $entity := .Aggregate.Entities}}
    if err:= a.{{$entity.FirstLowerName}}DomainService.DeleteAll(ctx, tenantId); err!=nil {
        return err
    }
    {{- end }}
	return a.{{.name}}DomainService.DeleteAll(ctx, tenantId)
}


//
// FindById
// @Description:  按ID查询{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return *view.{{.Name}}View
// @return bool 是否查询到数据
// @return error
//
func (a *{{.Name}}QueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	return a.{{.name}}DomainService.FindById(ctx, tenantId, id)
}


//
// FindAll
// @Description: 查询所有view.{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return *[]*view.{{.Name}}View
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error) {
	return a.{{.name}}DomainService.FindAll(ctx, tenantId)
}


//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param query 分页查询条件
// @return *ddd_repository.FindPagingResult[*view.ScanBillView]
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return a.{{.name}}DomainService.FindPaging(ctx, query)
}
