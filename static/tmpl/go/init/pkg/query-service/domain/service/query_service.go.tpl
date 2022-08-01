package service

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/command"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryDomainService interface {
	Create(ctx context.Context, view *view.{{.Name}}View) error
	CreateMany(ctx context.Context, views []*view.{{.Name}}View) error

	Update(ctx context.Context, view *view.{{.Name}}View) error
	UpdateMany(ctx context.Context, views []*view.{{.Name}}View) error

	Delete(ctx context.Context, view *view.{{.Name}}View) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string) error
	DeleteByFilter(ctx context.Context, tenantId, filter string) error
	DeleteAll(ctx context.Context, tenantId string) error

	FindById(ctx context.Context, query *command.{{.Name}}FindByIdQuery) (*view.{{.Name}}View, bool, error)
	FindByIds(ctx context.Context, query *command.{{.Name}}FindByIdsQuery) ([]*view.{{.Name}}View, bool, error)
	FindAll(ctx context.Context, query *command.{{.Name}}FindAllQuery) ([]*view.{{.Name}}View, bool, error)
	FindPaging(ctx context.Context, query *command.{{.Name}}FindPagingQuery) (*command.{{.Name}}FindPagingResult, bool, error)
    {{- if .IsEntity}}
    FindBy{{.AggregateName}}Id(ctx context.Context, query *command.{{.Name}}FindBy{{.AggregateName}}IdQuery) ([]*view.{{.Name}}View, bool, error)
    {{- end}}
}
