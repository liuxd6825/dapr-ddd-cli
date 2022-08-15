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
	database := cfg.Configuration.Database
	if database.Mongo {
		res.initRepositoryImpl("mongodb")
	}
	if database.Neo4j {
		res.initRepositoryImpl("neo4j")
	}
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
	res.initDb(database)
	res.initBaseQueryHandler()
	return res
}

func (b *BuildInfrastructureLayer) initDb(database config.Database) {
	if database.Mongo {
		mongoOutFile := fmt.Sprintf("%s/db/dao/mongo_dao/dao.go", b.outDir)
		buildMongoDao := NewBuildDbDao(b.BaseBuild, b.aggregate, mongoOutFile, "mongo_dao")
		b.AddBuild(buildMongoDao)
	}

	if database.Neo4j {
		neo4jOutFile := fmt.Sprintf("%s/db/dao/neo4j_dao/dao.go", b.outDir)
		buildNeo4jDao := NewBuildDbDao(b.BaseBuild, b.aggregate, neo4jOutFile, "neo4j_dao")
		b.AddBuild(buildNeo4jDao)
	}

	if database.HaveDb() {
		sessionFile := fmt.Sprintf("%s/db/session/session.go", b.outDir)
		buildSession := NewBuildDbSession(b.BaseBuild, sessionFile)
		b.AddBuild(buildSession)
	}

}

func (b *BuildInfrastructureLayer) initRepositoryImpl(dbType string) {
	outFile := fmt.Sprintf("%s/domain_impl/%s/repository_impl/%s/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), dbType, b.aggregate.FileName())
	buildRepository := NewBuildRepositoryImpl(b.BaseBuild, nil, dbType, utils.ToLower(outFile))
	b.AddBuild(buildRepository)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain_impl/%s/repository_impl/%s/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), dbType, item.FileName())
		build := NewBuildRepositoryImpl(b.BaseBuild, item, dbType, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildInfrastructureLayer) initQueryService() {
	outFile := fmt.Sprintf("%s/domain_impl/%s/service_impl/x_options.go", b.outDir, b.aggregate.FileName())
	buildOptions := NewBuildQueryServiceOptionsImpl(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(buildOptions)

	outFile = fmt.Sprintf("%s/domain_impl/%s/service_impl/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildService := NewBuildQueryServiceImpl(b.BaseBuild, nil, utils.ToLower(outFile))
	b.AddBuild(buildService)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain_impl/%s/service_impl/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildQueryServiceImpl(b.BaseBuild, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}
}

func (b *BuildInfrastructureLayer) initBaseQueryHandler() {
	outFile := fmt.Sprintf("%s/base/application/handler/base_query_handler.go", b.outDir)
	buildBaseQueryHandler := NewBuildBaseQueryHandler(b.BaseBuild, outFile)
	b.AddBuild(buildBaseQueryHandler)
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
