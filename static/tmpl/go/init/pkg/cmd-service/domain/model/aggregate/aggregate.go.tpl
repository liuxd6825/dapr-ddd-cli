package model
{{- $ClassName := .ClassName }}
{{- $EventPackage := .EventPackage}}
{{- $CommandPackage := .CommandPackage}}

import (
    "context"
    "time"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "{{.Namespace}}/pkg/cmd-service/domain/scan_file/field"
    "{{.Namespace}}/pkg/cmd-service/infrastructure/utils"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description:  {{.Description}} 聚合类型
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} *{{$property.LanType}}Items `json:"{{$property.JsonName}}" copier:"-" validate:"{{$property.GetValidate}}"` // {{$property.Description}}
    {{- else if  $property.IsArray }}
    {{$property.UpperName}} []{{$property.LanType}} `json:"{{$property.JsonName}}" validate:"{{$property.GetValidate}}"` // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.LanType}} `json:"{{$property.JsonName}}" validate:"{{$property.GetValidate}}"` // {{$property.Description}}
    {{- end }}
{{- end}}
}

const AggregateType = "{{.AggregateType}}"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init(){
    {{- range $name, $property := .Properties}}
    {{- if $property.IsArray}}
    aggMapperRemove = append(aggMapperRemove, "{{$property.UpperName}}")
    {{- end}}
    {{- end}}
}

//
// New{{.ClassName}}
// @Description: 新建{{.Description}} 聚合对象
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{
    {{- range $name, $property := .Properties}}
        {{- if $property.IsArray}}
        {{$property.UpperName}} : New{{$property.LanType}}Items() ,
        {{- end}}
    {{- end}}
    }
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
func (a *{{$ClassName}}) {{$cmd.Name}}(ctx context.Context, cmd *command.{{$cmd.Name}}, metadata *map[string]string) (any, error) {
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
    {{- if $event.IsAggregateCreateEvent }}
    return utils.Mapper(e.Data, a)
    {{- else if $event.IsAggregateUpdateEvent }}
    return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
    {{- else if $event.IsAggregateDeleteByIdEvent }}
	a.IsDeleted = true
	return nil
    {{- else if $event.IsEntityCreateEvent }}
    _, err := a.{{$event.ToPluralName}}.AddMapper(ctx, e.Data.Id, &e.Data)
    return err
    {{- else if $event.IsEntityUpdateEvent }}
    _, _, err := a.{{$event.ToPluralName}}.UpdateMapper(ctx, e.Data.Id, &e.Data, e.UpdateMask)
    return err
    {{- else if $event.IsEntityDeleteByIdEvent}}
    return a.{{$event.ToPluralName}}.DeleteById(ctx, e.Data.Id)
    {{- else }}
	panic("{{$ClassName}}.On{{$event.Name}} to={{$event.To}} error")
	return nil
    {{- end }}
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
