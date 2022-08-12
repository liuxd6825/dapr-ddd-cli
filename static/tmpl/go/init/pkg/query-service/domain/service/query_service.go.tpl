package service

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/query"
    "{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
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

	FindById(ctx context.Context, qry *query.FindByIdQuery, opts ...Options) (*view.{{.Name}}View, bool, error)
	FindByIds(ctx context.Context, qry *query.FindByIdsQuery, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	FindAll(ctx context.Context, qry *query.FindAllQuery, opts ...Options) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, qry *query.FindPagingQuery, opts ...Options) (*query.FindPagingResult, bool, error)
    {{- if .IsEntity}}
    FindBy{{.AggregateName}}Id(ctx context.Context, qry *query.FindBy{{.AggregateName}}IdQuery, opts ...Options) ([]*view.{{.Name}}View, bool, error)
    {{- end}}
}
