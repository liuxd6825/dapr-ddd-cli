package service
{{$AggregateName := .AggregateName}}
import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

//
// ScanCellQueryAppService
// @Description: {{.Description}}查询应用服务类
//
type {{.Name}}QueryAppService struct {
	domainService service.{{.Name}}QueryDomainService
}

//
// New{{.Name}}QueryAppService
// @Description: 创建{{.Name}}View查询应用服务
// @return *{{.Name}}QueryAppService
//
func New{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	return &{{.Name}}QueryAppService{
		domainService: service.New{{.Name}}QueryDomainService(),
	}
}

//
// Create
// @Description:  创建{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param v *view.{{.Name}}View
// @return error
//
func (a *{{.Name}}QueryAppService) Create(ctx context.Context, v *view.{{.Name}}View) error {
	return a.domainService.Create(ctx, v)
}

//
// Update
// @Description:  更新{{.Name}}View
// @receiver a
// @param ctx
// @param v 视图实体
// @return error
//
func (a *{{.Name}}QueryAppService) Update(ctx context.Context, v *view.{{.Name}}View) error {
	return a.domainService.Update(ctx, v)
}

//
// DeleteById
// @Description:  按ID删除{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
	return a.domainService.DeleteById(ctx, tenantId, id)
}

//
// DeleteAll
// @Description:  删除所有{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return a.domainService.DeleteAll(ctx, tenantId)
}

//
// DeleteBy{{.AggregateName}}Id
// @Description:  按{{.AggregateName}}Id删除{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error
//
func (a *{{.Name}}QueryAppService) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId, {{.aggregateName}}id string) error {
	return a.domainService.DeleteBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}id)
}


//
// FindById
// @Description:  按ID查询{{.Name}}View
// @receiver a
// @param ctx 上下文
// @param tenantId 租户ID
// @param id 视图ID
// @return *view.{{.Name}}View
// @return bool 是否查询到数据
// @return error
//
func (a *{{.Name}}QueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	return a.domainService.FindById(ctx, tenantId, id)
}

//
// FindAll
// @Description: 查询所有
// @receiver a
// @param ctx 上下文
// @param tenantId string 租户ID
// @return *[]*view.{{.Name}}View
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error) {
	return a.domainService.FindAll(ctx, tenantId)
}


//
// FindBy{{.AggregateName}}Id
// @Description: 根据{{.AggregateName}}Id查询
// @receiver a
// @param ctx 上下文
// @param tenantId string 租户ID
// @param {{.aggregateName}}Id  {{.AggregateName}}ID
// @return *ddd_repository.FindPagingResult[*view.{{.AggregateName}}View]
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) (*[]*view.{{.Name}}View, bool, error) {
	return a.domainService.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param query 分页查询条件
// @return *ddd_repository.FindPagingResult[*view.{{.AggregateName}}View]
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return a.domainService.FindPaging(ctx, query)
}
