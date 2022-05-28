package cmd_domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildDomainLayer struct {
	builds.BaseBuild
	aggregate                        *config.Aggregate
	outDir                           string
	buildFields                      []*BuildFields
	buildCommands                    []*BuildCommand
	buildRegisterAllEventType        *BuildRegisterAllEventType
	buildRegisterAggregateEventTypes []*BuildRegisterAggregateEventType
	buildEvents                      []*BuildEvent
	buildAggregate                   *BuildAggregate
	buildDomainService               *BuildDomainService
	buildBaseDomainService           *builds.BuildAnyFile
	buildValueObjects                []*BuildValueObject
	buildEntityObjects               []*BuildEntityObject
	buildRegisterAggregateType       *BuildRegisterAggregateType
}

func NewBuildDomainLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildDomainLayer {
	res := &BuildDomainLayer{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}

	res.initFields()
	res.initCommands()
	res.initEvents()
	res.initRegisterAllEventType()
	res.initRegisterAggregateEventTypes()
	res.initModel()
	res.initDomainService()
	res.initBuildValueObjects()
	res.initBuildEntityObjects()
	res.initBuildRegisterAggregateTypes()

	return res
}

func (b *BuildDomainLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildAggregate)

	// valueObject
	buildValueObjects := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildValueObjects {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildValueObjects()...)

	// entityObject
	buildEntityObjects := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildEntityObjects {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildEntityObjects()...)

	// domainService
	list = append(list, b.buildBaseDomainService)
	list = append(list, b.buildDomainService)

	// fields
	buildFields := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildFields {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildFields()...)

	// commands
	buildCommands := func() []builds.Build {
		var res []builds.Build
		for _, b := range b.buildCommands {
			res = append(res, b)
		}
		return res
	}
	list = append(list, buildCommands()...)

	// events
	buildEvents := func() []builds.Build {
		var res []builds.Build
		for _, b := range b.buildEvents {
			res = append(res, b)
		}
		return res
	}
	buildRegisterAggregateEventTypes := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildRegisterAggregateEventTypes {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildEvents()...)
	list = append(list, b.buildRegisterAllEventType)
	list = append(list, buildRegisterAggregateEventTypes()...)
	list = append(list, b.buildRegisterAggregateType)

	return b.DoBuild(list...)
}

func (b *BuildDomainLayer) doBuildDomainService() error {
	if b.buildBaseDomainService != nil {
		if err := b.buildBaseDomainService.Build(); err != nil {
			return err
		}
	}
	// 生成领域服务
	if b.buildDomainService != nil {
		if err := b.buildDomainService.Build(); err != nil {
			return err
		}
	}
	return nil
}

func (b *BuildDomainLayer) initCommands() {
	for name, command := range b.aggregate.Commands {
		outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/command/%s_commands/%s.go", b.outDir, b.aggregate.Name, utils.SnakeString(name))
		buildCommand := NewBuildCommand(b.BaseBuild, command, utils.ToLower(outFile))
		b.buildCommands = append(b.buildCommands, buildCommand)
	}
}

func (b *BuildDomainLayer) initEvents() {
	for name, event := range b.aggregate.Events {
		outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/event/%s_events/%s.go", b.outDir, b.aggregate.Name, utils.SnakeString(name))
		item := NewBuildEvent(b.BaseBuild, name, event, utils.ToLower(outFile))
		b.buildEvents = append(b.buildEvents, item)
	}
}

func (b *BuildDomainLayer) initRegisterAggregateEventTypes() {
	regs := []*BuildRegisterAggregateEventType{}
	for name, agg := range b.Config.Aggregates {
		outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/event/%s_events/event_type.go", b.outDir, utils.SnakeString(name))
		reg := NewBuildRegisterAggregateEventType(b.BaseBuild, agg, utils.ToLower(outFile))
		regs = append(regs, reg)
	}
	b.buildRegisterAggregateEventTypes = regs
}

func (b *BuildDomainLayer) initRegisterAllEventType() {
	outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/event/reg_all_event_type.go", b.outDir)
	b.buildRegisterAllEventType = NewBuildRegisterAllEventType(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initBuildRegisterAggregateTypes() {
	outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/model/reg_aggregate_type.go", b.outDir)
	b.buildRegisterAggregateType = NewBuildRegisterAggregateType(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initFields() {
	for name, fields := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/fields/%s_fields/%s.go", b.outDir, b.aggregate.Name, utils.SnakeString(fields.Name))
		item := NewBuildFields(b.BaseBuild, name, fields, utils.ToLower(outFile))
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildDomainLayer) initModel() {
	outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/model/%s_model/%s_aggregate.go", b.outDir, b.aggregate.Name, b.aggregate.Name)
	b.buildAggregate = NewBuildAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initDomainService() {
	values := make(map[string]interface{})
	tmplFile := "static/tmpl/go/init/pkg/cmd-service/domain/service/base.go.tpl"
	outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/service/base.go", b.outDir)
	b.buildBaseDomainService = builds.NewBuildAnyFile(b.BaseBuild, values, tmplFile, utils.ToLower(outFile))

	outFile = fmt.Sprintf("%s/pkg/cmd-service/domain/service/%s_domain_service.go", b.outDir, b.aggregate.Name)
	b.buildDomainService = NewBuildDomainService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initBuildValueObjects() {
	b.buildValueObjects = []*BuildValueObject{}
	for _, item := range b.aggregate.ValueObjects {
		outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/model/%s_model/%s_value_object.go", b.outDir, b.aggregate.Name, item.Name)
		buildValueObject := NewBuildValueObject(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildValueObjects = append(b.buildValueObjects, buildValueObject)
	}
}

func (b *BuildDomainLayer) initBuildEntityObjects() {
	b.buildEntityObjects = []*BuildEntityObject{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/pkg/cmd-service/domain/model/%s_model/%s_entity.go", b.outDir, b.aggregate.Name, item.Name)
		buildEntityObject := NewBuildEntityObject(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildEntityObjects = append(b.buildEntityObjects, buildEntityObject)
	}
}
