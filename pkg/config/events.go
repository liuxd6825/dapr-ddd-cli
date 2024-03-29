package config

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Events map[string]*Event

func (e *Events) isNil() bool {
	return e == nil || *e == nil
}

func (e *Events) Adds(events Events) {
	if e == nil || events == nil {
		return
	}
	em := *e
	for name, event := range events {
		if event != nil {
			if _, ok := em[name]; !ok {
				em[name] = event
			}
		}
	}
}

func (e *Events) init(a *Aggregate) {
	if e.isNil() {
		return
	}
	for name, event := range *e {
		if event.Properties == nil {
			event.Properties = make(Properties)
		}
		event.init(a, name)
	}
}

func (e *Events) Find(name string) (*Event, bool) {
	if e == nil {
		return nil, false
	}
	m := *e
	v, ok := m[name]
	return v, ok
}

func (e *Events) GetEventTypes() *[]string {
	typesMap := map[string]string{}
	var res []string
	if e.isNil() {
		return &res
	}

	for _, event := range *e {
		_, ok := typesMap[event.EventType]
		if !ok {
			typesMap[event.EventType] = event.EventType
			res = append(res, event.EventType)
		}
	}
	return &res
}

func (e *Events) GetAggregateEvents() []*Event {
	var events []*Event
	if e.isNil() {
		return events
	}

	for _, event := range *e {
		if event.To == "" || strings.ToLower(event.To) == strings.ToLower(event.Aggregate.Name) {
			events = append(events, event)
		}
	}

	return events
}

func (e *Events) GetEntityEvents(entityName string) []*Event {
	var events []*Event
	if e.isNil() {
		return events
	}

	for _, event := range *e {
		if strings.ToLower(event.To) == strings.ToLower(entityName) {
			events = append(events, event)
		}
	}

	return events
}

type Event struct {
	Name                string     // 事件名称
	AggregateId         string     `yaml:"aggregateId"` // 聚合id属性名称
	EventType           string     `yaml:"eventType"`   // 事件类型
	Action              string     `yaml:"action"`      // 活动类型: create, update, delete
	Version             string     `yaml:"version"`     // 版本号， 默认：V1.0
	To                  string     `yaml:"to"`          // 事件所应用到对象类型
	Description         string     `yaml:"description"` // 事件说明
	Properties          Properties `yaml:"properties"`  // 属性
	hasDataProperty     bool
	DataProperty        *Property   // 关联的数据属性
	DataFields          *Fields     // 关联的数据字段
	DataFieldProperties *Properties // 关联的数据字段
	ItemFieldProperties *Properties
	Aggregate           *Aggregate // 聚合
	Route               string     // dapr消息监听的web地址
}

func (e *Event) init(a *Aggregate, name string) {
	if len(e.Version) == 0 {
		e.Version = "v1.0"
	}
	e.Aggregate = a
	e.Name = name
	e.Route = fmt.Sprintf("%s/%s", utils.SnakeString(e.Aggregate.Name), utils.SnakeString(e.Name))
	e.hasDataProperty = false
	e.Version = strings.ToLower(e.Version)
	e.Properties.Init(a, a.Config)

	data := e.Properties["data"]
	if data == nil {
		data = e.Properties["Data"]
	}

	if data != nil {
		e.DataProperty = data
		e.hasDataProperty = true
		if a.FieldsObjects != nil && e.DataProperty.Type != "" {
			fields := a.FieldsObjects[e.DataProperty.Type]
			if fields != nil {
				e.DataFields = fields
				e.DataFieldProperties = &fields.Properties

				if itemsProperty, ok := fields.Properties.Find("Items"); ok {
					fieldItem := a.FieldsObjects[itemsProperty.Type]
					if fieldItem != nil {
						e.ItemFieldProperties = &fieldItem.Properties
					} else {
						e.ItemFieldProperties = &Properties{}
					}
				}
			}
		}
	}
	if e.DataProperty == nil {
		e.DataProperty = &Property{}
	}
	if e.DataFields == nil {
		e.DataFields = &Fields{}
	}
	if e.DataFieldProperties == nil {
		e.DataFieldProperties = &Properties{}
	}
	e.initEventTo()
}

func (e *Event) initEventTo() {
	if e.To != "" {
		return
	}
	e.To = e.getEntityNameByEventName()
}

func (e *Event) getEntityNameByEventName() string {
	if name, ok := e.getEntityName("CreateEvent"); ok {
		return name
	} else if name, ok := e.getEntityName("UpdateEvent"); ok {
		return name
	} else if name, ok := e.getEntityName("DeleteEvent"); ok {
		return name
	}
	return ""
}

