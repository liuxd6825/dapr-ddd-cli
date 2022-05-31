package {{.aggregate_name}}_handler

{{- $AggregateName := .AggregateName}}
{{- $EventPackage := .AggregateEventPackage}}
import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/event/{{.AggregateEventPackage}}"
	view "{{.Namespace}}/pkg/query-service/domain/projection/{{.aggregate_name}}_view"
	domain_factory "{{.Namespace}}/pkg/query-service/domain/factory/{{.AggregateFactoryPackage}}"
	domain_service "{{.Namespace}}/pkg/query-service/domain/service/{{.aggregate_name}}_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.Name}}QueryHandler struct {
	domainService domain_service.{{.Name}}QueryDomainService
	restapp.BaseQueryHandler
}

func New{{.Name}}Subscribes() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
	{{- range $eventName, $event := .Events}}
	{{- if $event.IsAggregate }}
		{PubsubName: "pubsub", Topic: {{$EventPackage}}.{{$event.EventType}}Type.String(), Route: "/event/command-service/{{$event.Route}}"},
	{{- end }}
	{{- end }}
	}
	return restapp.NewRegisterSubscribe(subscribes, New{{.Name}}QueryHandler())
}

func New{{.Name}}QueryHandler() ddd.QueryEventHandler {
	return &{{.Name}}QueryHandler{
		domainService: domain_service.New{{.Name}}QueryDomainService(),
	}
}
{{- $FactoryPackage := .AggregateFactoryPackage}}
{{- range $event := .Events}}
{{- if $event.IsAggregate }}
func (h *{{$AggregateName}}QueryHandler) On{{$event.Name}}(ctx context.Context, event *{{$EventPackage}}.{{$event.Name}}) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		v := domain_factory.New{{$AggregateName}}ViewBy{{$event.Name}}(event)
		{{- if $event.IsCreate }}
		return h.domainService.Create(ctx, v)
		{{- end}}
        {{- if $event.IsUpdate }}
        return h.domainService.Update(ctx, v)
        {{- end}}
        {{- if $event.IsDelete }}
        return h.domainService.DeleteById(ctx, v.TenantId, v.Id)
        {{- end}}
	})
}
{{- end }}
{{- end }}


func (h *{{.Name}}QueryHandler) GetStructName() string {
	return "{{.ServiceName}}.{{.Name}}QueryHandler"
}
