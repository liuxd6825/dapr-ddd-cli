{{- $aggregate_name := .aggregate_name}}
{{- $name := .Name}}
{{- $defaultProperties := .DefaultProperties}}
package factory

import (
	"context"
	"{{.Namespace}}/pkg/query-service/domain/{{$aggregate_name}}/view"
	"{{.Namespace}}/pkg/cmd-service/domain/{{$aggregate_name}}/event"
)
{{- range $eventName, $event := .Events }}

func New{{$name}}ViewBy{{$event.Name}}(ctx context.Context, e *event.{{$event.Name}}) *view.{{$name}}View {
	v := &view.{{$name}}View{
    {{- range $propertyName, $property := $event.DataFieldProperties }}
        {{$propertyName}} : e.Data.{{$propertyName}},
    {{- end }}
	}
	utils.SetViewDefaultFields(ctx, v)
	return v
}
{{- end }}