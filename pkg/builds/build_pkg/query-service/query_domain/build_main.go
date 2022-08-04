package query_domain

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

func (b *BuildDomainLayer) initFields() {
	for name, field := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/%s/field/%s.go", b.outDir, b.Aggregate.FileName(), field.FileName())
		build := NewBuildFields(b.BaseBuild, name, field, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initQueryServices() {

	outFile := fmt.Sprintf("%s/%s/service/%s_query_service.go", b.outDir, b.Aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildQueryService(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/service/%s_query_service.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		build := NewBuildQueryService(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initRepository() {
	outFile := fmt.Sprintf("%s/%s/repository/%s_view_repository.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	build := NewBuildRepository(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/repository/%s_view_repository.go", b.outDir, b.Aggregate.FileName(), utils.SnakeString(item.Name))
		build := NewBuildRepository(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initProjectionAggregate() {
	outFile := fmt.Sprintf("%s/%s/view/%s_view.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	build := NewBuildProjectionAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildDomainLayer) initProjectionEntities() {
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/view/%s_view.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		build := NewBuildProjectionEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initFactoryEntities() {
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/factory/%s_factory.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		build := NewBuildFactoryEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildDomainLayer) initFactoryAggregate() {
	outFile := fmt.Sprintf("%s/%s/factory/%s_factory.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	build := NewBuildFactoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildDomainLayer) initEnumObjects() {
	if b.aggregate.EnumObjects != nil {
		for _, item := range b.aggregate.EnumObjects {
			outFile := fmt.Sprintf("%s/%s/view/%s_enum.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
			build := NewBuildEnumObject(b.BaseBuild, item, utils.ToLower(outFile))
			b.AddBuild(build)
		}
	}
}

func (b *BuildDomainLayer) initCommands() {
	outFile := fmt.Sprintf("%s/%s/command/%s_query.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(b.aggregate.Name))
	build := NewBuildCommand(b.BaseBuild, b.aggregate, nil, utils.ToLower(outFile))
	b.AddBuild(build)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/command/%s_query.go", b.outDir, b.aggregate.FileName(), utils.SnakeString(item.Name))
		build := NewBuildCommand(b.BaseBuild, nil, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}
