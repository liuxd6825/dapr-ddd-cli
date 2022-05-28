package {{.aggregate_name}}_queryservice

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	"{{.Namespace}}/pkg/query-service/infrastructure/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryService interface {
	FindById(ctx context.Context, tenantId, id string) (*projection.{{.Name}}View, bool, error)
	Create(ctx context.Context, user *projection.{{.Name}}View) error
	Update(ctx context.Context, user *projection.{{.Name}}View) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.{{.Name}}View], bool, error)
}

func New{{.Name}}QueryService() {{.Name}}QueryService {
	return queryservice_impl.mongodb.New{{.Name}}QueryService()
}