func (e *Event) getEntityName(eventTypeName string) (string, bool) {
	index := strings.Index(e.Name, eventTypeName)
	if index > -1 {
		return e.Name[:index], true
	}
	return "", false
}

func (e *Event) ClassName() string {
	return e.Name
}

func (e *Event) MethodName() string {
	methodName := e.Name
	if strings.HasSuffix(methodName, "Event") {
		methodName = methodName[0 : len(methodName)-5]
	}
	return methodName
}

//
// EventSourcingHandler
// @Description: 根据事件类型名称获取接受事件方法名称
// @param eventType
// @param revision
// @return string
//
func (e *Event) EventSourcingHandler() string {
	names := strings.Split(e.Name, ".")
	name := names[len(names)-1]
	ver := strings.Replace(e.Version, ".", "s", -1)
	if strings.HasPrefix(ver, "v") || strings.HasPrefix(ver, "V") {
		ver = "V" + ver[1:]
	} else {
		ver = "V" + ver
	}
	return fmt.Sprintf("On%s%s", name, ver)
}

func (e *Event) FirstUpperName() string {
	return utils.FirstUpper(e.Name)
}

func (e *Event) FirstLowerName() string {
	return utils.FirstLower(e.Name)
}

func (e *Event) IsCreate() bool {
	if e.Action == "" && strings.Contains(e.Name, "CreateEvent") {
		return true
	}
	return strings.ToLower(e.Action) == "create"
}

func (e *Event) IsUpdate() bool {
	if e.Action == "" && strings.Contains(e.Name, "UpdateEvent") {
		return true
	}
	return strings.ToLower(e.Action) == "update"
}

func (e *Event) IsDelete() bool {
	if e.Action == "" && strings.Contains(e.Name, "DeleteEvent") {
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
	if e.To == "" || strings.ToLower(e.To) == strings.ToLower(e.Aggregate.Name) {
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
	if e.To == "" {
		return false
	}
	res := strings.ToLower(e.To) == strings.ToLower(entityName)
	return res
}

func (e *Event) HasDataProperty() bool {
	return e.hasDataProperty
}

func (e *Event) IsHasDataProperty() bool {
	return e.hasDataProperty
}

func (e *Event) SnakeName() string {
	return utils.SnakeString(e.Name)
}

//
// ToPluralName
// @Description: To的复数名称
// @receiver e
// @return string
//
func (e *Event) ToPluralName() string {
	return utils.FirstUpper(utils.Plural(e.To))
}

func (e *Event) DataIsItems() bool {
	b := e.DataFieldProperties.HasItems()
	return b
}

func (e *Event) IsAggregateDeleteByIdEvent() bool {
	if e != nil && e.IsAggregate() && strings.HasPrefix(e.Name, e.Aggregate.Name+"DeleteEvent") {
		return true
	}
	return false
}

func (e *Event) IsAggregateCreateEvent() bool {
	if e != nil && e.IsAggregate() && strings.HasPrefix(e.Name, e.Aggregate.Name+"CreateEvent") {
		return true
	}
	return false
}

func (e *Event) IsAggregateUpdateEvent() bool {
	if e != nil && e.IsAggregate() && strings.HasPrefix(e.Name, e.Aggregate.Name+"UpdateEvent") {
		return true
	}
	return false
}

func (e *Event) IsAggregateCustomEvent() bool {
	if e != nil && e.IsAggregate() && !e.IsAggregateDeleteByIdEvent() && !e.IsAggregateCreateEvent() && !e.IsAggregateUpdateEvent() {
		return true
	}
	return false
}

func (e *Event) IsEntityDeleteByIdEvent() bool {
	if e != nil && !e.IsAggregate() && strings.HasPrefix(e.Name, e.To+"DeleteEvent") {
		return true
	}
	return false
}

func (e *Event) IsEntityCreateEvent() bool {
	if e != nil && !e.IsAggregate() && strings.HasPrefix(e.Name, e.To+"CreateEvent") {
		return true
	}
	return false
}

func (e *Event) IsEntityUpdateEvent() bool {
	if e != nil && !e.IsAggregate() && strings.HasPrefix(e.Name, e.To+"UpdateEvent") {
		return true
	}
	return false
}

func (e *Event) IsEntityCustomEvent() bool {
	if e != nil && e.IsAggregate() && !e.IsEntityDeleteByIdEvent() && !e.IsEntityCreateEvent() && !e.IsEntityUpdateEvent() {
		return true
	}
	return false
}
