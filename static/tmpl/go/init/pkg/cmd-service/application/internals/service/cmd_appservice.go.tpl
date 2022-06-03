package {{.aggregate_name}}_service

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/model"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/service"
)

type {{.ClassName}} struct {
	domainService *service.{{.Aggregate.Name}}CommandDomainService
}

//
// New{{.ClassName}}
// @Description:  {{.Description}}
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
	return &{{.ClassName}}{
		domainService: &service.{{.Aggregate.Name}}CommandDomainService{},
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
