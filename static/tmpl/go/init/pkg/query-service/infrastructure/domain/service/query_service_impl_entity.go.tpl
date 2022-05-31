package {{.aggregate_name}}_service

import (
	"context"
    view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	domain_service "{{.Namespace}}/pkg/query-service/domain/service/{{.aggregate_name}}_service"
	domain_repository "{{.Namespace}}/pkg/query-service/domain/repository/{{.aggregate_name}}_repository"
	repository_impl "{{.Namespace}}/pkg/query-service/infrastructure/domain/repository/mongodb/{{.aggregate_name}}_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryDomainServiceImpl struct {
	repos domain_repository.{{.Name}}ViewRepository
}

func New{{.Name}}QueryDomainService() domain_service.{{.Name}}QueryDomainService {
	return &{{.Name}}QueryDomainServiceImpl{
		repos: repository_impl.New{{.Name}}ViewRepository(),
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

func (u *{{.Name}}QueryDomainServiceImpl) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error {
	return u.repos.DeleteBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string) ([]*view.{{.Name}}View, bool, error){
    return u.repos.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return u.repos.FindPaging(ctx, query)
}

