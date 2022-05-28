package queryservice_impl

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	"{{.Namespace}}/pkg/query-service/domain/repository"
	"{{.Namespace}}/pkg/query-service/infrastructure/repository_impl/mongodb/{{.Name}}_repository_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryService struct {
	repos repository.{{.Name}}ViewRepository
}

func New{{.Name}}QueryService() *{{.Name}}QueryService {
	return &{{.Name}}QueryService{
		repos: mongodb.New{{.Name}}Repository(),
	}
}

func (u *{{.Name}}QueryService) FindById(ctx context.Context, tenantId, userId string) (*projection.{{.Name}}View, bool, error) {
	return u.repos.FindById(ctx, tenantId, userId)
}

func (u *{{.Name}}QueryService) Create(ctx context.Context, view *projection.{{.Name}}View) error {
	_, err := u.repos.CreateById(ctx, view)
	return err
}

func (u *{{.Name}}QueryService) Update(ctx context.Context, view *projection.{{.Name}}View) error {
	_, err := u.repos.UpdateById(ctx, view)
	return err
}

func (u *{{.Name}}QueryService) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id)
}

func (u *{{.Name}}QueryService) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.{{.Name}}View], bool, error) {
	return u.repos.FindPaging(ctx, query).Result()
}
