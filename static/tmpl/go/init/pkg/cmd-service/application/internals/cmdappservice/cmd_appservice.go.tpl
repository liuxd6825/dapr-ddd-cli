package cmdappservice

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/command/{{.CommandPackage}}"
	"{{.Namespace}}/pkg/cmd-service/domain/model"
	domain_service "{{.Namespace}}/pkg/cmd-service/domain/service"
)

type {{.ClassName}} struct {
	{{.Aggregate.FirstLowerName}}DomainService *domain_service.{{.Aggregate.Name}}DomainService
}


//
// New{{.ClassName}}()
// @Description:  {{.Description}}
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
	return &{{.ClassName}}{
		{{.Aggregate.FirstLowerName}}DomainService: &domain_service.{{.Aggregate.Name}}DomainService{},
	}
}

{{- $ClassName := .ClassName}}
{{- $CommandPackage := .CommandPackage}}
{{- range $cmdName, $cmd := .Commands}}
//
// {{$cmd.ServiceFuncName}}
// @Description: {{$cmd.Description}}
// @receiver s
// @param ctx 上下文
// @param cmd {{$cmd.Description}}
// @return error
//
func (s *{{$ClassName}}) {{$cmd.ServiceFuncName}}(ctx context.Context, cmd *{{$CommandPackage}}.{{$cmdName}}) error {
	if err := cmd.Validate(); err != nil {
		return err
	}
	if cmd.GetIsValidOnly() {
		return nil
	}
	_, err := s.{{.Aggregate.FirstLowerName}}DomainService.{{$cmd.ServiceFuncName}}(ctx, cmd)
	return err
}
{{- end }}

//
// GetAggregateById
// @Description:
// @receiver s
// @param ctx 上下文
// @param tenantId 租户Id
// @param id 聚合根Id
// @return error
//
func (s *{{.ClassName}}) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.{{.Aggregate.Name}}Aggregate, bool, error) {
	return s.{{.Aggregate.FirstLowerName}}DomainService.GetAggregateById(ctx, tenantId, id)
}
