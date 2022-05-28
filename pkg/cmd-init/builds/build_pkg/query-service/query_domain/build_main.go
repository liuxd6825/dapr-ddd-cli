package query_domain

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
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
	outFile := fmt.Sprintf("%s/queryservice/%s_queryservice/%s_query_service.go", b.outDir, b.Aggregate.FileName(), b.aggregate.FileName())
	b.buildQueryServiceAggregate = NewBuildQueryServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initQueryServiceEntities() {
	b.buildQueryServiceEntities = []*BuildQueryServiceEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/queryservice/%s_queryservice/%s_query_service.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildQueryServiceEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServiceEntities = append(b.buildQueryServiceEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initEntityQueryHandlerEntities() {
	outFile := fmt.Sprintf("%s/queryhandler/%s_queryhandler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildQueryHandlerAggregate = NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initProjectionEntities() {
	b.buildProjectionEntities = []*BuildProjectionEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/projection/%s_view.go", b.outDir, item.FileName())
		buildEntityObject := NewBuildProjectionEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildProjectionEntities = append(b.buildProjectionEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initRepositoryEntities() {
	b.buildRepositoryEntities = []*BuildRepositoryEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/repository/%s_repository/%s_repository.go", b.outDir, b.Aggregate.FileName(), utils.SnakeString(item.Name))
		buildEntityObject := NewBuildRepositoryEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryEntities = append(b.buildRepositoryEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initRepositoryAggregate() {
	outFile := fmt.Sprintf("%s/repository/%s_repository/%s_repository.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildRepositoryAggregate = NewBuildRepositoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initProjectionAggregate() {
	outFile := fmt.Sprintf("%s/projection/%s_view.go", b.outDir, b.Aggregate.FileName())
	b.buildProjectionAggregate = NewBuildProjectionAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initQueryHandlerEntities() {
	b.buildQueryHandlerEntities = []*BuildQueryHandlerEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/queryhandler/%s_queryhandler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), item.FileName())
		buildEntityObject := NewBuildQueryHandlerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.buildQueryHandlerEntities = append(b.buildQueryHandlerEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initQueryHandlerAggregate() {
	outFile := fmt.Sprintf("%s/queryhandler/%s_queryhandler/%s_query_handler.go", b.outDir, b.Aggregate.FileName(), b.Aggregate.FileName())
	b.buildQueryHandlerAggregate = NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}
