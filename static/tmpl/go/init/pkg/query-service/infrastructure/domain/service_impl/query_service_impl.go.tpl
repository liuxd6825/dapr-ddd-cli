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

func (s *{{.Name}}QueryDomainServiceImpl) Create(ctx context.Context, view *view.{{.Name}}View) error {
	return s.repos.Create(ctx, view)
}

func (s *{{.Name}}QueryDomainServiceImpl) CreateMany(ctx context.Context, views []*view.{{.Name}}View) error {
	return s.repos.CreateMany(ctx, views)
}

func (s *{{.Name}}QueryDomainServiceImpl) Update(ctx context.Context, view *view.{{.Name}}View) error {
	return s.repos.Update(ctx, view)
}

func (s *{{.Name}}QueryDomainServiceImpl) UpdateMany(ctx context.Context, views []*view.{{.Name}}View) error {
	return s.repos.UpdateMany(ctx, views)
}

func (s *{{.Name}}QueryDomainServiceImpl) Delete(ctx context.Context, view *view.{{.Name}}View) error {
	return s.repos.Delete(ctx, view)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View) error {
	return s.repos.DeleteMany(ctx, tenantId, views)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string) error {
	return s.repos.DeleteById(ctx, tenantId, id)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string) error {
	return s.repos.DeleteByIds(ctx, tenantId, ids)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId, filter string) error {
	return s.repos.DeleteByFilter(ctx, tenantId, filter)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string) error {
	return s.repos.DeleteAll(ctx, tenantId)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindById(ctx context.Context, query *command.{{.Name}}FindByIdQuery) (*view.{{.Name}}View, bool, error) {
	return s.repos.FindById(ctx, query.TenantId, query.Id)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindByIds(ctx context.Context, query *command.{{.Name}}FindByIdsQuery) ([]*view.{{.Name}}View, bool, error) {
	return s.repos.FindByIds(ctx, query.TenantId, query.Ids)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindAll(ctx context.Context, query *command.{{.Name}}FindAllQuery) ([]*view.{{.Name}}View, bool, error) {
	return s.repos.FindAll(ctx, query.TenantId)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindPaging(ctx context.Context, query *command.{{.Name}}FindPagingQuery) (*command.{{.Name}}FindPagingResult, bool, error) {
	q := ddd_repository.NewFindPagingQuery()
	q.SetTenantId(query.TenantId)
	q.SetFilter(query.Filter)
	q.SetPageNum(query.PageNum)
	q.SetFields(query.Fields)
	q.SetSort(query.Sort)
	q.SetPageSize(query.PageSize)

	data, _, _ := s.repos.FindPaging(ctx, q)

	res := command.New{{.Name}}FindPagingResult()
	res.Data = data.GetData()
	res.TotalRows = data.GetTotalRows()
	res.TotalPages = data.GetTotalPages()
	res.PageNum = data.GetPageNum()
	res.PageSize = data.GetPageSize()
	res.Filter = data.GetFilter()
	res.Sort = data.GetSort()
	res.Error = data.GetError()
	res.IsFound = data.GetIsFound()

	return res, res.IsFound, res.Error
}


{{- if .IsEntity }}
func (s *{{.Name}}QueryDomainServiceImpl) FindByGraphId(ctx context.Context, query *command.{{.Name}}FindByGraphIdQuery) ([]*view.{{.Name}}View, bool, error) {
	return s.repos.FindByGraphId(ctx, query.TenantId, query.GraphId)
}
{{- end }}
