package {{.aggregate_name}}_repository

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepository interface {
	Create(ctx context.Context, view *projection.{{.Name}}View) (*projection.{{.Name}}View, error)
	UpdateById(ctx context.Context, view *projection.{{.Name}}View) (*projection.{{.Name}}View, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindById(ctx context.Context, tenantId string, id string) (*projection.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query *ddd_repository.PagingQuery) (res *ddd_repository.PagingData[*projection.{{.Name}}View], isFound bool, err error)
}
