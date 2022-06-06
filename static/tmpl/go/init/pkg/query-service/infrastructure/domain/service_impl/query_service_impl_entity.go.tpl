package service_impl

import (
    "sync"
	"context"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
    "{{.Namespace}}/pkg/query-service/infrastructure/domain/{{.aggregate_name}}/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

//
// {{.Name}}QueryDomainServiceImpl
// @Description: 查询领域服务实现类
//
type {{.Name}}QueryDomainServiceImpl struct {
	repos repository.{{.Name}}ViewRepository
}

// 单例应用服务
var {{.name}}QueryDomainService service.{{.Name}}QueryDomainService

// 并发安全
var once{{.Name}} sync.Once

//
// Get{{.Name}}QueryDomainService
// @Description: 获取单例领域服务
// @return service.{{.Name}}QueryDomainService
//
func Get{{.Name}}QueryDomainService() service.{{.Name}}QueryDomainService {
    once{{.Name}}.Do(func() {
       {{.name}}QueryDomainService = new{{.Name}}QueryDomainService()
    })
	return {{.name}}QueryDomainService
}

func new{{.Name}}QueryDomainService() service.{{.Name}}QueryDomainService {
	return &{{.Name}}QueryDomainServiceImpl{
		repos: mongodb.New{{.Name}}ViewRepository(),
	}
}

func (u *{{.Name}}QueryDomainServiceImpl) FindById(ctx context.Context, tenantId, userId string) (*view.{{.Name}}View, bool, error) {
	return u.repos.FindById(ctx, tenantId, userId)
}

func (u *{{.Name}}QueryDomainServiceImpl) Create(ctx context.Context, view *view.{{.Name}}View) error {
	_, err := u.repos.Create(ctx, view)
	return err
}

func (u *{{.Name}}QueryDomainServiceImpl) Update(ctx context.Context, view *view.{{.Name}}View) error {
	_, err := u.repos.Update(ctx, view)
	return err
}

func (u *{{.Name}}QueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id)
}

func (u *{{.Name}}QueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string) error {
	return u.repos.DeleteAll(ctx, tenantId)
}

func (u *{{.Name}}QueryDomainServiceImpl) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error {
	return u.repos.DeleteBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string) (*[]*view.{{.Name}}View, bool, error){
    return u.repos.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error) {
	return u.repos.FindAll(ctx, tenantId)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return u.repos.FindPaging(ctx, query)
}

