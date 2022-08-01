package query_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildDomainLayer struct {
	builds.BaseBuild
	aggregate                *config.Aggregate
	outDir                   string
	buildFields              []*BuildFields
	buildProjectionAggregate *BuildProjectionAggregate
	buildProjectionEntities  []*BuildProjectionEntity

	buildQueryServices []*BuildQueryService
	buildRepositories  []*BuildRepository

	buildFactoryAggregate *BuildFactoryAggregate
	buildFactoryEntities  []*BuildFactoryEntity
	buildEnumObjects      []*BuildEnumObject

	buildCommands []*BuildCommand
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

	res.initProjectionAggregate()
	res.initProjectionEntities()

	res.initRepository()

	res.initQueryServices()

	res.initFactoryEntities()
	res.initFactoryAggregate()

	res.initEnumObjects()
	res.initCommands()
	return res
}

func (b *BuildDomainLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildProjectionAggregate)

	// entityObject
	buildEntityObjects := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildProjectionEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildEntityObjects()...)

	// fields
	/*	buildFieldsObjects := func() []builds.Build {
			var res []builds.Build
			for _, item := range b.buildFields {
				res = append(res, item)
			}
			return res
		}
		list = append(list, buildFieldsObjects()...)*/

	// repository
	for _, item := range b.buildRepositories {
		list = append(list, item)
	}

	// query_service
	for _, item := range b.buildQueryServices {
		list = append(list, item)
	}

	// factory
	for _, item := range b.buildFactoryEntities {
		list = append(list, item)
	}
	list = append(list, b.buildFactoryAggregate)

	// enum
	for _, item := range b.buildEnumObjects {
		list = append(list, item)
	}

	// command
	for _, item := range b.buildCommands {
		list = append(list, item)
	}
	return b.DoBuild(list...)
}

func (b *BuildDomainLayer) initFields() {
	for name, field := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/%s/field/%s.go", b.outDir, b.Aggregate.FileName(), field.FileName())
		item := NewBuildFields(b.BaseBuild, name, field, utils.ToLower(outFile))
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildDomainLayer) initQueryServices() {
	b.buildQueryServices = []*BuildQueryService{}

	outFile := fmt.Sprintf("%s/%s/service/%s_query_service.go", b.outDir, b.Aggregate.FileName(), b.aggregate.FileName())
	buildQuery := NewBuildQueryService(b.BaseBuild, nil, utils.ToLower(outFile))
	b.buildQueryServices = append(b.buildQueryServices, buildQuery)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/service/%s_query_service.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildQueryService(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServices = append(b.buildQueryServices, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initRepository() {
	b.buildRepositories = []*BuildRepository{}

	outFile := fmt.Sprintf("%s/%s/repository/%s_view_repository.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	buildRepository := NewBuildRepository(b.BaseBuild, nil, utils.ToLower(outFile))
	b.buildRepositories = append(b.buildRepositories, buildRepository)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/repository/%s_view_repository.go", b.outDir, b.Aggregate.FileName(), utils.SnakeString(item.Name))
		buildRepository := NewBuildRepository(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositories = append(b.buildRepositories, buildRepository)
	}
}

func (b *BuildDomainLayer) initProjectionAggregate() {
	outFile := fmt.Sprintf("%s/%s/view/%s_view.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildProjectionAggregate = NewBuildProjectionAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initProjectionEntities() {
	b.buildProjectionEntities = []*BuildProjectionEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/view/%s_view.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildProjectionEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildProjectionEntities = append(b.buildProjectionEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initFactoryEntities() {
	b.buildFactoryEntities = []*BuildFactoryEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/factory/%s_factory.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildFactoryEntity := NewBuildFactoryEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.buildFactoryEntities = append(b.buildFactoryEntities, buildFactoryEntity)
	}
}

func (b *BuildDomainLayer) initFactoryAggregate() {
	outFile := fmt.Sprintf("%s/%s/factory/%s_factory.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildFactoryAggregate = NewBuildFactoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initEnumObjects() {
	b.buildEnumObjects = []*BuildEnumObject{}
	if b.aggregate.EnumObjects != nil {
		for _, item := range b.aggregate.EnumObjects {
			outFile := fmt.Sprintf("%s/%s/view/%s_enum.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
			buildEnumObject := NewBuildEnumObject(b.BaseBuild, item, utils.ToLower(outFile))
			b.buildEnumObjects = append(b.buildEnumObjects, buildEnumObject)
		}
	}
}

func (b *BuildDomainLayer) initCommands() {
	b.buildCommands = []*BuildCommand{}

	outFile := fmt.Sprintf("%s/%s/command/%s_query.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(b.aggregate.Name))
	buildCommand := NewBuildCommand(b.BaseBuild, b.aggregate, nil, utils.ToLower(outFile))
	b.buildCommands = append(b.buildCommands, buildCommand)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/command/%s_query.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
		buildCommand := NewBuildCommand(b.BaseBuild, nil, item, utils.ToLower(outFile))
		b.buildCommands = append(b.buildCommands, buildCommand)
	}
}
