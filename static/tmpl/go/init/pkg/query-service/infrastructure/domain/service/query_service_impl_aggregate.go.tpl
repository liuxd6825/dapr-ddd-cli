package service_impl

import (
	"context"
    view "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain/{{.aggregate_name}}/repository/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryDomainServiceImpl struct {
	repos repository.{{.Name}}ViewRepository
}

func New{{.Name}}QueryDomainService() service.{{.Name}}QueryDomainService {
	return &{{.Name}}QueryDomainServiceImpl{
		repos: mongodb.New{{.Name}}ViewRepository(),
	}
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

func (u *{{.Name}}QueryDomainServiceImpl) FindById(ctx context.Context, tenantId, id string) (*view.{{.Name}}View, bool, error) {
	return u.repos.FindById(ctx, tenantId, id)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error) {
	return u.repos.FindAll(ctx, tenantId)
}

func (u *{{.Name}}QueryDomainServiceImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return u.repos.FindPaging(ctx, query)
}
