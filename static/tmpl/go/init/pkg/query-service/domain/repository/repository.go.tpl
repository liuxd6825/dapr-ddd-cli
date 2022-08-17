package repository

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

{{if .IsAggregate}}
type Options interface {
    Timeout() *int
	UpdateMask() *[]string
}
{{- end }}

type {{.Name}}ViewRepository interface {
	Create(ctx context.Context, view *view.{{.Name}}View, opt Options) error
    CreateMany(ctx context.Context, views []*view.{{.Name}}View, opt Options) error

	Update(ctx context.Context, view *view.{{.Name}}View, opt Options) error
    UpdateMany(ctx context.Context, views []*view.{{.Name}}View, opt Options) error

	Delete(ctx context.Context, view *view.{{.Name}}View, opt Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View, opt Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opt Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opt Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opt Options) error
	DeleteAll(ctx context.Context, tenantId string, opt Options) error
    {{- if .IsEntity }}
	DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opt Options) error
	{{- end}}

	FindById(ctx context.Context, tenantId string, id string, opt Options) (*view.{{.Name}}View, bool, error)
	FindByIds(ctx context.Context, tenantId string, ids []string, opt Options) ([]*view.{{.Name}}View, bool, error)
	FindAll(ctx context.Context, tenantId string, opt Options) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt Options) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error)
	{{- if .IsEntity }}
	FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string, opt Options) ([]*view.{{.Name}}View, bool, error)
	{{- end}}
}
