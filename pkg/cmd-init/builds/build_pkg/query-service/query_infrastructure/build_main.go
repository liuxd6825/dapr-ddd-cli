package query_infrastructure

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildInfrastructureLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string

	buildRepositoryImplAggregate *BuildRepositoryImplAggregate
	buildRepositoryImplEntities  []*BuildRepositoryImplEntity

	buildQueryServiceImplAggregate *BuildQueryServiceImplAggregate
	buildQueryServiceImplEntities  []*BuildQueryServiceImplEntity

	buildRegisterAllEventType   *BuildRegisterEventType
	buildRegisterAggregateType  *BuildRegisterAggregateType
	buildRegisterRestController *BuildRegisterRestApi

	buildRepositoryBase    *BuildRepositoryBase
	buildRegisterSubscribe *BuildRegisterSubscribe

	buildDtoBase       *builds.BuildAnyFile
	buildTypesDateTime *BuildTypesDateTime
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
	res.initRegisterSubscribe()

	res.initRegisterEventType()
	res.initRegisterAggregateType()
	res.initRegisterRestController()
	res.initDtoBase()
	res.initTypes()

	return res
}

func (b *BuildInfrastructureLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildRepositoryImplAggregate)

	// registerController
	list = append(list, b.buildRegisterRestController)

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
	list = append(list, b.buildRegisterSubscribe)
	list = append(list, b.buildRegisterAggregateType)
	list = append(list, b.buildRegisterAllEventType)
	list = append(list, b.buildDtoBase)
	list = append(list, b.buildTypesDateTime)

	return b.DoBuild(list...)
}

func (b *BuildInfrastructureLayer) initRepositoryEntities() {
	b.buildRepositoryImplEntities = []*BuildRepositoryImplEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/%s/repository/mongodb/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRepositoryImplEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryImplEntities = append(b.buildRepositoryImplEntities, build)
	}
}

func (b *BuildInfrastructureLayer) initRepositoryAggregate() {
	outFile := fmt.Sprintf("%s/domain/%s/repository/mongodb/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildRepositoryImplAggregate = NewBuildRepositoryImplAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initQueryServiceEntities() {
	b.buildQueryServiceImplEntities = []*BuildQueryServiceImplEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/%s/service/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildQueryServiceImplEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServiceImplEntities = append(b.buildQueryServiceImplEntities, build)
	}
}

func (b *BuildInfrastructureLayer) initQueryServiceAggregate() {
	outFile := fmt.Sprintf("%s/domain/%s/service/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildQueryServiceImplAggregate = NewBuildQueryServiceImplAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initBuildRepositoryBase() {
	outFile := fmt.Sprintf("%s/base/domain/repository/mongodb_base/base_repository.go", b.outDir)
	b.buildRepositoryBase = NewBuildRepositoryBase(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initRegisterSubscribe() {
	outFile := fmt.Sprintf("%s/register/register_subscribe.go", b.outDir)
	b.buildRegisterSubscribe = NewBuildRegisterSubscribe(b.BaseBuild, outFile)
}

func (b *BuildInfrastructureLayer) initRegisterEventType() {
	outFile := fmt.Sprintf("%s/register/register_event_type.go", b.outDir)
	b.buildRegisterAllEventType = NewBuildRegisterEventType(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initRegisterAggregateType() {
	outFile := fmt.Sprintf("%s/register/register_aggregate_type.go", b.outDir)
	b.buildRegisterAggregateType = NewBuildRegisterAggregateType(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initRegisterRestController() {
	outFile := fmt.Sprintf("%s/register/register_rest_api.go", b.outDir)
	b.buildRegisterRestController = NewBuildRegisterRestApi(b.BaseBuild, outFile)
}

func (b *BuildInfrastructureLayer) initDtoBase() {
	values := b.BaseBuild.Values()
	outFile := fmt.Sprintf("%s/base/userinterface/rest/dto/base_dto.go", b.outDir)
	tmplFile := "static/tmpl/go/init/pkg/query-service/infrastructure/base/userinterface/rest/dto/base_dto.go.tpl"
	b.buildDtoBase = builds.NewBuildAnyFile(b.BaseBuild, values, tmplFile, outFile)
}

func (b *BuildInfrastructureLayer) initTypes() {
	outFile := fmt.Sprintf("%s/types/date_time.go", b.outDir)
	b.buildTypesDateTime = NewBuildTypesDateTime(b.BaseBuild, outFile)
}
