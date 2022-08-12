{{- $aggregate_name := .aggregate_name}}
{{- $defaultProperties := .DefaultProperties}}
{{- $Name := .Name}}
{{- $name := .name}}
package factory

import (
	"context"
	"{{.Namespace}}/pkg/cmd-service/domain/{{$aggregate_name}}/event"
	"{{.Namespace}}/pkg/query-service/infrastructure/utils"
	"{{.Namespace}}/pkg/query-service/domain/{{$aggregate_name}}/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

type {{$name}}ViewFactory struct {
}

var {{$Name}}View = &{{$name}}ViewFactory{}

{{- range $eventName, $event := .Events }}
{{- if $event.DataIsItems }}

func (f *{{$name}}ViewFactory) NewBy{{$event.Name}} (ctx context.Context, e *event.{{$event.Name}}) ([]*view.{{$Name}}View, error) {
	if e == nil || len(e.Data.Items) == 0 {
		return []*view.{{$Name}}View{}, nil
	}

    var vList []*view.{{$Name}}View
    {{- if $event.IsCreate}}
    setViewType := utils.SetViewCreated
    {{- else  if $event.IsUpdate}}
    setViewType := utils.SetViewUpdated
    {{- else if $event.IsDelete}}
    setViewType := utils.SetViewDeleted
    {{- else }}
    setViewType := utils.SetViewOther
    {{- end }}
    for _, item := range e.Data.Items {
        v := &view.{{$Name}}View{}
        {{- range $propertyName, $property := $event.ItemFieldProperties }}
        v.{{$propertyName}} = item.{{$propertyName}}
        {{- end }}
        /*
        if err := utils.Mapper(item, v); err != nil {
            return nil, err
        }*/
        if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
            return nil, err
        }
        vList = append(vList, v)
    }
    return vList, nil
}
{{- else }}

func (f *{{$name}}ViewFactory) NewBy{{$event.Name}} (ctx context.Context, e *event.{{$event.Name}}) (*view.{{$Name}}View, error) {
    if e==nil {
        return nil, errors.New("NewBy{{$event.Name}}(ctx, e) error: e is nil")
    }
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

    {{- range $propertyName, $property := $event.DataFieldProperties }}
    v.{{$propertyName}} = e.Data.{{$propertyName}}
    {{- end }}
    if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err!=nil {
        return nil, err
    }
    return v, nil
}
{{- end }}
{{- end }}