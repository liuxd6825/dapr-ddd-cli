package appservice

import (
	"context"
	view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	domain_service "{{.Namespace}}/pkg/query-service/domain/service/{{.aggregate_name}}_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

//
// New{{.Name}}AppQueryService
// @Description: 创建{{.Name}}应用服务
// @return *{{.Name}}AppQueryService
//
type {{.Name}}QueryAppService struct {
	domainService domain_service.{{.Name}}QueryDomainService
}

//
// FindById
// @Description:  按ID查询{{.Name}}View
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @param id ID
// @return *view.{{.Name}}View
// @return bool 是否查询到数据
// @return error
//
func New{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	return &{{.Name}}QueryAppService{
		domainService: domain_service.New{{.Name}}QueryDomainService(),
	}
}


//
// FindById
// @Description:  按ID查询{{.Name}}View
// @receiver a
// @param ctx 上下文
// @param tenantId 租户ID
// @param id ID
// @return *view.{{.Name}}View
// @return bool 是否查询到数据
// @return error
//
func (a *{{.Name}}QueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	return a.domainService.FindById(ctx, tenantId, id)
}

//
// FindBy{{.AggregateName}}Id
// @Description: 根据{{.AggregateName}}Id查询
// @receiver a
// @param ctx 上下文
// @param tenantId string 租户ID
// @param {{.aggregateName}}Id  {{.AggregateName}}ID
// @return *ddd_repository.FindPagingResult[*view.ScanBillView]
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) ([]*view.{{.Name}}View, bool, error) {
	return a.domainService.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

//
// FindPagingData
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param query 分页查询条件
// @return *ddd_repository.FindPagingResult[*view.ScanBillView]
// @return bool 是否查询到数据
// @return error 错误
//
func (a *{{.Name}}QueryAppService) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return a.domainService.FindPagingData(ctx, query)
}
