package neo4j

import (
	"context"
 	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
	"{{.Namespace}}/pkg/query-service/infrastructure/db/dao/neo4j_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepositoryImpl struct {
	dao *neo4j_dao.Dao[*view.{{.Name}}View]
}

func New{{.Name}}ViewRepository() repository.{{.Name}}ViewRepository {
	return &{{.Name}}ViewRepositoryImpl{
		dao: neo4j_dao.NewDao[*view.{{.Name}}View]([]string{"{{.aggregate_name}}"}),
	}
}

func (r *{{.Name}}ViewRepositoryImpl) Create(ctx context.Context, view *view.{{.Name}}View, opt repository.Options) error {
	return r.dao.Insert(ctx, view)
}

func (r *{{.Name}}ViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.{{.Name}}View, opt repository.Options) error {
	return r.dao.InsertMany(ctx, views)
}

func (r *{{.Name}}ViewRepositoryImpl) Update(ctx context.Context, view *view.{{.Name}}View, opt repository.Options) error {
	return r.dao.Update(ctx, view)
}

func (r *{{.Name}}ViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.{{.Name}}View, opt repository.Options) error {
	return r.dao.UpdateMany(ctx, views)
}

func (r *{{.Name}}ViewRepositoryImpl) Delete(ctx context.Context, view *view.{{.Name}}View, opt repository.Options) error {
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId())
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View, opt repository.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opt)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) error {
	return r.dao.DeleteByIds(ctx, tenantId, ids)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opt repository.Options) error {
	return r.dao.DeleteByFilter(ctx, tenantId, filter)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opt repository.Options) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opt repository.Options) error {
	return r.dao.DeleteAll(ctx, tenantId)
}

func (r *{{.Name}}ViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opt repository.Options) (*view.{{.Name}}View, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r *{{.Name}}ViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) ([]*view.{{.Name}}View, bool, error) {
	return r.dao.FindByIds(ctx, tenantId, ids)
}

func (r *{{.Name}}ViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opt repository.Options) ([]*view.{{.Name}}View, bool, error) {
	return r.dao.FindAll(ctx, tenantId).Result()
}

func (r *{{.Name}}ViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt repository.Options) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return r.dao.FindPaging(ctx, query).Result()
}

{{- if .IsEntity }}
func (r {{.Name}}ViewRepositoryImpl) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opt repository.Options) ([]*view.{{.Name}}View, bool, error) {
	filterMap := map[string]interface{}{
	    "{{.aggregateName}}Id": {{.aggregateName}}Id,
	}
	return r.dao.FindListByMap(ctx, tenantId, filterMap).Result()
}
{{- end }}