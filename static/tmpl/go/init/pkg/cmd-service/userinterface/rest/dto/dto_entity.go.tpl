package dto

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
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
	CommandId   string                `json:"commandId"`
	IsValidOnly bool                  `json:"isValidOnly"`
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
    {{- if $property.IsTimeType }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
    {{- else if $property.IsDateType }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
    {{- else if $property.IsEntityType }}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}}Dto `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
    {{- else }}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
    {{- end }}
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



