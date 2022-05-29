import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	repository "{{.Namespace}}/pkg/query-service/domain/repository/{{.aggregate_name}}_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepositoryImpl struct {
	base *BaseRepository[*projection.{{.Name}}View]
}

func New{{.Name}}ViewRepository(opts ...*RepositoryOptions) repository.{{.Name}}ViewRepository {
	newFunc := func() *projection.{{.Name}}View {
		return &projection.{{.Name}}View{}
	}
	return &{{.Name}}ViewRepositoryImpl{
		base: NewBaseRepository[*projection.{{.Name}}View](newFunc, "views", opts...),
	}
}

func (r *{{.Name}}ViewRepositoryImpl) CreateById(ctx context.Context, view *projection.{{.Name}}View) (*projection.{{.Name}}View, error) {
	return r.base.CreateById(ctx, view)
}

func (r {{.Name}}ViewRepositoryImpl) UpdateById(ctx context.Context, view *projection.{{.Name}}View) (*projection.{{.Name}}View, error) {
	return r.base.UpdateById(ctx, view)
}

func (r {{.Name}}ViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string) error {
	return r.base.DeleteById(ctx, tenantId, id)
}

func (r {{.Name}}ViewRepositoryImpl) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) (*projection.{{.Name}}View, bool, error) {
	return r.base.DeleteBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (r {{.Name}}ViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string) (*projection.{{.Name}}View, bool, error) {
	return r.base.FindById(ctx, tenantId, id)
}

func (r {{.Name}}ViewRepositoryImpl) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) (*[]projection.{{.Name}}View, bool, error) {
	return r.base.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (r {{.Name}}ViewRepositoryImpl) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*projection.{{.Name}}View] {
	return r.base.FindPaging(ctx, query)
}
