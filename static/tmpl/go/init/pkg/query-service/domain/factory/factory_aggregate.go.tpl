{{- $aggregate_name := .aggregate_name}}
{{- $name := .Name}}
{{- $defaultProperties := .DefaultProperties}}
package factory

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{$aggregate_name}}/event"
	"{{.Namespace}}/pkg/query-service/infrastructure/utils"
	"{{.Namespace}}/pkg/query-service/domain/{{$aggregate_name}}/view"
)

type {{.aggregateName}}ViewFactory struct {
}

var {{.AggregateName}}View = &{{.aggregateName}}ViewFactory{}

{{- range $eventName, $event := .Events }}

func NewBy{{$event.Name}}(ctx context.Context, e *event.{{$event.Name}}) (*view.{{$name}}View, error) {
    v := &view.{{$name}}View{}
    {{- if $event.IsCreate}}
    setViewType := utils.SetViewCreated
    {{- else  if $event.IsUpdate}}
    setViewType := utils.SetViewUpdated
    {{- else if $event.IsDelete}}
    setViewType := utils.SetViewDeleted
    {{- else }}
    setViewType := utils.SetViewOther
    {{- end }}

	if err := utils.ViewMapper(ctx, v, e, setViewType); err != nil {
		return nil, err
	}
	return v, nil

	/*
    {{- range $propertyName, $property := $event.DataFieldProperties }}
    v.{{$propertyName}} = e.Data.{{$propertyName}}
    {{- end }}
    if err := utils.SetViewDefaultFields(ctx, v, e.Time, setViewType); err!=nil {
        return nil, err
    }
    */
}
{{- end }}