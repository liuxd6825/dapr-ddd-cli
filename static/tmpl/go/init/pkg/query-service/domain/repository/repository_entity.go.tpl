package {{.aggregate_name}}_repository

import (
	"context"
	view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepository interface {
	Create(ctx context.Context, view *view.{{.Name}}View) (*view.{{.Name}}View, error)
	Update(ctx context.Context, view *view.{{.Name}}View) (*view.{{.Name}}View, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
	DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) error
	FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error)
	FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error)
}
