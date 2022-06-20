package handler

{{- $AggregateName := .AggregateName}}
{{- $EventPackage := .AggregateEventPackage}}
{{- $entityName := .Name}}
import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/factory"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}QueryHandler struct {
	service *service.{{.Name}}QueryAppService
	restapp.BaseQueryHandler
}
{{$serviceName := .ServiceName}}
func New{{.Name}}Subscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
	{{- range $eventName, $event := .Events}}
		{PubsubName: "pubsub", Topic: event.{{$event.EventType}}Type.String(), Route: "/{{$serviceName}}/domain-event/{{$event.Route}}"},
	{{- end }}
	}
	return restapp.NewRegisterSubscribe(subscribes, New{{.Name}}QueryHandler())
}

func New{{.Name}}QueryHandler() ddd.QueryEventHandler {
	return &{{.Name}}QueryHandler{
		service: service.Get{{.Name}}QueryAppService(),
	}
}
{{- $factoryPackage := .AggregateFactoryPackage}}
{{- range $eventName, $event := .Events}}

//
// On{{$event.Name}}
// @Description: {{$event.Name}}事件处理器
// @receiver h
// @param ctx 上下文
// @param event {{$event.Name}} {{$event.Description}}
// @return error 错误
//
func (h *{{$entityName}}QueryHandler) On{{$event.Name}}(ctx context.Context, event *event.{{$event.Name}}) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {

		{{- if $event.IsEntityCreateEvent }}
        v, err := factory.{{$entityName}}View.NewBy{{$event.Name}}(ctx, event)
        if err != nil {
            return err
        }
		return h.service.Create(ctx, v)
		{{- else if $event.IsEntityUpdateEvent }}
        v, err := factory.{{$entityName}}View.NewBy{{$event.Name}}(ctx, event)
        if err != nil {
            return err
        }
        return h.service.Update(ctx, v)
        {{- else if $event.IsEntityDeleteByIdEvent }}
        return h.service.DeleteById(ctx, event.GetTenantId(), event.Id)
        {{- else }}
        return h.service.{{$event.MethodName}}(ctx, v)
        {{- end }}
	})
}
{{- end }}

func (h *{{$entityName}}QueryHandler) GetStructName() string {
	return "{{.ServiceName}}.{{.Aggregate.SnakeName}}.{{$entityName}}QueryHandler"
}
