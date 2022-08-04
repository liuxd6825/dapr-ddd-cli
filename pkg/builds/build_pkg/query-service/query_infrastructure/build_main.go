package query_infrastructure

import (
	"fmt"
	builds "github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildInfrastructureLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
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

	res.initRepository()
	res.initQueryService()
	res.initRegisterSubscribe()
	res.initRegisterEventType()
	res.initRegisterAggregateType()
	res.initRegisterRestController()
	res.initTypes()
	res.initUtils()
	res.initBaseApi()
	res.initBaseAssembler()
	res.initBaseView()
	res.initBaseDto()
	res.initBaseDao()

	return res
}

func (b *BuildInfrastructureLayer) initBaseDao() {
	mongoOutFile := fmt.Sprintf("%s/base/domain/dao/mongo_dao/dao.go", b.outDir)
	buildMongoDao := NewBuildBaseDao(b.BaseBuild, b.aggregate, mongoOutFile, "mongo_dao")
	b.AddBuild(buildMongoDao)

	neo4jOutFile := fmt.Sprintf("%s/base/domain/dao/neo4j_dao/dao.go", b.outDir)
	buildNeo4jDao := NewBuildBaseDao(b.BaseBuild, b.aggregate, neo4jOutFile, "neo4j_dao")
	b.AddBuild(buildNeo4jDao)
}

func (b *BuildInfrastructureLayer) initRepository() {
	outFile := fmt.Sprintf("%s/domain/%s/repository_impl/mongodb/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildRepository := NewBuildRepositoryImpl(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(buildRepository)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/%s/repository_impl/mongodb/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRepositoryImpl(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildInfrastructureLayer) initQueryService() {
	outFile := fmt.Sprintf("%s/domain/%s/service_impl/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildQueryServiceImpl(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(build)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/%s/service_impl/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildQueryServiceImpl(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildInfrastructureLayer) initRegisterSubscribe() {
	outFile := fmt.Sprintf("%s/register/register_subscribe.go", b.outDir)
	buildRegisterSubscribe := NewBuildRegisterSubscribe(b.BaseBuild, outFile)
	b.AddBuild(buildRegisterSubscribe)
}

func (b *BuildInfrastructureLayer) initRegisterEventType() {
	outFile := fmt.Sprintf("%s/register/register_event_type.go", b.outDir)
	buildRegisterAllEventType := NewBuildRegisterEventType(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(buildRegisterAllEventType)
}

func (b *BuildInfrastructureLayer) initRegisterAggregateType() {
	outFile := fmt.Sprintf("%s/register/register_aggregate_type.go", b.outDir)
	buildRegisterAggregateType := NewBuildRegisterAggregateType(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(buildRegisterAggregateType)
}

func (b *BuildInfrastructureLayer) initRegisterRestController() {
	outFile := fmt.Sprintf("%s/register/register_rest_api.go", b.outDir)
	buildRegisterRestController := NewBuildRegisterRestApi(b.BaseBuild, outFile)
	b.AddBuild(buildRegisterRestController)
}

func (b *BuildInfrastructureLayer) initTypes() {
	outFile := fmt.Sprintf("%s/types/types.go", b.outDir)
	buildTypes := NewBuildTypes(b.BaseBuild, outFile)
	b.AddBuild(buildTypes)
}

func (b *BuildInfrastructureLayer) initUtils() {
	outFile := fmt.Sprintf("%s/utils/utils.go", b.outDir)
	buildUtils := NewBuildUtils(b.BaseBuild, outFile)
	b.AddBuild(buildUtils)
}

func (b *BuildInfrastructureLayer) initBaseApi() {
	outFile := fmt.Sprintf("%s/base/userinterface/rest/facade/base_api.go", b.outDir)
	buildBaseApi := NewBuildBaseApi(b.BaseBuild, outFile)
	b.AddBuild(buildBaseApi)
}

func (b *BuildInfrastructureLayer) initBaseAssembler() {
	outFile := fmt.Sprintf("%s/base/userinterface/rest/assembler/base_assembler.go", b.outDir)
	buildBaseAssembler := NewBuildBaseAssembler(b.BaseBuild, outFile)
	b.AddBuild(buildBaseAssembler)
}

func (b *BuildInfrastructureLayer) initBaseView() {
	outFile := fmt.Sprintf("%s/base/domain/view/base_view.go", b.outDir)
	buildBaseView := NewBuildBaseView(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildBaseView)
}

func (b *BuildInfrastructureLayer) initBaseDto() {
	outFile := fmt.Sprintf("%s/base/userinterface/rest/dto/base_dto.go", b.outDir)
	buildBaseDto := NewBuildBaseDto(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildBaseDto)
}
