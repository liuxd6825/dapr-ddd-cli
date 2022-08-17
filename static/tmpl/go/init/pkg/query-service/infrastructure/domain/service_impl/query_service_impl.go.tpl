package service_impl

import (
    "sync"
	"context"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
	{{- if .IsMongo}}
    repos_impl "{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/repository_impl/mongodb"
    {{- else if .IsNeo4j}}
    repos_impl "{{.Namespace}}/pkg/query-service/infrastructure/domain_impl/{{.aggregate_name}}/repository_impl/neo4j"
    {{- end }}
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
		repos: repos_impl.New{{.Name}}ViewRepository(),
	}
}

func (s *{{.Name}}QueryDomainServiceImpl) Create(ctx context.Context, view *view.{{.Name}}View, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.Create(ctx, view, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) CreateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.CreateMany(ctx, views, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) Update(ctx context.Context, view *view.{{.Name}}View, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.Update(ctx, view, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) UpdateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.UpdateMany(ctx, views, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) Delete(ctx context.Context, view *view.{{.Name}}View, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.Delete(ctx, view, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.DeleteMany(ctx, tenantId, views, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.DeleteById(ctx, tenantId, id, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.DeleteByIds(ctx, tenantId, ids, opt)
}

{{- if .IsEntity }}
func (s *{{.Name}}QueryDomainServiceImpl) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opts ...service.Options) error {
    opt := MergeOptions(opts...)
    return s.repos.DeleteBy{{.AggregateName}}Id (ctx, tenantId, {{.aggregateName}}Id, opt)
}
{{- end }}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.DeleteByFilter(ctx, tenantId, filter, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
    opt := MergeOptions(opts...)
	return s.repos.DeleteAll(ctx, tenantId, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindById(ctx context.Context, qry *query.{{.Name}}FindByIdQuery, opts ...service.Options) (*view.{{.Name}}View, bool, error) {
    opt := MergeOptions(opts...)
	return s.repos.FindById(ctx, qry.TenantId, qry.Id, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindByIds(ctx context.Context, qry *query.{{.Name}}FindByIdsQuery, opts ...service.Options) ([]*view.{{.Name}}View, bool, error) {
    opt := MergeOptions(opts...)
	return s.repos.FindByIds(ctx, qry.TenantId, qry.Ids, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindAll(ctx context.Context, qry *query.{{.Name}}FindAllQuery, opts ...service.Options) ([]*view.{{.Name}}View, bool, error) {
    opt := MergeOptions(opts...)
	return s.repos.FindAll(ctx, qry.TenantId, opt)
}

func (s *{{.Name}}QueryDomainServiceImpl) FindPaging(ctx context.Context, qry *query.{{.Name}}FindPagingQuery, opts ...service.Options) (*query.{{.Name}}FindPagingResult, bool, error) {
	q := ddd_repository.NewFindPagingQuery()
	q.SetTenantId(qry.TenantId)
	q.SetFilter(qry.Filter)
	q.SetPageNum(qry.PageNum)
	q.SetFields(qry.Fields)
	q.SetSort(qry.Sort)
	q.SetPageSize(qry.PageSize)

    opt := MergeOptions(opts...)
	data, _, _ := s.repos.FindPaging(ctx, q, opt)

	res := query.New{{.Name}}FindPagingResult()
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
func (s *{{.Name}}QueryDomainServiceImpl) FindBy{{.AggregateName}}Id(ctx context.Context, qry *query.{{.Name}}FindBy{{.AggregateName}}IdQuery, opts ...service.Options) ([]*view.{{.Name}}View, bool, error) {
    opt := MergeOptions(opts...)
	return s.repos.FindBy{{.AggregateName}}Id(ctx, qry.TenantId, qry.{{.AggregateName}}Id, opt)
}
{{- end }}
