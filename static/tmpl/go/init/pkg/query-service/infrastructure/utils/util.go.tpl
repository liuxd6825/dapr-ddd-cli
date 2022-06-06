package utils

import (
    "context"
    {{- if .HasTimeType }}
    "time"
    {{- end}}
)

type ViewDefaultFields interface {
{{- range $name, $property := .DefaultViewProperties}}
{{- if not $property.IsArray}}

    //
    // Get{{$property.UpperName}}
    // @Description: 获取 {{$property.Description}}
    //
    Get{{$property.UpperName}}() {{$property.LanType}}

    //
    // Set{{$property.UpperName}}
    // @Description: 设置 {{$property.Description}}
    //
    Set{{$property.UpperName}}({{$property.LanType}})



{{- end}}
{{- end}}
}


func SetViewDefaultFields(ctx context.Context, viewFields ViewDefaultFields) {

}
