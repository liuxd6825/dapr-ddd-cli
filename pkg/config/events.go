package config

import (
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
	"strings"
)

type Event struct {
	Name         string
	AggregateId  string     `yaml:"aggregateId"`
	EventType    string     `yaml:"eventType"`
	Action       string     `yaml:"action"`
	Version      string     `yaml:"version"`
	Description  string     `yaml:"description"`
	Properties   Properties `yaml:"properties"`
	DataProperty *Property
	DataFields   *Fields
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
	res := []string{}
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
	e.Name = name
	if len(e.Version) == 0 {
		e.Version = "1"
	}
	e.Properties.init(a)
	if data := e.Properties["data"]; data != nil {
		e.DataProperty = data
		if a.FieldsObjects != nil && e.DataProperty.DataType != "" {
			fields := a.FieldsObjects[e.DataProperty.DataType]
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
