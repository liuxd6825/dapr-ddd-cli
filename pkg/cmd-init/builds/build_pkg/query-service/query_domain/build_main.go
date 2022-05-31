package query_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
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

	buildQueryHandlerAggregate *BuildQueryHandlerAggregate
	buildQueryHandlerEntities  []*BuildQueryHandlerEntity

	buildFactoryAggregate *BuildFactoryAggregate
	buildFactoryEntities  []*BuildFactoryEntity
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

	res.initQueryHandlerAggregate()
	res.initQueryHandlerEntities()

	res.initFactoryEntities()
	res.initFactoryAggregate()

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
	buildRepositoryEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildRepositoryEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, b.buildRepositoryAggregate)
	list = append(list, buildRepositoryEntities()...)

	// query_handler
	buildQueryHandlerEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildQueryHandlerEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, b.buildQueryHandlerAggregate)
	list = append(list, buildQueryHandlerEntities()...)

	// query_service
	buildQueryServiceEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildQueryServiceEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, b.buildQueryServiceAggregate)
	list = append(list, buildQueryServiceEntities()...)

	// factory
	buildFactoryEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildFactoryEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, b.buildFactoryAggregate)
	list = append(list, buildFactoryEntities()...)

	return b.DoBuild(list...)
}

func (b *BuildDomainLayer) initFields() {
	for name, field := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/field/%s_field/%s.go", b.outDir, b.Aggregate.FileName(), field.FileName())
		item := NewBuildFields(b.BaseBuild, name, field, utils.ToLower(outFile))
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildDomainLayer) initQueryServiceAggregate() {
	outFile := fmt.Sprintf("%s/service/%s_service/%s_queryservice.go", b.outDir, b.Aggregate.FileName(), b.aggregate.FileName())
	b.buildQueryServiceAggregate = NewBuildQueryServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initQueryServiceEntities() {
	b.buildQueryServiceEntities = []*BuildQueryServiceEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/service/%s_service/%s_queryservice.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildQueryServiceEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServiceEntities = append(b.buildQueryServiceEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initEntityQueryHandlerEntities() {
	outFile := fmt.Sprintf("%s/handler/%s_handler/%s_queryhandler.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildQueryHandlerAggregate = NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initRepositoryAggregate() {
	outFile := fmt.Sprintf("%s/repository/%s_repository/%s_viewrepository.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildRepositoryAggregate = NewBuildRepositoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initRepositoryEntities() {
	b.buildRepositoryEntities = []*BuildRepositoryEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/repository/%s_repository/%s_viewrepository.go", b.outDir, b.Aggregate.FileName(), utils.SnakeString(item.Name))
		buildEntityObject := NewBuildRepositoryEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryEntities = append(b.buildRepositoryEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initProjectionAggregate() {
	outFile := fmt.Sprintf("%s/projection/%s_view/%s_view.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildProjectionAggregate = NewBuildProjectionAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initProjectionEntities() {
	b.buildProjectionEntities = []*BuildProjectionEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/projection/%s_view/%s_view.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildProjectionEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildProjectionEntities = append(b.buildProjectionEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initQueryHandlerEntities() {
	b.buildQueryHandlerEntities = []*BuildQueryHandlerEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/handler/%s_handler/%s_queryhandler.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildQueryHandlerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.buildQueryHandlerEntities = append(b.buildQueryHandlerEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initQueryHandlerAggregate() {
	outFile := fmt.Sprintf("%s/handler/%s_handler/%s_queryhandler.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildQueryHandlerAggregate = NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initFactoryEntities() {
	b.buildFactoryEntities = []*BuildFactoryEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/factory/%s_factory/%s_factory.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildFactoryEntity := NewBuildFactoryEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.buildFactoryEntities = append(b.buildFactoryEntities, buildFactoryEntity)
	}
}

func (b *BuildDomainLayer) initFactoryAggregate() {
	outFile := fmt.Sprintf("%s/factory/%s_factory/%s_factory.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildFactoryAggregate = NewBuildFactoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}
