package appservice

import (
	"context"
	view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	domain_service "{{.Namespace}}/pkg/query-service/domain/service/{{.aggregate_name}}_service"
	"github.com/dapr/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type {{.Name}}QueryAppService struct {
	domainService domain_service.{{.Name}}QueryDomainService
}

func New{{.Name}}QueryAppService() *{{.Name}}QueryAppService {
	return &{{.Name}}QueryAppService{
		domainService: domain_service.New{{.Name}}QueryDomainService(),
	}
}

func (a *{{.Name}}QueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.{{.Name}}View, bool, error) {
	return a.domainService.FindById(ctx, tenantId, id)
}


func (a *{{.Name}}QueryAppService) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.{{.Name}}View], bool, error) {
	return a.domainService.FindPagingData(ctx, query)
}
