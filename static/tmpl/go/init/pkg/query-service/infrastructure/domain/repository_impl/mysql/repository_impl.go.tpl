package mongodb

import (
	"context"
	"fmt"
 	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/repository"
    "{{.Namespace}}/pkg/query-service/infrastructure/db/dao/mysql_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepositoryImpl struct {
	dao *mysql_dao.Dao[*view.{{.Name}}View]
}

var new{{.Name}}Func = func() *view.{{.Name}}View {
	return &view.{{.Name}}View{}
}

var newList{{.Name}}Func = func() []*view.{{.Name}}View {
	return []*view.{{.Name}}View{}
}

func New{{.Name}}ViewRepository() repository.{{.Name}}ViewRepository {
	return &{{.Name}}ViewRepositoryImpl{
		dao: mysql_dao.NewDao[*view.{{.Name}}View](nil, new{{.Name}}Func, newList{{.Name}}Func),
	}
}

func (r *{{.Name}}ViewRepositoryImpl) Create(ctx context.Context, view *view.{{.Name}}View, opts ...repository.Options) error {
    ops := newOptions(opts...)
	return r.dao.Insert(ctx, view, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...repository.Options) error {
    ops := newOptions(opts...)
	return r.dao.InsertMany(ctx, views, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) Update(ctx context.Context, view *view.{{.Name}}View, opts ...repository.Options) error {
    ops := newOptions(opts...)
	return r.dao.Update(ctx, view, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.{{.Name}}View, opts ...repository.Options) error {
    ops := newOptions(opts...)
	return r.dao.UpdateMany(ctx, views, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) Delete(ctx context.Context, view *view.{{.Name}}View, opts ...repository.Options) error {
    ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId(), ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.{{.Name}}View, opts ...repository.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opts...)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...repository.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, tenantId, id, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...repository.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByIds(ctx, tenantId, ids, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opts ...repository.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteAll(ctx, tenantId, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...repository.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}

{{- if .IsEntity }}
func (r *{{.Name}}ViewRepositoryImpl) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opts ...repository.Options) error {
    ops := newOptions(opts...)
	filter := fmt.Sprintf(`{{.aggregateName}}Id == "%s"`, {{.aggregateName}}Id)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}
{{- end}}

func (r *{{.Name}}ViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opts ...repository.Options) (*view.{{.Name}}View, bool, error) {
    ops := newOptions(opts...)
	return r.dao.FindById(ctx, tenantId, id, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opts ...repository.Options) ([]*view.{{.Name}}View, bool, error) {
    ops := newOptions(opts...)
	return r.dao.FindByIds(ctx, tenantId, ids, ops...)
}

func (r *{{.Name}}ViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opts ...repository.Options) ([]*view.{{.Name}}View, bool, error) {
    ops := newOptions(opts...)
	return r.dao.FindAll(ctx, tenantId, ops...).Result()
}

func (r *{{.Name}}ViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...repository.Options) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
    ops := newOptions(opts...)
	return r.dao.FindPaging(ctx, query, ops...).Result()
}

{{- if .IsEntity }}
func (r *{{.Name}}ViewRepositoryImpl) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string, opts ...repository.Options) ([]*view.{{.Name}}View, bool, error) {
	filterMap := map[string]interface{}{
	    "{{.aggregateName}}Id": {{.aggregateName}}Id,
	}
    ops := newOptions(opts...)
	return r.dao.FindListByMap(ctx, tenantId, filterMap, ops...).Result()
}
{{- else }}

func newOptions(opts ...repository.Options) []ddd_repository.Options {
	var options []ddd_repository.Options
	for _,o := range options {
		options = append(options, o)
	}
	return options
}
{{- end }}

