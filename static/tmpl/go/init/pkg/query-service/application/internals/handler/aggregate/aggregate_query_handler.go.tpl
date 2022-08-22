package handler

{{- $AggregateName := .AggregateName}}
{{- $EventPackage := .AggregateEventPackage}}
import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/factory"
	"{{.Namespace}}/pkg/query-service/application/internals/{{.aggregate_name}}/service"
	"{{.Namespace}}/pkg/query-service/infrastructure/base/application/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}QueryHandler struct {
	service *service.{{.Name}}QueryAppService
	handler.BaseQueryHandler
}

{{- $serviceName := .ServiceName}}
//
// New{{.Name}}Subscribe
// @Description: 创建dapr消息订阅器，用于接受领域事件
// @return restapp.RegisterSubscribe  消息注册器
//
func New{{.Name}}Subscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
	{{- range $eventName, $event := .Events}}
	{{- if $event.IsAggregate }}
		{PubsubName: "pubsub", Topic: event.{{$event.EventType}}Type.String(), Route: "/{{$serviceName}}/domain-event/{{$event.Route}}"},
	{{- end }}
	{{- end }}
	}
	return restapp.NewRegisterSubscribe(subscribes, New{{.Name}}QueryHandler())
}

//
// New{{.Name}}QueryHandler
// @Description: 创建{{.Description}}领域事件处理器
// @return ddd.QueryEventHandler 领域事件处理器
//
func New{{.Name}}QueryHandler() ddd.QueryEventHandler {
	return &{{.Name}}QueryHandler{
		service: service.Get{{.Name}}QueryAppService(),
	}
}
{{- $FactoryPackage := .AggregateFactoryPackage}}
{{- range $event := .Events}}
{{- if $event.IsAggregate }}

//
// On{{$event.Name}}
// @Description: {{$event.Name}}事件处理器
// @receiver h
// @param ctx 上下文
// @param event {{$event.Name}} {{$event.Description}}
// @return error 错误
//
func (h *{{$AggregateName}}QueryHandler) On{{$event.EventSourcingHandler}}(ctx context.Context, event *event.{{$event.Name}}) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		{{- if $event.IsAggregateCreateEvent }}
        v, err := factory.{{$AggregateName}}View.NewBy{{$event.Name}}(ctx, event)
        if err != nil {
            return err
        }
        {{- if $event.DataIsItems }}
        return h.service.CreateMany(ctx, v)
        {{- else }}
        return h.service.Create(ctx, v)
        {{- end }}

		{{- else if $event.IsAggregateUpdateEvent }}
        v, err := factory.{{$AggregateName}}View.NewBy{{$event.Name}}(ctx, event)
        if err != nil {
            return err
        }
        {{- if $event.DataIsItems }}
        return h.service.UpdateMany(ctx, v)
        {{- else }}
        return h.service.Update(ctx, v)
        {{- end }}

        {{- else if $event.IsAggregateDeleteByIdEvent }}
        {{- if $event.DataIsItems }}
        v, err := factory.{{$AggregateName}}View.NewBy{{$event.Name}}(ctx, event)
        if err != nil {
            return err
        }
        return h.service.DeleteMany(ctx, event.GetTenantId(), v)
        {{- else }}
        return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
        {{- end }}

        {{- else }}
        v, err := factory.{{$AggregateName}}View.NewBy{{$event.Name}}(ctx, event)
        if err != nil {
            return err
        }
        return h.service.{{$event.MethodName}}(ctx, v)
        {{- end }}
	})
}
{{- end }}
{{- end }}

func (h *{{.Name}}QueryHandler) GetStructName() string {
	return "{{.ServiceName}}.{{.Aggregate.SnakeName}}.{{.Name}}QueryHandler"
}
