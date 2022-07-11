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

func (s *{{.Name}}QueryDomainServiceImpl) FindById(ctx context.Context, tenantId, userId string) (*view.{{.Name}}View, bool, error) {
	return s.repos.FindById(ctx, tenantId, userId)
}


func (s *{{.Name}}QueryDomainServiceImpl) CreateMany(ctx context.Context, views *[]*view.{{.Name}}View) error {
	return s.repos.CreateMany(ctx, views)
}

func (s *{{.Name}}QueryDomainServiceImpl) UpdateManyById(ctx context.Context, views *[]*view.{{.Name}}View) error {
	return s.repos.UpdateManyById(ctx, views)
}

func (s *{{.Name}}QueryDomainServiceImpl) UpdateManyByFilter(ctx context.Context, tenantId, filter string, data interface{}) error {
	return s.repos.UpdateManyByFilter(ctx, tenantId, filter, data)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId string, filter string) error {
	return s.repos.DeleteByFilter(ctx, tenantId, filter)
}

func (s *{{.Name}}QueryDomainServiceImpl) Create(ctx context.Context, view *view.{{.Name}}View) error {
	_, err := s.repos.Create(ctx, view)
	return err
}

func (s *{{.Name}}QueryDomainServiceImpl) Update(ctx context.Context, view *view.{{.Name}}View) error {
	_, err := s.repos.Update(ctx, view)
	return err
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string) error {
	return s.repos.DeleteById(ctx, tenantId, id)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string) error {
	return s.repos.DeleteAll(ctx, tenantId)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error {
	return s.repos.DeleteBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string) (*[]*view.{{.Name}}View, bool, error){
    return s.repos.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error) {
	return s.repos.FindAll(ctx, tenantId)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return s.repos.FindPaging(ctx, query)
}

