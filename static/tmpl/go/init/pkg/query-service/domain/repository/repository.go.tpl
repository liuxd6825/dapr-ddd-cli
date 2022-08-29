package repository

import (
	"context"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)



type {{.Name}}ViewRepository interface {
	Create(ctx context.Context, view *view.{{.Name}}View, opts ...service.Options) error
    CreateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...service.Options) error

	Update(ctx context.Context, view *view.{{.Name}}View, opts ...service.Options) error
    UpdateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...service.Options) error

	Delete(ctx context.Context, view *view.{{.Name}}View, opts ...service.Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View, opts ...service.Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error
    {{- if .IsEntity }}
	DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opts ...service.Options) error
	{{- end}}

	FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.{{.Name}}View, bool, error)
	FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.{{.Name}}View, bool, error)
	FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error)
	{{- if .IsEntity }}
	FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string, opts ...service.Options) ([]*view.{{.Name}}View, bool, error)
	{{- end}}
}
