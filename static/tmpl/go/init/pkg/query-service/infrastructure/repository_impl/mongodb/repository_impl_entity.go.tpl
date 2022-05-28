import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	"{{.Namespace}}/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}ViewRepository struct {
	base *BaseRepository[*projection.{{.Name}}View]
}

func New{{.Name}}ViewRepository(opts ...*RepositoryOptions) repository.{{.Name}}ViewRepository {
	newFunc := func() *projection.{{.Name}}View {
		return &projection.{{.Name}}View{}
	}
	return &UserRepository{
		base: NewBaseRepository[*projection.{{.Name}}View](newFunc, "views", opts...),
	}
}

func (u *{{.Name}}ViewRepository) CreateById(ctx context.Context, view *projection.{{.Name}}View) (*projection.{{.Name}}View, error) {
	return u.base.CreateById(ctx, view)
}

func (u {{.Name}}ViewRepository) UpdateById(ctx context.Context, view *projection.{{.Name}}View) (*projection.{{.Name}}View, error) {
	return u.base.UpdateById(ctx, view)
}

func (u {{.Name}}ViewRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.base.DeleteById(ctx, tenantId, id)
}

func (u {{.Name}}ViewRepository) DeleteBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) (*projection.{{.Name}}View, bool, error) {
	return u.base.DeleteBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (u {{.Name}}ViewRepository) FindById(ctx context.Context, tenantId string, id string) (*projection.{{.Name}}View, bool, error) {
	return u.base.FindById(ctx, tenantId, id)
}

func (u {{.Name}}ViewRepository) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) (*[]projection.{{.Name}}View, bool, error) {
	return u.base.FindBy{{.AggregateName}}Id(ctx, tenantId, {{.aggregateName}}Id)
}

func (u {{.Name}}Repository) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*projection.{{.Name}}View] {
	return u.base.FindPaging(ctx, query)
}
