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

	buildQueryServiceAggregate *BuildQueryServiceAggregate
	buildQueryServiceEntities  []*BuildQueryServiceEntity

	buildRepositoryAggregate *BuildRepositoryAggregate
	buildRepositoryEntities  []*BuildRepositoryEntity

	buildFactoryAggregate *BuildFactoryAggregate
	buildFactoryEntities  []*BuildFactoryEntity
	buildEnumObjects      []*BuildEnumObject
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

	res.initRepositoryAggregate()
	res.initRepositoryEntities()

	res.initQueryServiceAggregate()
	res.initQueryServiceEntities()

	res.initFactoryEntities()
	res.initFactoryAggregate()

	res.initEnumObjects()
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
	for _, item := range b.buildRepositoryEntities {
		list = append(list, item)
	}
	list = append(list, b.buildRepositoryAggregate)

	// query_service
	for _, item := range b.buildQueryServiceEntities {
		list = append(list, item)
	}
	list = append(list, b.buildQueryServiceAggregate)

	// factory
	for _, item := range b.buildFactoryEntities {
		list = append(list, item)
	}
	list = append(list, b.buildFactoryAggregate)

	// enum
	for _, item := range b.buildEnumObjects {
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

func (b *BuildDomainLayer) initQueryServiceAggregate() {
	outFile := fmt.Sprintf("%s/%s/service/%s_query_service.go", b.outDir, b.Aggregate.FileName(), b.aggregate.FileName())
	b.buildQueryServiceAggregate = NewBuildQueryServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initQueryServiceEntities() {
	b.buildQueryServiceEntities = []*BuildQueryServiceEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/service/%s_query_service.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildQueryServiceEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServiceEntities = append(b.buildQueryServiceEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initRepositoryAggregate() {
	outFile := fmt.Sprintf("%s/%s/repository/%s_view_repository.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildRepositoryAggregate = NewBuildRepositoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initRepositoryEntities() {
	b.buildRepositoryEntities = []*BuildRepositoryEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/%s/repository/%s_view_repository.go", b.outDir, b.Aggregate.FileName(), utils.SnakeString(item.Name))
		buildEntityObject := NewBuildRepositoryEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryEntities = append(b.buildRepositoryEntities, buildEntityObject)
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
