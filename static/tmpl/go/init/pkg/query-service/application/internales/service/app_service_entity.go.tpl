package appservice

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
	domain_service "{{.Namespace}}/pkg/query-service/domain/service/{{.aggregate_name}}_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryAppService struct {
	domainService domain_service.{{.Name}}QueryDomainService
}

func New{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	return &{{.Name}}AppQueryService{
		domainService: domain_service.New{{.Name}}QueryService(),
	}
}

func (a *{{.Name}}QueryAppService) FindById(ctx context.Context, tenantId string, id string) (*projection.{{.Name}}View, bool, error) {
	return a.domainService.FindById(ctx, tenantId, id)
}

func (a *{{.Name}}QueryAppService) FindBy{{.AggregateName}}Id(ctx context.Context, tenantId string, {{.aggregateName}}Id string) ([]*projection.{{.Name}}View, bool, error) {
	return a.domainService.FindById(ctx, tenantId, id)
}

func (a *{{.Name}}QueryAppService) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.{{.Name}}View], bool, error) {
	return a.domainService.FindPagingData(ctx, query)
}
