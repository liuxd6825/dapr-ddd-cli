package view

import (
	"time"

)

//
// BaseView
// @Description: 视图基类
//
type BaseView struct {
{{- range $name, $property := .Properties}}
{{- if not $property.IsArray}}
    {{$property.UpperName}} {{if $property.IsArray}}[]*{{end}}{{$property.LanType}} `json:"{{$property.JsonName}},omitempty"  bson:"{{$property.BsonName}}" {{$property.GormTagName}} {{if $property.HasValidate}} validate:"{{$property.Validate}}"{{- end}}` {{if $property.HasDescription }} // {{$property.Description}}{{ end }}
{{- end}}
{{- end}}
}

{{- range $name, $property := .Properties}}
{{- if not $property.IsArray}}

//
// Get{{$property.UpperName}}
// @Description: 获取 {{$property.Description}}
//
func (v *BaseView) Get{{$property.UpperName}}(){{$property.LanType}} {
    return v.{{$property.UpperName}}
}

//
// Set{{$property.UpperName}}
// @Description: 设置 {{$property.Description}}
//
func (v *BaseView) Set{{$property.UpperName}}(value {{$property.LanType}}) {
    v.{{$property.UpperName}} = value
}
{{- end}}
{{- end}}

{{- if .IsNeo4j }}
//
// BaseNeo4jView
// @Description: 视图基类
//
type BaseNeo4jView struct {
	BaseView
	Nid    int64    `json:"-" bson:"nid"`
	Labels []string `json:"-" bson:"labels"`
}

func (b *BaseNeo4jView) SetNid(v int64) {
	b.Nid = v
}

func (b *BaseNeo4jView) GetNid() int64 {
	return b.Nid
}

func (b *BaseNeo4jView) SetLabels(v []string) {
	b.Labels = v
}

func (b *BaseNeo4jView) GetLabels() []string {
	return b.Labels
}

{{- end }}
