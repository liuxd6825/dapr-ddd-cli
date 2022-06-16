package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/dto"
    base "{{.Namespace}}/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
)

type {{.Name}}Assembler struct {
    base.BaseAssembler
}

var {{.Name}} = &{{.Name}}Assembler{}

func (a *{{.Name}}Assembler) AssFindByIdResponse(ctx iris.Context, v *view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindByIdResponse()
	err := mapper.Mapper(v, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *{{.Name}}Assembler) AssFindPagingResponse(ctx iris.Context, v *ddd_repository.FindPagingResult[*view.{{.Name}}View], isFound bool, findErr error) (*dto.{{.Name}}FindPagingResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	response := dto.New{{.Name}}FindPagingResponse()
	err := mapper.Mapper(v, response)
	if err != nil {
		return nil, false, err
	}
	return response, isFound, nil
}

func (a *{{.Name}}Assembler) AssFindAllResponse(ctx iris.Context, vList *[]*view.{{.Name}}View, isFound bool, findErr error) (*dto.{{.Name}}FindAllResponse, bool, error) {
	if findErr != nil  {
		return nil, isFound, findErr
	}
	res := dto.New{{.Name}}FindAllResponse()
	err := mapper.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}
