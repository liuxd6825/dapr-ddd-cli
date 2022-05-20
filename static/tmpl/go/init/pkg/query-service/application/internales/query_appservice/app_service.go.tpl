package appservice

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	"{{.Namespace}}/pkg/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}AppQueryService struct {
	service queryservice.{{.Name}}QueryService
}

func New{{.Name}}AppQueryService() *{{.Name}}AppQueryService {
	return &{{.Name}}AppQueryService{
		service: queryservice.New{{.Name}}QueryService(),
	}
}

func (a *{{.Name}}AppQueryService) FindById(ctx context.Context, tenantId string, userId string) (*projection.{{.Name}}View, bool, error) {
	return a.service.FindById(ctx, tenantId, userId)
}

func (a *{{.Name}}AppQueryService) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.{{.Name}}View], bool, error) {
	return a.service.FindPagingData(ctx, query)
}
