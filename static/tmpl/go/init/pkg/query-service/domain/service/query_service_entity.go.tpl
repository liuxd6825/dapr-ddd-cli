package {{.aggregate_name}}_queryservice

import (
	"context"
	view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	service_impl "{{.Namespace}}/pkg/query-service/infrastructure/domain/service/{{.aggregate_name}}_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryDomainService interface {
	Create(ctx context.Context, user *view.{{.Name}}View) error
	Update(ctx context.Context, user *view.{{.Name}}View) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error
	FindById(ctx context.Context, tenantId, id string) (*view.{{.Name}}View, bool, error)
	FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string) ([]*view.{{.Name}}View, bool, error)
	FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error)
}


func New{{.Name}}QueryDomainService() {{.Name}}QueryDomainService {
	return service_impl.New{{.Name}}QueryDomainService()
}
