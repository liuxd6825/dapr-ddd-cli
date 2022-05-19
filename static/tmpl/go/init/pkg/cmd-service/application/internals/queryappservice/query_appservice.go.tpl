package queryappservice

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/projection"
)

type {{.AggregateName}}QueryAppService interface {
	QueryAppService[*projection.{{.Aggregate.Name}}View]
	GetById(ctx context.Context, tenantId, id string) (data *projection.{{.AggregateName}}View, isFound bool, err error)
}

type {{.aggregateName}}QueryAppService struct {
	BaseQueryAppService[*projection.{{.AggregateName}}View]
}

var _{{.aggregateName}}QueryAppService {{.AggregateName}}QueryAppService

func init() {
	_{{.aggregateName}}QueryAppService = new{{.AggregateName}}QueryAppService()
}

func Get{{.AggregateName}}QueryAppService() {{.AggregateName}}QueryAppService {
	return _{{.aggregateName}}QueryAppService
}

func new{{.AggregateName}}QueryAppService() {{.AggregateName}}QueryAppService {
	res := &{{.aggregateName}}QueryAppService{}
	res.appId = "query-service"
	res.resourceName = "{{.aggregateName}}s"
	res.apiVersion = "v1.0"
	return res
}

func (s *{{.aggregateName}}QueryAppService) GetById(ctx context.Context, tenantId, id string) (data *projection.{{.AggregateName}}View, isFound bool, err error) {
	view := &projection.{{.AggregateName}}View{}
	isFound, err = s.GetData(ctx, tenantId, id, data)
	if isFound {
	    data = view
	}
	return
}
