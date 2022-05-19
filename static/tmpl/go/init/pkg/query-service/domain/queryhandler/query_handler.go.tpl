package queryhandler
{{$AggregateName := .AggregateName}}
{{$EventPackage := .EventPackage}}
import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/event/{{.EventPackage}}"
	"{{.Namespace}}/pkg/query-service/domain/factory/{{.FactoryPackage}}"
	"{{.Namespace}}/pkg/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type {{.AggregateName}}QueryHandler struct {
	{{.aggregateName}}Service queryservice.{{.AggregateName}}QueryService
{{- range $entityName, $entity := .Entities}}
	{{$entity.FirstLowerName}}Service queryservice.{{$entity.Name}}QueryService
{{- end }}
	restapp.BaseQueryHandler
}

func New{{.AggregateName}}Subscribes() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
	{{- range $eventName, $event := .Events}}
		{PubsubName: "pubsub", Topic: {{$EventPackage}}.{{$event.EventType}}Type.String(), Route: "/event/command-service/users/user-create-event"},
	{{- end }}
	}
	return restapp.NewRegisterSubscribe(subscribes, New{{.AggregateName}}QueryHandler())
}

func New{{.AggregateName}}QueryHandler() ddd.QueryEventHandler {
	return &{{.AggregateName}}QueryHandler{
		{{.aggregateName}}Service: queryservice.New{{.AggregateName}}QueryService(),
		{{- range $entityName, $entity := .Entities}}
		{{$entity.FirstLowerName}}Service: queryservice.New{{$entityName}}QueryService(),
		{{- end }}
	}
}
{{- $FactoryPackage := .FactoryPackage}}
{{- range $eventName, $event := .Events}}
func (h *{{$AggregateName}}QueryHandler) On{{$eventName}}(ctx context.Context, event *{{$EventPackage}}.{{$eventName}}) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		view := {{$FactoryPackage}}.NewAddressViewByUserCreateEventV1(event)
		return h.addrService.Create(ctx, view)
	})
}
{{- end}}


func (h *{{.AggregateName}}QueryHandler) GetStructName() string {
	return "{{.ServiceName}}.{{.AggregateName}}QueryHandler"
}
