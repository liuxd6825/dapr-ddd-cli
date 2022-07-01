package dto

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
)

{{- $DefaultProperties := .DefaultProperties}}
{{- range $i, $cmd := .Commands}}

// {{$cmd.Name}}

//
// {{$cmd.Name}}Request
// @Description: {{$cmd.Description}}
// @receiver c
// @param ctx
//
type {{$cmd.Name}}Request struct {
	CommandId   string  `json:"commandId" validate:"required"`
	IsValidOnly bool `json:"isValidOnly" validate:"required"`
	Data        {{$cmd.Name}}RequestData `json:"data"`
}

//
// {{$cmd.Name}}RequestData
// @Description: {{$cmd.Description}}
// @receiver c
// @param ctx
//
type {{$cmd.Name}}RequestData struct {
{{- range $name, $property := $cmd.Properties.GetDataFieldProperties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} []*{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEntityType}}
    {{$property.UpperName}} *{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsDateType }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsTimeType }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEnumType }}
    {{$property.UpperName}} field.{{$property.LanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.LanType}}`json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- end}}
{{- end}}
}

//
// {{$cmd.Name}}Response
// @Description: {{$cmd.Description}}
// @receiver c
// @param ctx
//
type {{$cmd.Name}}Response struct {
}

{{- end }}


//
// {{.Name}}Dto
// @Description: {{.Description}}  请求或响应业务数据
//
type {{.Name}}Dto struct {
{{- range $name, $property := .Properties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} []*{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEntityType}}
    {{$property.UpperName}} *{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsDateType }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsTimeType }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEnumType }}
    {{$property.UpperName}} field.{{$property.LanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.LanType}}`json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- end}}
{{- end}}
}


