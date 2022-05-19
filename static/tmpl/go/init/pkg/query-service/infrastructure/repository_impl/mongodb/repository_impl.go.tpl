import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	"{{.Namespace}}/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}Repository struct {
	base *BaseRepository[*projection.{{.Name}}View]
}

func New{{.Name}}Repository(opts ...*RepositoryOptions) repository.{{.Name}}ViewRepository {
	newFunc := func() *projection.{{.Name}}View {
		return &projection.{{.Name}}View{}
	}
	return &UserRepository{
		base: NewBaseRepository[*projection.{{.Name}}View](newFunc, "users", opts...),
	}
}

func (u *{{.Name}}Repository) CreateById(ctx context.Context, user *projection.UserView) (*projection.{{.Name}}View, error) {
	return u.base.CreateById(ctx, user)
}

func (u {{.Name}}Repository) UpdateById(ctx context.Context, user *projection.UserView) (*projection.{{.Name}}View, error) {
	return u.base.UpdateById(ctx, user)
}

func (u {{.Name}}Repository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.base.DeleteById(ctx, tenantId, id)
}

func (u {{.Name}}Repository) FindById(ctx context.Context, tenantId string, id string) (*projection.{{.Name}}View, bool, error) {
	return u.base.FindById(ctx, tenantId, id)
}

func (u {{.Name}}Repository) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*projection.{{.Name}}View] {
	return u.base.FindPaging(ctx, query)
}
