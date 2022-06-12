{{- $ClassName := .ClassName }}
{{- $EventPackage := .EventPackage}}
{{- $CommandPackage := .CommandPackage}}
package model

import (
    "context"
    {{- if .Properties.HasTimeType }}
    "time"
    {{- end}}
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description:  {{.Description}} 聚合类型
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}map[string]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` // {{$property.Description}}
{{- end}}
}

const AggregateType = "{{.AggregateType}}"

//
// New{{.ClassName}}
// @Description: 新建{{.Description}} 聚合对象
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{}
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return New{{.ClassName}}()
}

{{- range $cmdName, $cmd := .Commands }}

//
// {{$cmd.Name}}
// @Description: 执行 {{$cmd.Name}} {{$cmd.Description}} 命令
// @receiver a
// @param ctx 上下文
// @param cmd {{$cmd.Name}} 命令
// @param metadata 元数据
// @return error 错误
//
func (a *{{$ClassName}}) {{$cmd.Name}}(ctx context.Context, cmd *command.{{$cmd.Name}}, metadata *map[string]string) error {
    {{- if $cmd.IsCreateAggregate }}
    return ddd.CreateEvent(ctx, a, cmd.NewDomainEvent(), ddd.NewApplyEventOptions(metadata))
    {{- else if  $cmd.IsUpdateAggregate }}
    return ddd.ApplyEvent(ctx, a, cmd.NewDomainEvent(), ddd.NewApplyEventOptions(metadata))
    {{- else if  $cmd.IsDeleteAggregate }}
    return ddd.DeleteEvent(ctx, a, cmd.NewDomainEvent(), ddd.NewApplyEventOptions(metadata))
    {{- else }}
    return ddd.ApplyEvent(ctx, a, cmd.NewDomainEvent(), ddd.NewApplyEventOptions(metadata))
    {{- end }}
}
{{- end }}

{{- range $eventName, $event := .Events }}
//
// On{{$event.Name}}
// @Description: {{$event.Name}} {{$event.Description}} 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *{{$ClassName}}) On{{$event.Name}}(ctx context.Context, e *event.{{$event.Name}}) error {
    {{- if $event.IsCreateOrUpdate }}
        {{- range $propName, $prop := $event.DataFields.Properties }}
    a.{{$propName}} = e.Data.{{$propName}}
	    {{- end }}
	{{- end }}
    return nil
}
{{- end }}


//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *{{.ClassName}}) GetAggregateVersion() string {
    return "{{.Aggregate.Version}}"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *{{.ClassName}}) GetAggregateType() string {
    return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *{{.ClassName}}) GetAggregateId() string {
    return a.{{.Aggregate.Id.Name}}
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *{{.ClassName}}) GetTenantId() string {
    return a.TenantId
}
