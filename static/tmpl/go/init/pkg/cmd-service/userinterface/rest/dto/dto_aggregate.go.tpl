package dto

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

{{- range $i, $cmd := .Commands}}

// {{$cmd.Name}}

//
// {{$cmd.Name}}Request
// @Description: {{$cmd.Description}}
// @receiver c
// @param ctx
//
type {{$cmd.Name}}Request struct {
	CommandId   string                `json:"commandId"`     // 命令ID
	IsValidOnly bool                  `json:"isValidOnly"`   // 是否仅验证，不执行
    {{- if $cmd.IsUpdate }}
	UpdateMask  []string              `json:"updateMask"`    // 要更新的字段项，空值：更新所有字段
    {{- end }}
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
    {{- if $property.TypeIsDateTime}}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.JsonName}}"{{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }}// {{$property.Description}}{{ end }}
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



