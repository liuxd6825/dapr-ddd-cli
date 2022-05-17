package {{.Package}}

import (
    "{{.Namespace}}/pkg/cmd-service/domain/event/{{.aggregateName}}_events"
    "{{.Namespace}}/pkg/cmd-service/domain/fields/{{.aggregateName}}_fields"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
    ddd.BaseCommand
{{- range $name, $property := .Properties}}
    {{$property.UpperName}}   {{$property.DataType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}

//
// NewDomainEvent
// @Description: 创建领域事件
//
func (c *{{.ClassName}}) NewDomainEvent() ddd.DomainEvent {
    return &user_events.UserUpdateEventV1{
        EventId: c.CommandId,
        Data:    c.Data,
    }
}

//
// GetAggregateId
// @Description: 获取聚合根Id
//
func (c *{{.ClassName}}) GetAggregateId() ddd.AggregateId {
    return ddd.NewAggregateId(c.{{.AggregateId}})
}

//
// GetCommandId
// @Description: 获取命令Id
//
func (c *{{.ClassName}}) GetCommandId() string {
    return c.CommandId
}

//
// GetCommandId
// @Description: 获取租户Id
//
func (c *{{.ClassName}}) GetTenantId() string {
    return c.Data.TenantId
}

//
// GetCommandId
// @Description: 命令数据验证
//
func (c *{{.ClassName}}) Validate() error {
    errs := ddd_errors.NewVerifyError()
    c.BaseCommand.ValidateError(errs)
    if c.Data.TenantId == "" {
        errs.AppendField("data.tenantId", "不能为空")
    }

    /* 添加其他数据检查  */

    return errs.GetError()
}
