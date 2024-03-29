package dto
{{- $aggregate := .Aggregate}}
import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
	"{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
)

{{- range $i, $cmd := .Commands}}

// {{$cmd.Name}}

//
// {{$cmd.Name}}Request
// @Description: {{$cmd.Description}}
//
type {{$cmd.Name}}Request struct {
	CommandId   string                `json:"commandId"  validate:"required"`     // 命令ID
	IsValidOnly bool                  `json:"isValidOnly"  validate:"-"`   // 是否仅验证，不执行
    {{- if $cmd.IsUpdate }}
	UpdateMask  []string              `json:"updateMask"  validate:"-"`    // 要更新的字段项，空值：更新所有字段
    {{- end }}
	Data        {{$cmd.Name}}RequestData `json:"data"  validate:"required"`
}

//
// {{$cmd.Name}}RequestData
// @Description: {{$cmd.Description}}
//
type {{$cmd.Name}}RequestData struct {
{{- range $name, $property := $cmd.Properties.GetDataFieldProperties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} []*{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEntityType}}
    {{$property.UpperName}} *{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsDates }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsTimes }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEnumType }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsItems}}
    {{$property.UpperName}} []*{{$property.LanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- end}}
{{- end}}
}

{{- if $cmd.Properties.GetDataFieldProperties.HasItems }}
{{- $itemProperty := $cmd.Properties.GetDataFieldProperties.GetItems }}
{{- if  $aggregate.HasType $itemProperty.Type }}
{{- $type := $aggregate.FindType $itemProperty.Type}}


//
// {{$itemProperty.LanType}}
// @Description: {{$itemProperty.Description}}
//
type {{$itemProperty.LanType}} struct {
{{- range $name, $property := $type.Properties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} []*{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEntityType}}
    {{$property.UpperName}} *{{$property.LanType}}Dto `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsDates }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsTimes }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEnumType }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsItems}}
    {{$property.UpperName}} []*{{$property.LanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- end}}
{{- end}}
}
{{- end }}
{{- end }}

//
// {{$cmd.Name}}Response
// @Description: {{$cmd.Description}}
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
    {{- else if $property.IsDates }}
    {{$property.UpperName}} *types.JSONDate `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsTimes }}
    {{$property.UpperName}} *types.JSONTime `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else if $property.IsEnumType }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.GoLanType}}  `json:"{{$property.LowerName}},omitempty" validate:"{{$property.GetValidate}}"`  // {{$property.Description}}
    {{- end}}
{{- end}}
}


