package cmd_service

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildLayerDomain struct {
	builds.BaseBuild
	aggregate                        *config.Aggregate
	dir                              string
	saveFile                         string
	buildFields                      []*BuildFields
	buildCommands                    []*BuildCommand
	buildRegisterAllEventType        *BuildRegisterEventType
	buildRegisterAggregateEventTypes *[]BuildRegisterEventType
	buildEvents                      []*BuildEvent
	buildAggregate                   *BuildAggregate
	buildDomainService               *BuildDomainService
}

func NewBuildLayerDomain(cfg *config.Config, aggregate *config.Aggregate, dir string) *BuildLayerDomain {
	res := &BuildLayerDomain{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		dir:       dir,
	}

	res.initFields()
	res.initCommands()
	res.initEvents()
	res.initRegisterAllEventType()
	res.initRegisterAggregateEventTypes()
	res.initModel()
	res.initDomainService()

	return res
}

func (b *BuildLayerDomain) Build() error {
	// 生成领域服务
	if b.buildDomainService != nil {
		b.buildDomainService.Build()
	}

	return nil
	
	//  生成聚合根
	if b.buildAggregate != nil {
		b.buildAggregate.Build()
	}

	// 生成Fields
	if b.buildFields != nil {
		for _, build := range b.buildFields {
			build.Build()
		}
	}

	// 生成命令
	if b.buildCommands != nil {
		for _, build := range b.buildCommands {
			build.Build()
		}
	}

	// 生成事件
	if b.buildEvents != nil {
		for _, build := range b.buildEvents {
			build.Build()
		}
	}

	// 生成全部事件注册器
	if b.buildRegisterAllEventType != nil {
		b.buildRegisterAllEventType.Build()
	}

	// 生成聚合根的事件注册器
	if b.buildRegisterAggregateEventTypes != nil {
		for _, build := range *b.buildRegisterAggregateEventTypes {
			build.Build()
		}
	}

	return nil
}

func (b *BuildLayerDomain) initCommands() {
	// commandDir := fmt.Sprintf("%s/command/%s_commands", b.dir, b.name)
	for _, command := range b.aggregate.Commands {
		buildCommand := NewBuildCommand(b.BaseBuild, command)
		b.buildCommands = append(b.buildCommands, buildCommand)
	}
}

func (b *BuildLayerDomain) initEvents() {
	dir := fmt.Sprintf("%s/event/%s_events", b.dir, b.aggregate.Name)
	for name, event := range b.aggregate.Events {
		item := NewBuildEvent(b.BaseBuild, name, event, dir)
		b.buildEvents = append(b.buildEvents, item)
	}

}

func (b *BuildLayerDomain) initRegisterAggregateEventTypes() {
	regs := []BuildRegisterEventType{}
	for name, agg := range b.Config.Aggregates {
		dir := fmt.Sprintf("%s/event/%s_events/event_type", b.dir, name)
		reg := NewBuildRegisterAggregateEventType(b.BaseBuild, agg, dir)
		regs = append(regs, *reg)
	}
	b.buildRegisterAggregateEventTypes = &regs
}

func (b *BuildLayerDomain) initRegisterAllEventType() {
	dir := fmt.Sprintf("%s/event/reg_all_event_type", b.dir)
	b.buildRegisterAllEventType = NewBuildRegisterAllEventType(b.BaseBuild, b.aggregate, dir)
}

func (b *BuildLayerDomain) initFields() {
	dir := fmt.Sprintf("%s/fields/%s_fields", b.dir, b.aggregate.Name)
	for name, fields := range b.aggregate.FieldsObjects {
		item := NewBuildFields(b.BaseBuild, name, fields, dir)
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildLayerDomain) initModel() {
	b.buildAggregate = NewBuildAggregate(b.BaseBuild, b.aggregate)
}

func (b *BuildLayerDomain) initDomainService() {
	b.buildDomainService = NewBuildDomainService(b.BaseBuild, b.aggregate)
}
