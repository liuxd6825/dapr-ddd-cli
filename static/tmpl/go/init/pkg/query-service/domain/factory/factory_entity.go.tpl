{{- $aggregate_name := .aggregate_name}}
{{- $Name := .Name}}
{{- $name := .name}}
{{- $defaultProperties := .DefaultProperties}}
package factory

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{$aggregate_name}}/event"
	"{{.Namespace}}/pkg/query-service/domain/{{$aggregate_name}}/view"
	"{{.Namespace}}/pkg/query-service/infrastructure/utils"

)

type {{.name}}ViewFactory struct {
}

var {{.Name}}View = &{{.name}}ViewFactory{}

{{- range $eventName, $event := .Events }}

func (f *{{$name}}ViewFactory) NewBy{{$event.Name}}(ctx context.Context, e *event.{{$event.Name}}) (*view.{{$Name}}View, error) {
    v := &view.{{$Name}}View{}
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