package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
    "{{.Namespace}}/pkg/query-service/infrastructure/utils"
    "{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/appquery"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/dto"
    base "{{.Namespace}}/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"{{.Namespace}}/pkg/query-service/infrastructure/utils"
)

type {{.Name}}Assembler struct {
    base.BaseAssembler
}

var {{.Name}} = &{{.Name}}Assembler{}

func (a *{{.Name}}Assembler) AssFindByIdResponse(ictx iris.Context, v *view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindByIdResponse()
	err := utils.Mapper(v, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *{{.Name}}Assembler) AssFindPagingRequest(ctx iris.Context) (*appquery.{{.Name}}FindPagingAppQuery, error) {
	fpr, err := a.BaseAssembler.AssFindPagingRequest(ctx)
	if err != nil {
		return nil, err
	}
	query := &appquery.{{.Name}}FindPagingAppQuery{}
	query.Filter = fpr.Filter
	query.TenantId = fpr.TenantId
	query.Sort = fpr.Sort
	query.PageSize = fpr.PageSize
	query.PageNum = fpr.PageNum
	query.Fields = fpr.Fields
	return query, nil
}

func (a *{{.Name}}Assembler) AssFindPagingResponse(ictx iris.Context, fpr *appquery.{{.Name}}FindPagingResult, isFound bool, findErr error) (*dto.{{.Name}}FindPagingResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindPagingResponse()
	err := utils.Mapper(fpr, res)
	if err != nil {
		return nil, false, err
	}
	return res, isFound, nil
}

func (a *{{.Name}}Assembler) AssFindByIdsResponse(ictx iris.Context, vList []*view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindByIdsResponse, bool, error) {
	if findErr != nil  {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindByIdsResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *{{.Name}}Assembler) AssFindAllResponse(ictx iris.Context, vList []*view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindAllResponse, bool, error) {
	if findErr != nil  {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindAllResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}
