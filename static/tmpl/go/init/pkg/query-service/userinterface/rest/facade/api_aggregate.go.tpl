{{$AggregatePluralName := .AggregatePluralName}}
package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/assembler"

)

var {{.Name}}Assembler = assembler.{{.Name}}

type {{.Name}}QueryApi struct {
	queryService *service.{{.Name}}QueryAppService
}

func New{{.Name}}QueryApi() *{{.Name}}QueryApi {
	return &{{.Name}}QueryApi{
		queryService: service.Get{{.Name}}QueryAppService(),
	}
}

func (a *{{.Name}}QueryApi) BeforeActivation(b mvc.BeforeActivation) {
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{$AggregatePluralName}}/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{$AggregatePluralName}}:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{$AggregatePluralName}}", "FindPaging")
}

// FindById godoc
// @Summary      按ID查询
// @Description  get string by ID
// @Tags         {{$AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "User ID"
// @Success      200        {object}   dto.{{.Name}}FindByIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/{{$AggregatePluralName}}/{id} [get]
func (a *{{.Name}}QueryApi) FindById(ictx iris.Context {
	_, _, _ = restapp.DoQueryOne(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := {{.Name}}Assembler.AssFindByIdRequest(ictx)
    	if err != nil {
    		return nil, false, err
    	}
		v, b, e := a.queryService.FindById(ctx, req.TenantId, req.Id)
		return {{.Name}}Assembler.AssFindByIdResponse(ictx, v, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有用户
// @Description  get string by ID
// @Tags         {{$AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.UserFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/{{$AggregatePluralName}}:all [get]
func (a *{{.Name}}QueryApi) FindAll(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := {{.Name}}Assembler.AssFindAllRequest(ictx)
    	if err != nil {
    		return nil, false, err
    	}
		fpr, b, e := a.queryService.FindAll(ctx, req.TenantId)
		return {{.Name}}Assembler.AssFindAllResponse(ictx, fpr, b, e)
	})
}


// FindPaging godoc
// @Summary      分页查询
// @Description  分页查询
// @Tags         {{$AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.{{.Name}}FindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/{{$AggregatePluralName}} [get]
func (a *{{.Name}}QueryApi) FindPaging(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := {{.Name}}Assembler.AssFindPagingRequest(ictx)
    	if err != nil {
    		return nil, false, err
    	}
		fpr, b, e := a.queryService.FindPaging(ctx, req)
		return {{.Name}}Assembler.AssFindPagingResponse(ictx, fpr, b, e)
	})
}

