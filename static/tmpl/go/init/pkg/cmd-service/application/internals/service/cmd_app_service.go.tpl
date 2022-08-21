package service

import (
	"context"
    "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/appcmd"
    "{{.Namespace}}/pkg/cmd-service/application/internals/{{.aggregate_name}}/executor"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/model"
	domain_service "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/service"
	query_dto "{{.Namespace}}/pkg/query-service/userinterface/rest/{{.aggregate_name}}/dto"
	base "{{.Namespace}}/pkg/cmd-service/infrastructure/base/application/service"
)

type {{.ClassName}} struct {
    base.BaseQueryAppService
    {{- range $i, $cmd := .Commands}}
    {{$cmd.FirstLowerName}}Executor executor.{{$cmd.Name}}Executor
    {{- end }}
    findAggregateByIdExecutor executor.FindAggregateByIdExecutor
}

//
// New{{.ClassName}}
// @Description:  {{.Description}}
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
	res := &{{.ClassName}}{
        {{- range $i, $cmd := .Commands}}
        {{$cmd.FirstLowerName}}Executor: executor.Get{{$cmd.Name}}Executor(),
        {{- end }}
        findAggregateByIdExecutor: executor.GetFindAggregateByIdExecutor(),
	}
    res.Init("{{.QueryServiceName}}", "{{.Aggregate.PluralName}}", "{{.ApiVersion}}")
    return res
}

{{- $ClassName := .ClassName}}
{{- $CommandPackage := .CommandPackage}}
{{- range $i, $cmd := .Commands}}

//
// {{$cmd.ServiceFuncName}}
// @Description: {{$cmd.Description}}
// @receiver s
// @param ctx 上下文
// @param cmd {{$cmd.Description}}命令DTO对象
// @return error
//
func (s *{{$ClassName}}) {{$cmd.ServiceFuncName}}(ctx context.Context, appCmd *appcmd.{{$cmd.AppName}}) error {
	return s.{{$cmd.FirstLowerName}}Executor.Execute(ctx, appCmd)
}
{{- end }}

//
// FindAggregateById
// @Description:
// @receiver s
// @param ctx 上下文
// @param tenantId 租户Id
// @param id 聚合根Id
// @return error
//
func (s *{{.ClassName}}) FindAggregateById(ctx context.Context, tenantId string, id string) (*model.{{.Aggregate.Name}}Aggregate, bool, error) {
	return s.findAggregateByIdExecutor.Execute(ctx, tenantId, id)
}

//
// QueryById
// @Description: 按id获取{{.Description}}投影类
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param id {{.Description}} Id
// @return data {{.Description}} 信息
// @return isFound 是否找到
// @return err 错误信息
//
func (s *{{.ClassName}}) QueryById(ctx context.Context, tenantId string, id string) (*query_dto.{{.Aggregate.Name}}FindByIdResponse, bool, error) {
	var resp query_dto.{{.AggregateName}}FindByIdResponse
	isFound, err := s.BaseQueryAppService.QueryById(ctx, tenantId, id, &resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}

//
// QueryByIds
// @Description: 按ids获取{{.Description}}投影类
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param ids 多个{{.Description}}Id
// @return data {{.Description}} 信息
// @return isFound 是否找到
// @return err 错误信息
//
func (s *{{.ClassName}}) QueryByIds(ctx context.Context, tenantId string, ids []string) (*query_dto.{{.Name}}FindByIdsResponse, bool, error) {
	var resp query_dto.{{.Name}}FindByIdsResponse
	isFound, err := s.BaseQueryAppService.QueryByIds(ctx, tenantId, ids, &resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}
