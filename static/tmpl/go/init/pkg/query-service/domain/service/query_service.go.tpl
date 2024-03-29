package service

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
	{{- if .IsAggregate}}
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
    {{- end }}
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
)

{{if .IsAggregate}}
type Options interface {
	repository.Options
}
{{- end }}

type {{.Name}}QueryDomainService interface {
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
    {{- if .IsEntity}}
    DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opts ...Options) error
    {{- end}}

	FindById(ctx context.Context, qry *query.{{.Name}}FindByIdQuery, opts ...Options) (*view.{{.Name}}View, bool, error)
	FindByIds(ctx context.Context, qry *query.{{.Name}}FindByIdsQuery, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	FindAll(ctx context.Context, qry *query.{{.Name}}FindAllQuery, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, qry *query.{{.Name}}FindPagingQuery, opts ...Options) (*query.{{.Name}}FindPagingResult, bool, error)
    {{- if .IsEntity}}
    FindBy{{.AggregateName}}Id(ctx context.Context, qry *query.{{.Name}}FindBy{{.AggregateName}}IdQuery, opts ...Options) ([]*view.{{.Name}}View, bool, error)
    {{- end}}
}
