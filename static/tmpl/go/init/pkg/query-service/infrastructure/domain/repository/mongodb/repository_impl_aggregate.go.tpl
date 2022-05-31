package {{.aggregate_name}}_repository

import (
	"context"
	"{{.Namespace}}/pkg/query-service/infrastructure/domain/repository/mongodb"
 	view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	repository "{{.Namespace}}/pkg/query-service/domain/repository/{{.aggregate_name}}_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepositoryImpl struct {
	base *mongodb.BaseRepository[*view.{{.Name}}View]
}

func New{{.Name}}ViewRepository(opts ...*mongodb.RepositoryOptions) repository.{{.Name}}ViewRepository {
	newFunc := func() *view.{{.Name}}View {
		return &view.{{.Name}}View{}
	}
	return &{{.Name}}ViewRepositoryImpl{
		base: mongodb.NewBaseRepository[*view.{{.Name}}View](newFunc, "{{.aggregate_name}}", opts...),
	}
}

func (r *{{.Name}}ViewRepositoryImpl) Create(ctx context.Context, view *view.{{.Name}}View) (*view.{{.Name}}View, error) {
	return r.base.Create(ctx, view)
}

func (r {{.Name}}ViewRepositoryImpl) Update(ctx context.Context, view *view.{{.Name}}View) (*view.{{.Name}}View, error) {
	return r.base.Update(ctx, view)
}

func (r {{.Name}}ViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string) error {
	return r.base.DeleteById(ctx, tenantId, id)
}

func (r {{.Name}}ViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	return r.base.FindById(ctx, tenantId, id)
}

func (r {{.Name}}ViewRepositoryImpl) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return r.base.FindPaging(ctx, query).Result()
}
