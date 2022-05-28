package {{.entityName}}_queryhandler

{{- $AggregateName := .AggregateName}}
{{- $EventPackage := .EventPackage}}
{{- $EntityName := .EntityName}}
import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/event/{{.EventPackage}}"
	"{{.Namespace}}/pkg/query-service/domain/factory/{{.FactoryPackage}}"
	"{{.Namespace}}/pkg/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.EntityName}}QueryHandler struct {
	service queryservice.{{.EntityName}}QueryService
	restapp.BaseQueryHandler
}

func New{{.EntityName}}Subscribes() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
	{{- range $eventName, $event := .Events}}
	    {{- if $event.IsEntity $EntityName }}
		{PubsubName: "pubsub", Topic: {{$EventPackage}}.{{$event.EventType}}Type.String(), Route: "/event/command-service/{{$event.Route}}"},
		{{- end }}
	{{- end }}
	}
	return restapp.NewRegisterSubscribe(subscribes, New{{.EntityName}}QueryHandler())
}

func New{{.EntityName}}QueryHandler() ddd.QueryEventHandler {
	return &{{.EntityName}}QueryHandler{
		service: queryservice.New{{$EntityName}}QueryService(),
	}
}
{{- $FactoryPackage := .FactoryPackage}}
{{- range $eventName, $event := .Events}}
{{- if $event.IsEntity $EntityName }}
func (h *{{$EntityName}}QueryHandler) On{{$eventName}}(ctx context.Context, event *{{$EventPackage}}.{{$eventName}}) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		view := {{$FactoryPackage}}.New{{$AggregateName}}ViewBy{{$eventName}}(event)
		return h.service.Create(ctx, view)
	})
}
{{- end }}
{{- end }}


func (h *{{$EntityName}}QueryHandler) GetStructName() string {
	return "{{.ServiceName}}.{{$EntityName}}QueryHandler"
}
