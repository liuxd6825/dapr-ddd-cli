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
	buildQueryHandler        *BuildQueryHandler

	buildQueryServiceAggregate *BuildQueryServiceAggregate
	buildQueryServiceEntities  []*BuildQueryServiceEntity

	buildRepositoryAggregate *BuildRepositoryAggregate
	buildRepositoryEntities  []*BuildRepositoryEntity
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

	res.initQueryHandler()

	res.initProjectionAggregate()
	res.initProjectionEntities()

	res.initRepositoryAggregate()
	res.initRepositoryEntities()

	res.initQueryServiceAggregate()
	res.initQueryServiceEntities()

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
	buildFieldsObjects := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildFields {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildFieldsObjects()...)

	// service
	list = append(list, b.buildQueryServiceAggregate)

	// handler
	list = append(list, b.buildQueryHandler)

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

	return b.DoBuild(list...)
}

func (b *BuildDomainLayer) initFields() {
	for name, fields := range b.aggregate.FieldsObjects {
		outFile := fmt.Sprintf("%s/fields/%s_fields/%s.go", b.outDir, b.aggregate.Name, utils.SnakeString(fields.Name))
		item := NewBuildFields(b.BaseBuild, name, fields, utils.ToLower(outFile))
		b.buildFields = append(b.buildFields, item)
	}
}

func (b *BuildDomainLayer) initQueryServiceAggregate() {
	outFile := fmt.Sprintf("%s/service/%s_query_service.go", b.outDir, b.aggregate.Name)
	b.buildQueryServiceAggregate = NewBuildQueryServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initQueryHandler() {
	outFile := fmt.Sprintf("%s/handler/%s_query_handler.go", b.outDir, b.aggregate.Name)
	b.buildQueryHandler = NewBuildQueryHandler(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initProjectionEntities() {
	b.buildProjectionEntities = []*BuildProjectionEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/projection/%s_view.go", b.outDir, item.Name)
		buildEntityObject := NewBuildProjectionEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildProjectionEntities = append(b.buildProjectionEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initRepositoryEntities() {
	b.buildRepositoryEntities = []*BuildRepositoryEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/repository/%s_repository/%s_repository.go", b.outDir, b.aggregate.Name, item.Name)
		buildEntityObject := NewBuildRepositoryEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryEntities = append(b.buildRepositoryEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initQueryServiceEntities() {
	b.buildQueryServiceEntities = []*BuildQueryServiceEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/queryservice/%s_service.go", b.outDir, item.Name)
		buildEntityObject := NewBuildRepositoryEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryEntities = append(b.buildRepositoryEntities, buildEntityObject)
	}
}

func (b *BuildDomainLayer) initRepositoryAggregate() {
	outFile := fmt.Sprintf("%s/repository/%s_repository/%s_repository.go", b.outDir, b.aggregate.Name, b.aggregate.Name)
	b.buildRepositoryAggregate = NewBuildRepositoryAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildDomainLayer) initProjectionAggregate() {
	outFile := fmt.Sprintf("%s/projection/%s_view.go", b.outDir, b.aggregate.Name)
	b.buildProjectionAggregate = NewBuildProjectionAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}
