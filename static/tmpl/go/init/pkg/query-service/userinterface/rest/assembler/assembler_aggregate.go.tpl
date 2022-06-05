package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"{{.Namespace}}/pkg/query-service/infrastructure/types"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/dto"
)

type {{.Name}}Assembler struct {
}

var {{.Name}} = &{{.Name}}Assembler{}

func (a *{{.Name}}Assembler) AssFindByIdRequest(ctx iris.Context) (*dto.{{.Name}}FindByIdRequest, error) {
	res := &dto.{{.Name}}FindByIdRequest{}
	err := res.Init(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *{{.Name}}Assembler) AssGetPagingRequest(ctx iris.Context) (*dto.{{.Name}}FindPagingRequest, error) {
	res := &dto.{{.Name}}FindPagingRequest{}
	err := res.Init(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *{{.Name}}Assembler) AssFindPagingResponse(ctx iris.Context, v *ddd_repository.FindPagingResult[*view.{{.Name}}View], isFound bool, findErr error) (*dto.{{.Name}}FindPagingResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	dtoList := a.AssViewDtoList(v.Data)
	res := &dto.{{.Name}}FindPagingResponse{}
	err := res.Init(v, dtoList)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, false, err
	}

	return res, isFound, nil
}

func (a *{{.Name}}Assembler) AssFindAllRequest(ctx iris.Context) (*dto.{{.Name}}FindAllRequest, error) {
	res := &dto.{{.Name}}FindAllRequest{}
	err := res.Init(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *{{.Name}}Assembler) AssOneResponse(ctx iris.Context, v *view.{{.Name}}View, isFound bool, err error) (*dto.{{.Name}}ViewDto, bool, error) {
	if err != nil || !isFound {
		return nil, isFound, err
	}
	res := a.AssViewDto(v)
	return res, isFound, nil
}

func (a *{{.Name}}Assembler) AssListResponse(ctx iris.Context, v *[]*view.{{.Name}}View, isFound bool, findErr error) (*[]*dto.{{.Name}}ViewDto, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}

	res := a.AssViewDtoList(v)
	return res, isFound, nil
}

func (a *{{.Name}}Assembler) AssViewDtoList(vList *[]*view.{{.Name}}View) *[]*dto.{{.Name}}ViewDto {
	var dtoList []*dto.{{.Name}}ViewDto
	if vList != nil {
		for _, v := range *vList {
			vDto := a.AssViewDto(v)
			dtoList = append(dtoList, vDto)
		}
	}
	return &dtoList
}

func (a *{{.Name}}Assembler) AssViewDto(v *view.{{.Name}}View) *dto.{{.Name}}ViewDto {
	d := dto.{{.Name}}ViewDto{}
	{{- range $name, $property := .DataFieldsProperties}}
      d.{{$property.Name}} = v.{{$property.Name}}
    {{- end}}
	return &d
}
