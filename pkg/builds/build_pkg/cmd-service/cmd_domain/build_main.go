package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildDomainLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
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
	res.initModel()
	res.initDomainService()
	res.initBuildValueObjects()
	res.initBuildEntityObjects()
	res.initEnumObjects()
	res.initEventTypes()
	res.initEntityItems()
	res.initEventFactory()
	return res
}

func (b *BuildDomainLayer) initEventFactory() {
	outFile := fmt.Sprintf("%s/%s/factory/event_factory.go", b.outDir, b.aggregate.FileName())
	build := NewBuildEventFactory(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildDomainLayer) initCommands() {
	for name, command := range b.aggregate.Commands {
		outFile := fmt.Sprintf("%s/%s/command/%s.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(name))
		build := NewBuildCommand(b.BaseBuild, command, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initEvents() {
	for name, event := range b.aggregate.Events {
		outFile := fmt.Sprintf("%s/%s/event/%s.go", b.outDir, b.aggregate.FileName(), event.SnakeName())
		build := NewBuildEvent(b.BaseBuild, name, event, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initFields() {
	for name, field := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/%s/field/%s.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(field.Name))
		build := NewBuildField(b.BaseBuild, name, field, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initModel() {
	outFile := fmt.Sprintf("%s/%s/model/%s_aggregate.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildAggregate := NewBuildAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildAggregate)

	outFile = fmt.Sprintf("%s/%s/model/%s_aggregate_event.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildAggregateEvent := NewBuildAggregateEvent(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildAggregateEvent)

	outFile = fmt.Sprintf("%s/%s/model/%s_aggregate_command.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildAggregateCommand := NewBuildAggregateCommand(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildAggregateCommand)
}

func (b *BuildDomainLayer) initDomainService() {
	outFile := fmt.Sprintf("%s/%s/service/%s_domain_service.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildDomainService := NewBuildDomainService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildDomainService)
}

func (b *BuildDomainLayer) initBuildValueObjects() {
	for _, item := range b.aggregate.ValueObjects {
		outFile := fmt.Sprintf("%s/%s/model/%s_value_object.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
		buildValueObject := NewBuildValueObject(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(buildValueObject)
	}
}

func (b *BuildDomainLayer) initBuildEntityObjects() {
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/model/%s_entity.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
		buildEntityObject := NewBuildEntityObject(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(buildEntityObject)
	}
}

func (b *BuildDomainLayer) initEnumObjects() {
	if b.aggregate.EnumObjects != nil {
		for _, item := range b.aggregate.EnumObjects {
			outFile := fmt.Sprintf("%s/%s/field/%s_enum.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
			buildEnumObject := NewBuildEnumObject(b.BaseBuild, item, utils.ToLower(outFile))
			b.AddBuild(buildEnumObject)
		}
	}
}

func (b *BuildDomainLayer) initEventTypes() {
	outFile := fmt.Sprintf("%s/%s/event/event_type.go", b.outDir, b.aggregate.FileName())
	buildEventType := NewBuildEventType(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildEventType)
}

func (b *BuildDomainLayer) initEntityItems() {
	for _, item := range b.aggregate.Entities {
		fileName := utils.SnakeString(item.Name)
		outFile := fmt.Sprintf("%s/%s/model/%s_items.go", b.outDir, b.aggregate.FileName(), fileName)
		buildEntityList := NewBuildEntityItems(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(buildEntityList)
	}
}
