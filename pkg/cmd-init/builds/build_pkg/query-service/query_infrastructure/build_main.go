package query_infrastructure

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildInfrastructureLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string

	buildRepositoryImplAggregate *BuildRepositoryImplAggregate
	buildRepositoryImplEntities  []*BuildRepositoryImplEntity

	buildQueryServiceImplAggregate *BuildQueryServiceImplAggregate
	buildQueryServiceImplEntities  []*BuildQueryServiceImplEntity

	buildRepositoryBase *BuildRepositoryBase
}

func NewBuildInfrastructureLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildInfrastructureLayer {
	res := &BuildInfrastructureLayer{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}

	res.initRepositoryAggregate()
	res.initRepositoryEntities()

	res.initQueryServiceAggregate()
	res.initQueryServiceEntities()

	res.initBuildRepositoryBase()
	return res
}

func (b *BuildInfrastructureLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildRepositoryImplAggregate)

	// entityObject
	buildRepositoryImplEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildRepositoryImplEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildRepositoryImplEntities()...)

	// aggregate
	list = append(list, b.buildQueryServiceImplAggregate)

	// entityObject
	buildQueryServiceImplEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildQueryServiceImplEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildQueryServiceImplEntities()...)

	list = append(list, b.buildRepositoryBase)

	return b.DoBuild(list...)
}

func (b *BuildInfrastructureLayer) initRepositoryEntities() {
	b.buildRepositoryImplEntities = []*BuildRepositoryImplEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/repository/mongodb/%s_repository/%s_viewrepository_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRepositoryImplEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryImplEntities = append(b.buildRepositoryImplEntities, build)
	}
}

func (b *BuildInfrastructureLayer) initRepositoryAggregate() {
	outFile := fmt.Sprintf("%s/domain/repository/mongodb/%s_repository/%s_viewrepository_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildRepositoryImplAggregate = NewBuildRepositoryImplAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initQueryServiceEntities() {
	b.buildQueryServiceImplEntities = []*BuildQueryServiceImplEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/service/%s_service/%s_queryservice_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildQueryServiceImplEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServiceImplEntities = append(b.buildQueryServiceImplEntities, build)
	}
}

func (b *BuildInfrastructureLayer) initQueryServiceAggregate() {
	outFile := fmt.Sprintf("%s/domain/service/%s_service/%s_queryservice_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildQueryServiceImplAggregate = NewBuildQueryServiceImplAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initBuildRepositoryBase() {
	outFile := fmt.Sprintf("%s/domain/repository/mongodb/base_repository.go", b.outDir)
	b.buildRepositoryBase = NewBuildRepositoryBase(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}
