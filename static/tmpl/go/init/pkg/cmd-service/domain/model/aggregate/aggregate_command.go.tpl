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
    e, err := factory.Event.New{{$cmd.EventName}}(ctx, cmd, metadata)
    if err!=nil {
        return nil, err
    }
    {{- if $cmd.IsCreateAggregate }}
    return ddd.CreateEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
    {{- else if  $cmd.IsUpdateAggregate }}
    return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
    {{- else if  $cmd.IsDeleteAggregate }}
    return ddd.DeleteEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
    {{- else }}
    return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
    {{- end }}
}
{{- end }}
