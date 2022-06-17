package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildDomainLayer struct {
	builds.BaseBuild
	aggregate     *config.Aggregate
	outDir        string
	buildFields   []*BuildField
	buildCommands []*BuildCommand

	buildEvents        []*BuildEvent
	buildAggregate     *BuildAggregate
	buildDomainService *BuildDomainService
	buildValueObjects  []*BuildValueObject

	buildEntityObjects []*BuildEntityObject
	buildEntityLists   []*BuildEntityList

	buildEventType   *BuildEventType
	buildEnumObjects []*BuildEnumObject
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
	res.initEntityLists()

	return res
}

func (b *BuildDomainLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildAggregate)

	// valueObject
	for _, item := range b.buildValueObjects {
		list = append(list, item)
	}

	// entityObject
	for _, item := range b.buildEntityObjects {
		list = append(list, item)
	}

	// entityList
	for _, item := range b.buildEntityLists {
		list = append(list, item)
	}

	// domainService
	list = append(list, b.buildDomainService)

	// fields
	for _, item := range b.buildFields {
		list = append(list, item)
	}

	// commands
	for _, item := range b.buildCommands {
		list = append(list, item)
	}

	// event
	for _, item := range b.buildEvents {
		list = append(list, item)
	}

	// enumObject
	for _, item := range b.buildEnumObjects {
		list = append(list, item)
	}

	// event type
	list = append(list, b.buildEventType)

	return b.DoBuild(list...)
}

func (b *BuildDomainLayer) doBuildDomainService() error {
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
		outFile := fmt.Sprintf("%s/%s/command/%s.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(name))
		buildCommand := NewBuildCommand(b.BaseBuild, command, utils.ToLower(outFile))
		b.buildCommands = append(b.buildCommands, buildCommand)
	}
}

func (b *BuildDomainLayer) initEvents() {
	for name, event := range b.aggregate.Events {
		outFile := fmt.Sprintf("%s/%s/event/%s.go", b.outDir, b.aggregate.FileName(), event.SnakeName())
		item := NewBuildEvent(b.BaseBuild, name, event, utils.ToLower(outFile))
		b.buildEvents = append(b.buildEvents, item)
	}
}

func (b *BuildDomainLayer) initFields() {
	for name, field := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/%s/field/%s.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(field.Name))
		item := NewBuildField(b.BaseBuild, name, field, utils.ToLower(outFile))
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildDomainLayer) initModel() {
	outFile := fmt.Sprintf("%s/%s/model/%s_aggregate.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildAggregate = NewBuildAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initDomainService() {
	outFile := fmt.Sprintf("%s/%s/service/%s_domain_service.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildDomainService = NewBuildDomainService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initBuildValueObjects() {
	b.buildValueObjects = []*BuildValueObject{}
	for _, item := range b.aggregate.ValueObjects {
		outFile := fmt.Sprintf("%s/%s/model/%s_value_object.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
		buildValueObject := NewBuildValueObject(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildValueObjects = append(b.buildValueObjects, buildValueObject)
	}
}

func (b *BuildDomainLayer) initBuildEntityObjects() {
	b.buildEntityObjects = []*BuildEntityObject{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/model/%s_entity.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
		buildEntityObject := NewBuildEntityObject(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildEntityObjects = append(b.buildEntityObjects, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initEnumObjects() {
	b.buildEnumObjects = []*BuildEnumObject{}
	if b.aggregate.EnumObjects != nil {
		for _, item := range b.aggregate.EnumObjects {
			outFile := fmt.Sprintf("%s/%s/model/%s_enum.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
			buildEnumObject := NewBuildEnumObject(b.BaseBuild, item, utils.ToLower(outFile))
			b.buildEnumObjects = append(b.buildEnumObjects, buildEnumObject)
		}
	}
}

func (b *BuildDomainLayer) initEventTypes() {
	outFile := fmt.Sprintf("%s/%s/event/event_type.go", b.outDir, b.aggregate.FileName())
	b.buildEventType = NewBuildEventType(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

}

func (b *BuildDomainLayer) initEntityLists() {
	b.buildEntityLists = []*BuildEntityList{}
	for _, item := range b.aggregate.Entities {
		fileName := utils.SnakeString(utils.Plural(item.Name))
		outFile := fmt.Sprintf("%s/%s/model/%s.go", b.outDir, b.aggregate.FileName(), fileName)
		buildEntityList := NewBuildEntityList(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildEntityLists = append(b.buildEntityLists, buildEntityList)
	}
}
