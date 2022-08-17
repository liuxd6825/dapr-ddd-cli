package factory

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

{{- range $i, $cmd := .Commands}}
{{$event := $cmd.Event}}

//
// New{{$event.Name}}
// @Description: 创建{{$event.Description}}
//
func New{{$event.Name}} (ctx context.Context, cmd *command.{{$cmd.Name}}, metadata *map[string]string) (*event.{{$event.Name}}, error) {
	err := checkNewEventParas("New{{$event.Name}}", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
    e := event.New{{$event.Name}}(cmd.CommandId)
    e.Data = cmd.Data
    return e, nil
}

{{- end }}

//
// checkNewEventParas
// @Description: 检查参数是否正确
//
func checkNewEventParas(funcName string, ctx context.Context, cmd interface{}, metadata *map[string]string) error {
	if ctx == nil {
		return errors.ErrorOf("%s(ctx, cmd, metadata) error: ctx is nil", funcName)
	}
	if cmd == nil {
		return errors.ErrorOf("%s(ctx, cmd, metadata) error: cmd is nil", funcName)
	}
	if metadata == nil {
		return errors.ErrorOf("%s(ctx, cmd, metadata) error: metadata is nil", funcName)
	}
	return nil
}
