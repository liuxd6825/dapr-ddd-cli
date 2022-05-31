{{- $ClassName := .ClassName }}
{{- $EventPackage := .EventPackage}}
{{- $CommandPackage := .CommandPackage}}
package {{.Package}}

import (
    "context"
    _ "time"
    "{{.Namespace}}/pkg/cmd-service/domain/command/{{.AggregateCommandPackage}}"
    "{{.Namespace}}/pkg/cmd-service/domain/event/{{.AggregateEventPackage}}"
    "{{.Namespace}}/pkg/cmd-service/domain/factory/{{.AggregateFactoryPackage}}"
    "github.com/dapr/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description:  {{.Description}}
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsArray}}map[string]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
{{- end}}
}

const AggregateType = "{{.AggregateType}}"

func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{}
}

func NewAggregate() ddd.Aggregate {
	return New{{.ClassName}}()
}

{{- range $cmdName, $cmd := .Commands }}

func (a *{{$ClassName}}) {{$cmd.Name}}(ctx context.Context, cmd *{{$CommandPackage}}.{{$cmd.Name}}, metadata *map[string]string) error {
    return ddd.Apply(ctx, a, cmd.NewDomainEvent(), ddd.ApplyOptions{}.SetMetadata(metadata))
}
{{- end }}

{{- range $eventName, $event := .Events }}

func (a *{{$ClassName}}) On{{$event.Name}}(ctx context.Context, event *{{$EventPackage}}.{{$event.Name}}) error {
    {{- if $event.IsCreateOrUpdate }}
    {{- range $propName, $prop := $event.DataFields.Properties }}
    a.{{$propName}} = event.Data.{{$propName}}
	{{- end }}
	{{- end }}
    return nil
}
{{- end }}

func (a *{{.ClassName}}) GetAggregateVersion() string {
    return "{{.Aggregate.Version}}"
}

func (a *{{.ClassName}}) GetAggregateType() string {
    return AggregateType
}

func (a *{{.ClassName}}) GetAggregateId() string {
    return a.{{.Aggregate.Id.Name}}
}

func (a *{{.ClassName}}) GetTenantId() string {
    return a.TenantId
}
