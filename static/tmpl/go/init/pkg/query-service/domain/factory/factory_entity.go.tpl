{{- $aggregate_name := .aggregate_name}}
{{- $name := .Name}}
{{- $defaultProperties := .DefaultProperties}}
package {{$aggregate_name}}_factory

import (
	"{{.Namespace}}/pkg/query-service/domain/projection"
	domain_event "{{.Namespace}}/pkg/command-service/domain/event/{{$aggregate_name}}_event"
)
{{- range $eventName, $event := .Events }}

func New{{$name}}ViewBy{{$event.Name}}(event *domain_event.{{$event.Name}}) *projection.{{$name}}View {
	view := projection.{{$name}}View{
    {{- range $propertyName, $property := $event.DataFieldProperties }}
        {{$propertyName}} : event.Data.{{$propertyName}},
    {{- end }}
    {{- range $propertyName, $property := $defaultProperties}}
        {{- if not $property.IsArray}}
        {{$property.UpperName}} : "",
        {{- end}}
    {{- end}}
	}
	return &view
}
{{- end }}