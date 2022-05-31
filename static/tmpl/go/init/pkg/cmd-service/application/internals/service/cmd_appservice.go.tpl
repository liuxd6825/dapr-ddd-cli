package {{.aggregate_name}}_service

import (
	"context"
	command "{{.Namespace}}/pkg/cmd-service/domain/command/{{.AggregateCommandPackage}}"
	model "{{.Namespace}}/pkg/cmd-service/domain/model/{{.aggregate_name}}_model"
	domain_service "{{.Namespace}}/pkg/cmd-service/domain/service/{{.aggregate_name}}_service"
)

type {{.ClassName}} struct {
	domainService *domain_service.{{.Aggregate.Name}}CommandDomainService
}

//
// New{{.ClassName}}
// @Description:  {{.Description}}
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
	return &{{.ClassName}}{
		domainService: &domain_service.{{.Aggregate.Name}}CommandDomainService{},
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
func (s *{{$ClassName}}) {{$cmd.ServiceFuncName}}(ctx context.Context, cmd *command.{{$cmdName}}) error {
	if err := cmd.Validate(); err != nil {
		return err
	}
	if cmd.GetIsValidOnly() {
		return nil
	}
	_, err := s.domainService.{{$cmd.ServiceFuncName}}(ctx, cmd)
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
	return s.domainService.GetAggregateById(ctx, tenantId, id)
}
