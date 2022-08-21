package model
{{- $ClassName := .ClassName }}
{{- $EventPackage := .EventPackage}}
{{- $CommandPackage := .CommandPackage}}

import (
    "context"
    {{.time}}
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/factory"
    "{{.Namespace}}/pkg/cmd-service/infrastructure/utils"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

{{- range $eventName, $event := .Events }}
//
// {{$event.EventSourcingHandler}}
// @Description: {{$event.Name}} {{$event.Description}} 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *{{$ClassName}}) {{$event.EventSourcingHandler}}(ctx context.Context, e *event.{{$event.Name}}) error {
    {{- if $event.IsAggregateCreateEvent }}
    return utils.Mapper(e.Data, a)
    {{- else if $event.IsAggregateUpdateEvent }}
    return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
    {{- else if $event.IsAggregateDeleteByIdEvent }}
	a.IsDeleted = true
	return nil
    {{- else if $event.IsEntityCreateEvent }}
    {{- if $event.DataIsItems }}
    for _, item := range e.Data.Items {
        if _, err := a.SaleItems.AddMapper(ctx, item.Id, item); err != nil {
            return err
        }
    }
    return nil
    {{- else }}
    _, err := a.{{$event.ToPluralName}}.AddMapper(ctx, e.Data.Id, &e.Data)
    return err
    {{- end }}
    {{- else if $event.IsEntityUpdateEvent }}
    {{- if $event.DataIsItems }}
    for _, item := range e.Data.Items {
        if _, err := a.SaleItems.AddMapper(ctx, item.Id, item); err != nil {
            return err
        }
    }
    return nil
    {{- else }}
    _, _, err := a.{{$event.ToPluralName}}.UpdateMapper(ctx, e.Data.Id, &e.Data, e.UpdateMask)
    return err
    {{- end }}
    {{- else if $event.IsEntityDeleteByIdEvent}}
    {{- if $event.DataIsItems }}
    for _, item := range e.Data.Items {
        if err := a.SaleItems.DeleteById(ctx, item.Id); err != nil {
            return err
        }
    }
    return nil
    {{- else }}
    return a.{{$event.ToPluralName}}.DeleteById(ctx, e.Data.Id)
    {{- end }}
    {{- else }}
	panic("{{$ClassName}}.On{{$event.Name}} to={{$event.To}} error")
	return nil
    {{- end }}
}
{{- end }}

