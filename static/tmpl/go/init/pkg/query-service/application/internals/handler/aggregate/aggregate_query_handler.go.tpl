package {{.aggregate_name}}_handler

{{- $AggregateName := .AggregateName}}
{{- $EventPackage := .AggregateEventPackage}}
import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/domain/{{.aggregate_name}}/factory"
	"{{.Namespace}}/pkg/query-service/application/internals/service/{{.aggregate_name}}_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}QueryHandler struct {
	service *{{.aggregate_name}}_service.{{.Name}}QueryAppService
	restapp.BaseQueryHandler
}
{{$serviceName := .ServiceName}}
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

func New{{.Name}}QueryHandler() ddd.QueryEventHandler {
	return &{{.Name}}QueryHandler{
		service: {{.aggregate_name}}_service.New{{.Name}}QueryAppService(),
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
func (h *{{$AggregateName}}QueryHandler) On{{$event.Name}}(ctx context.Context, event *event.{{$event.Name}}) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		{{- if $event.IsCreate }}
		v := factory.New{{$AggregateName}}ViewBy{{$event.Name}}(event)
		return h.service.Create(ctx, v)
		{{- end}}
        {{- if $event.IsUpdate }}
        v := factory.New{{$AggregateName}}ViewBy{{$event.Name}}(event)
        return h.service.Update(ctx, v)
        {{- end}}
        {{- if $event.IsDelete }}
        return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
        {{- end}}
	})
}
{{- end }}
{{- end }}

func (h *{{.Name}}QueryHandler) GetStructName() string {
	return "{{.ServiceName}}.{{.Name}}QueryHandler"
}
