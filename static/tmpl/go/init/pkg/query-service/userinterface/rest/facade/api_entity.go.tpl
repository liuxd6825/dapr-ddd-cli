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
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{.EntityPluralName}}/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{.AggregatePluralName}}/{{.ParentId}}/{{.EntityPluralName}}", "FindBy{{.AggregateName}}Id")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{.EntityPluralName}}:{{.aggregateMidlineName}}-id/{{.EntityPluralName}}", "FindBy{{.AggregateName}}Id")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{.EntityPluralName}}:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/{{.EntityPluralName}}", "FindPaging")
}


// FindById godoc
// @Summary      按ID获取用户
// @Description  get string by ID
// @Tags         {{.AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "User ID"
// @Success      200        {object}   dto.{{.Name}}FindByIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/{{.AggregatePluralName}}/{id} [get]
func (a *{{.Name}}QueryApi) FindById(ictx iris.Context) {
	_, _, _ = restapp.DoQueryOne(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := {{.Name}}Assembler.AssFindByIdRequest(ictx)
    	if err != nil {
            return nil, false, err
    	}
		v, b, e := a.queryService.FindById(ctx, req.TenantId, req.Id)
		return {{.Name}}Assembler.AssFindByIdResponse(ictx, v, b, e)
	})
}

// FindBy{{.AggregateName}}Id godoc
// @Summary      按{{.AggregateName}}Id获取
// @Description  get string by ID
// @Tags         {{.AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "User ID"
// @Success      200        {object}   dto.{{.AggregateName}}FindBy{{.AggregateName}}IdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/{{.EntityPluralName}}:{{.aggregateMidlineName}}-id/{{.aggregateName}}Id [get]
func (a *{{.Name}}QueryApi) FindBy{{.AggregateName}}Id(ictx iris.Context, tenantId, {{.aggregateName}}Id string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := {{.Name}}Assembler.AssFindBy{{.AggregateName}}IdRequest(ictx)
    	if err != nil {
    		return nil, false, err
    	}
		v, b, e := a.queryService.FindBy{{.AggregateName}}Id(ctx, req.TenantId, req.{{.AggregateName}}Id)
		return {{.Name}}Assembler.AssFindBy{{.AggregateName}}IdResponse(ictx, v, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有
// @Description  get string by ID
// @Tags         {{.AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.UserFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/{{.AggregatePluralName}}:all [get]
func (a *{{.Name}}QueryApi) FindAll(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(c context.Context) (interface{}, bool, error) {
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
// @Tags         {{.AggregatePluralName}}
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.{{.Name}}FindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/{{.AggregatePluralName}} [get]
func (a *{{.Name}}QueryApi) FindPaging(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := {{.Name}}Assembler.AssFindPagingRequest(ictx)
    	if err != nil {
    		return  nil, false, err
    	}
		fpr, b, e := a.queryService.FindPaging(ctx, req)
		return {{.Name}}Assembler.AssFindPagingResponse(ictx, fpr, b, e)
	})
}

