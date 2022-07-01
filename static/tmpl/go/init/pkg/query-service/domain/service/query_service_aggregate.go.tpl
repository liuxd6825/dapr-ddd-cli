package service

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryDomainService interface {
	Create(ctx context.Context, v *view.{{.Name}}View) error
	Update(ctx context.Context, v *view.{{.Name}}View) error

    CreateMany(ctx context.Context, views *[]*view.{{.Name}}View) error
    UpdateManyById(ctx context.Context, views *[]*view.{{.Name}}View) error
    UpdateManyByFilter(ctx context.Context, tenantId, filter string, data interface{}) error
    DeleteByFilter(ctx context.Context, tenantId, filter string) error

	DeleteById(ctx context.Context, tenantId string, id string) error
	DeleteAll(ctx context.Context, tenantId string) error

    FindById(ctx context.Context, tenantId, id string) (*view.{{.Name}}View, bool, error)
    FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error)
}

