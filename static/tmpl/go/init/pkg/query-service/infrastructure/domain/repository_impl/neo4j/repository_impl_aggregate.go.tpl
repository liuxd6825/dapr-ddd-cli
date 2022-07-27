package neo4j

import (
	"context"
 	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
	"{{.Namespace}}/pkg/query-service/infrastructure/base/domain/repository/mongodb_base"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepositoryImpl struct {
	dao *Dao[*view.{{.Name}}View]
}

func New{{.Name}}ViewRepository(opts ...*mongodb_base.RepositoryOptions) repository.{{.Name}}ViewRepository {
	return &{{.Name}}ViewRepositoryImpl{
		dao: newDao[*view.{{.Name}}View]("{{.aggregate_name}}", opts...),
	}
}

func (r *{{.Name}}ViewRepositoryImpl) Insert(ctx context.Context, view *view.{{.Name}}View) (*view.{{.Name}}View, error) {
	return r.dao.Insert(ctx, view)
}

func (r *{{.Name}}ViewRepositoryImpl) InsertMany(ctx context.Context, views *[]*view.{{.Name}}View) error {
	return r.dao.CreateMany(ctx, views)
}

func (r {{.Name}}ViewRepositoryImpl) Update(ctx context.Context, view *view.{{.Name}}View) (*view.{{.Name}}View, error) {
	return r.dao.Update(ctx, view)
}

func (r *{{.Name}}ViewRepositoryImpl) UpdateManyById(ctx context.Context, views *[]*view.{{.Name}}View) error {
	return r.dao.UpdateManyById(ctx, views)
}

func (r *{{.Name}}ViewRepositoryImpl) UpdateManyByFilter(ctx context.Context, tenantId, filter string, data interface{}) error {
	return r.dao.UpdateManyByFilter(ctx, tenantId, filter, data)
}

func (r {{.Name}}ViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r {{.Name}}ViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string) error {
	return r.dao.DeleteAll(ctx, tenantId)
}

func (r {{.Name}}ViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r {{.Name}}ViewRepositoryImpl) FindAll(ctx context.Context, tenantId string) (*[]*view.{{.Name}}View, bool, error) {
	return r.dao.FindAll(ctx, tenantId).Result()
}

func (r {{.Name}}ViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return r.dao.FindPaging(ctx, query).Result()
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string) error {
	return r.dao.DeleteById(ctx, tenantId, filter)
}
