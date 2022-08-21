package factory

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// eventFactory
// @Description: {{.Aggregate.Description}}事件工厂
//
type eventFactory struct {
}

var Event = eventFactory{}

{{- range $i, $cmd := .Commands}}
{{$event := $cmd.Event}}

//
// New{{$event.Name}}
// @Description: 创建{{$event.Description}}
//
func (e eventFactory) New{{$event.Name}} (ctx context.Context, cmd *command.{{$cmd.Name}}, metadata *map[string]string) (*event.{{$event.Name}}, error) {
	err := e.checkNewEventParas("eventFactory.New{{$event.Name}}", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
    v := event.New{{$event.Name}}(cmd.CommandId)
    v.Data = cmd.Data
    return v, nil
}
{{- end }}

//
// checkNewEventParas
// @Description: 检查参数是否正确
//
func (e eventFactory) checkNewEventParas(funcName string, ctx context.Context, cmd interface{}, metadata *map[string]string) error {
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


