package model
{{- $ClassName := .ClassName }}
{{- $EventPackage := .EventPackage}}
{{- $CommandPackage := .CommandPackage}}

import (
    "context"
    {{.time}}
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/command"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/factory"
    "{{.Namespace}}/pkg/cmd-service/infrastructure/utils"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description:  {{.Description}} 聚合类型
//
type {{.ClassName}} struct {
{{- range $name, $property := .Properties}}
    {{- if $property.IsArrayEntityType }}
    {{$property.UpperName}} *{{$property.GoLanType}}Items `json:"{{$property.JsonName}}" copier:"-" validate:"{{$property.GetValidate}}"` // {{$property.Description}}
    {{- else if  $property.IsArray }}
    {{$property.UpperName}} []{{$property.GoLanType}} `json:"{{$property.JsonName}}" validate:"{{$property.GetValidate}}"` // {{$property.Description}}
    {{- else }}
    {{$property.UpperName}} {{$property.GoLanType}} `json:"{{$property.JsonName}}" validate:"{{$property.GetValidate}}"` // {{$property.Description}}
    {{- end }}
{{- end}}
}

const AggregateType = "{{.AggregateType}}"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init(){
    {{- range $name, $property := .Properties}}
    {{- if $property.IsArray}}
    aggMapperRemove = append(aggMapperRemove, "{{$property.UpperName}}")
    {{- end}}
    {{- end}}
}

//
// New{{.ClassName}}
// @Description: 新建{{.Description}} 聚合对象
// @return *{{.ClassName}}
//
func New{{.ClassName}}() *{{.ClassName}} {
    return &{{.ClassName}}{
    {{- range $name, $property := .Properties}}
        {{- if $property.IsArray}}
        {{$property.UpperName}} : New{{$property.LanType}}Items() ,
        {{- end}}
    {{- end}}
    }
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return New{{.ClassName}}()
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *{{.ClassName}}) GetAggregateVersion() string {
    return "{{.Aggregate.Version}}"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *{{.ClassName}}) GetAggregateType() string {
    return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *{{.ClassName}}) GetAggregateId() string {
    return a.{{.Aggregate.Id.Name}}
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *{{.ClassName}}) GetTenantId() string {
    return a.TenantId
}
