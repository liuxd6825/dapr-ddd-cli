package repository

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

{{if .IsAggregate}}
type Options interface {
	GetTimeout() *time.Duration
	SetTimeout(v *time.Duration) Options
	GetUpdateFields() *[]string
	SetUpdateFields(*[]string) Options
	Merge(opts ...Options) Options
}
{{- end }}

type {{.Name}}ViewRepository interface {
	Create(ctx context.Context, view *view.{{.Name}}View, opts ...Options) error
    CreateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...Options) error

	Update(ctx context.Context, view *view.{{.Name}}View, opts ...Options) error
    UpdateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...Options) error

	Delete(ctx context.Context, view *view.{{.Name}}View, opts ...Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View, opts ...Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...Options) error
    {{- if .IsEntity }}
	DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opts ...Options) error
	{{- end}}

	FindById(ctx context.Context, tenantId string, id string, opts ...Options) (*view.{{.Name}}View, bool, error)
	FindByIds(ctx context.Context, tenantId string, ids []string, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	FindAll(ctx context.Context, tenantId string, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...Options) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error)
	{{- if .IsEntity }}
	FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string,{{.aggregateName}}Id string, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	{{- end}}
}
