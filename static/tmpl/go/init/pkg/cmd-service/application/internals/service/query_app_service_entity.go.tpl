package service

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	base "{{.Namespace}}/pkg/cmd-service/infrastructure/application/service"
)

type {{.Name}}QueryAppService interface {
	base QueryAppService[*view.{{.Name}}View]
	GetById(ctx context.Context, tenantId, id string) (data *view.{{.Name}}View, isFound bool, err error)
}

type {{.name}}QueryAppService struct {
	base.BaseQueryAppService[*view.{{.AggregateName}}View]
}

var _{{.name}}QueryAppService {{.Name}}QueryAppService

func init() {
	_{{.name}}QueryAppService = new{{.Name}}QueryAppService()
}

func Get{{.Name}}QueryAppService() {{.Name}}QueryAppService {
	return _{{.name}}QueryAppService
}

func new{{.Name}}QueryAppService() {{.Name}}QueryAppService {
	res := &{{.aggregateName}}QueryAppService{}
    res.SetAppId("{{.QueryServiceName}}")
    res.SetResourceName("{{.Aggregate.PluralName}}/{{.Entity.PluralName}}")
    res.SetApiVersion("{{.ApiVersion}}")
	return res
}


//
// GetById
// @Description: 按id获取{{.Description}}投影类
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param id {{.Description}} Id
// @return data {{.Description}} 信息
// @return isFound 是否找到
// @return err 错误信息
//
func (s *{{.name}}QueryAppService) GetById(ctx context.Context, tenantId, id string) (data *view.{{.Name}}View, isFound bool, err error) {
	v := &view.{{.Name}}View{}
	isFound, err = s.GetData(ctx, tenantId, id, data)
	if isFound {
	    data = v
	}
	return
}
