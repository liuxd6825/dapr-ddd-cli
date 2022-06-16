package command

import (
	"github.com/google/uuid"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/event"
    "{{.Namespace}}/pkg/cmd-service/domain/{{.aggregate_name}}/field"
    "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// {{.ClassName}}
// @Description: {{.Description}}
//
type {{.ClassName}} struct {
	CommandId   string    `json:"commandId"     validate:"gt=0"`   // 命令ID
	IsValidOnly bool      `json:"isValidOnly"   validate:"gt=0"`   // 是否仅验证，不执行
    {{- if .Command.IsUpdate }}
	UpdateMask  []string  `json:"updateMask"`                      // 要更新的字段项，空值：更新所有字段
    {{- end }}
{{- range $name, $property := .Properties}}
    {{$property.UpperName}} {{if $property.IsData }} field.{{ end }}{{$property.LanType}}   `json:"{{$property.LowerName}}"{{if $property.HasValidate}}  validate:"{{$property.Validate}}"{{- end}}`  // {{$property.Description}}
{{- end}}
}

//
// NewDomainEvent
// @Description: 创建领域事件
//
func (c *{{.ClassName}}) NewDomainEvent() ddd.DomainEvent {
    return &event.{{.EventName}}{
        EventId:    uuid.New().String(),
        CommandId:  c.CommandId,
        {{- if .Command.IsUpdate }}
        UpdateMask: c.UpdateMask,
        {{- end }}
        Data:       c.Data,
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
// GetTenantId
// @Description: 获取租户Id
//
func (c *{{.ClassName}}) GetTenantId() string {
    return c.Data.TenantId
}

//
// GetIsValidOnly
// @Description: 是否只验证不执行。
//
func (c *{{.ClassName}}) GetIsValidOnly() bool {
	return c.IsValidOnly
}

//
// Validate
// @Description: 命令数据验证
//
func (c *{{.ClassName}}) Validate() error {
	ve := ddd.ValidateCommand(c, nil)
    /* 添加其他数据检查
    if c.Data.Id == "" {
        ve.AppendField("data.id", "不能为空")
    }
    */
	return ve.GetError()
}

