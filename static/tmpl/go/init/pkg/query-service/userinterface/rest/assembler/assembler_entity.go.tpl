package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"

	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/dto"
	"{{.Namespace}}/pkg/query-service/infrastructure/utils"
	base "{{.Namespace}}/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
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

func (a *{{.Name}}Assembler) AssFindPagingResponse(ictx iris.Context, v *ddd_repository.FindPagingResult[*view.{{.Name}}View], isFound bool, findErr error) (*dto.{{.Name}}FindPagingResponse, bool, error) {
    if findErr != nil {
        return nil, isFound, findErr
    }
	response := dto.New{{.Name}}FindPagingResponse()
	err := utils.Mapper(v, response)
	if err != nil {
		return nil, false, err
	}
	return response, isFound, nil
}

func (a *{{.Name}}Assembler) AssFindAllResponse(ictx iris.Context, vList []*view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindAllResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindAllResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}


func (a *{{.Name}}Assembler) AssFindBy{{.AggregateName}}IdRequest(ictx iris.Context) (*dto.{{.Name}}FindBy{{.AggregateName}}IdRequest, error) {
    tenantId, err := a.BaseAssembler.GetTenantId(ictx)
    if err!=nil {
        return nil, err
    }

    {{.aggregateName}}Id, err := a.BaseAssembler.GetIdParam(ictx, "{{.aggregateName}}Id")
    if err!=nil {
        return nil, err
    }

	res := dto.New{{.Name}}FindBy{{.AggregateName}}IdRequest()
	res.TenantId = tenantId
	res.{{.AggregateName}}Id = {{.aggregateName}}Id

	return res, nil
}


func (a *{{.Name}}Assembler) AssFindBy{{.AggregateName}}IdResponse(ictx iris.Context , vList []*view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindBy{{.AggregateName}}IdResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}

	res := dto.New{{.Name}}FindBy{{.AggregateName}}IdResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

