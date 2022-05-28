package config

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Event struct {
	Name         string     // 事件名称
	AggregateId  string     `yaml:"aggregateId"` // 聚合id属性名称
	EventType    string     `yaml:"eventType"`   // 事件类型
	Action       string     `yaml:"action"`      // 活动类型: create, update, delete
	Version      string     `yaml:"version"`     // 版本号， 默认：V1
	To           string     `yaml:"to"`          // 事件所应用到对象类型
	Description  string     `yaml:"description"` // 事件说明
	Properties   Properties `yaml:"properties"`  // 属性
	DataProperty *Property  // 关联的数据属性
	DataFields   *Fields    // 关联的数据字段
	Aggregate    *Aggregate // 聚合
	Route        string     // dapr消息监听的web地址
}

type Events map[string]*Event

func (e *Events) init(a *Aggregate) {
	if e != nil {
		for name, event := range *e {
			event.init(a, name)
		}
	}
}

func (e *Events) GetEventTypes() *[]string {
	typesMap := map[string]string{}
	var res []string
	if e != nil {
		for _, event := range *e {
			_, ok := typesMap[event.EventType]
			if !ok {
				typesMap[event.EventType] = event.EventType
				res = append(res, event.EventType)
			}
		}
	}
	return &res
}

func (e *Event) init(a *Aggregate, name string) {
	if len(e.Version) == 0 {
		e.Version = "V1"
	}

	e.Aggregate = a
	e.Name = name
	e.Route = fmt.Sprintf("%s/%s/ver:%s", e.Aggregate.Name, e.Name, e.Version)
	e.Route = strings.ToLower(e.Route)

	e.Properties.Init(a)
	data := e.Properties["data"]
	if data == nil {
		data = e.Properties["Data"]
	}
	if data != nil {
		e.DataProperty = data
		if a.FieldsObjects != nil && e.DataProperty.Type != "" {
			fields := a.FieldsObjects[e.DataProperty.Type]
			if fields != nil {
				e.DataFields = fields
			}
		}
	}
}

func (e *Event) ClassName() string {
	return e.Name
}

func (e *Event) FirstLowerName() string {
	return utils.FirstLower(e.Name)
}

func (e *Event) IsCreate() bool {
	if e.Action == "" && strings.HasSuffix(strings.ToLower(e.Name), "createevent") {
		return true
	}
	return strings.ToLower(e.Action) == "create"
}

func (e *Event) IsUpdate() bool {
	if e.Action == "" && strings.HasSuffix(strings.ToLower(e.Name), "updateevent") {
		return true
	}
	return strings.ToLower(e.Action) == "update"
}

func (e *Event) IsDelete() bool {
	if e.Action == "" && strings.HasSuffix(strings.ToLower(e.Name), "deleteevent") {
		return true
	}
	return strings.ToLower(e.Action) == "delete"
}

func (e *Event) IsCreateOrUpdate() bool {
	if e.IsUpdate() || e.IsUpdate() {
		return true
	}
	return false
}

//
// IsAggregate
// @Description: 是聚合对象上的事件
// @receiver e
// @return bool
//
func (e *Event) IsAggregate() bool {
	if e.To == "" || e.To == e.Aggregate.Name {
		return true
	}
	return false
}

//
// IsEntity
// @Description: 是实体对象上的事件
// @receiver e
// @param entityName 实体名称
// @return bool
//
func (e *Event) IsEntity(entityName string) bool {
	return e.To == entityName
}
