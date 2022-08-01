package query_infrastructure

import (
	"fmt"
	builds2 "github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildInfrastructureLayer struct {
	builds2.BaseBuild
	aggregate *config.Aggregate
	outDir    string

	buildRepositoryImpls   []*BuildRepositoryImpl
	buildQueryServiceImpls []*BuildQueryServiceImpl

	buildRegisterAllEventType   *BuildRegisterEventType
	buildRegisterAggregateType  *BuildRegisterAggregateType
	buildRegisterRestController *BuildRegisterRestApi

	buildRepositoryBase    *BuildBaseRepository
	buildRegisterSubscribe *BuildRegisterSubscribe

	buildDtoBase *builds2.BuildAnyFile
	buildTypes   *BuildTypes

	buildBaseApi       *BuildBaseApi
	buildBaseAssembler *BuildBaseAssembler
	buildUtils         *BuildUtils
	buildBaseView      *BuildBaseView
	buildBaseDto       *BuildBaseDto
}

func NewBuildInfrastructureLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildInfrastructureLayer {
	res := &BuildInfrastructureLayer{
		BaseBuild: builds2.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}

	res.initRepository()
	res.initQueryService()

	res.initBuildRepositoryBase()
	res.initRegisterSubscribe()

	res.initRegisterEventType()
	res.initRegisterAggregateType()
	res.initRegisterRestController()
	res.initDtoBase()
	res.initTypes()
	res.initUtils()
	res.initBaseApi()
	res.initBaseAssembler()
	res.initBaseView()
	res.initBaseDto()
	return res
}

func (b *BuildInfrastructureLayer) Build() error {
	var list []builds2.Build

	// RegisterRestController
	list = append(list, b.buildRegisterRestController)

	// RepositoryImpl
	for _, item := range b.buildRepositoryImpls {
		list = append(list, item)
	}

	// QueryService
	for _, item := range b.buildQueryServiceImpls {
		list = append(list, item)
	}

	list = append(list, b.buildRepositoryBase)
	list = append(list, b.buildRegisterSubscribe)
	list = append(list, b.buildRegisterAggregateType)
	list = append(list, b.buildRegisterAllEventType)
	list = append(list, b.buildDtoBase)
	list = append(list, b.buildTypes)
	list = append(list, b.buildUtils)
	list = append(list, b.buildBaseApi)
	list = append(list, b.buildBaseAssembler)
	list = append(list, b.buildBaseView)
	list = append(list, b.buildBaseDto)

	return b.DoBuild(list...)
}

func (b *BuildInfrastructureLayer) initRepository() {
	b.buildRepositoryImpls = []*BuildRepositoryImpl{}
	outFile := fmt.Sprintf("%s/domain/%s/repository_impl/mongodb/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildRepository := NewBuildRepositoryImpl(b.BaseBuild, nil, utils.ToLower(outFile))
	b.buildRepositoryImpls = append(b.buildRepositoryImpls, buildRepository)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/%s/repository_impl/mongodb/%s_view_repository_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRepositoryImpl(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRepositoryImpls = append(b.buildRepositoryImpls, build)
	}
}

func (b *BuildInfrastructureLayer) initQueryService() {
	b.buildQueryServiceImpls = []*BuildQueryServiceImpl{}

	outFile := fmt.Sprintf("%s/domain/%s/service_impl/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	build := NewBuildQueryServiceImpl(b.BaseBuild, nil, utils.ToLower(outFile))
	b.buildQueryServiceImpls = append(b.buildQueryServiceImpls, build)

	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/domain/%s/service_impl/%s_query_service_impl.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildQueryServiceImpl(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildQueryServiceImpls = append(b.buildQueryServiceImpls, build)
	}
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
	b.buildDtoBase = builds2.NewBuildAnyFile(b.BaseBuild, values, tmplFile, outFile)
}

func (b *BuildInfrastructureLayer) initTypes() {
	outFile := fmt.Sprintf("%s/types/types.go", b.outDir)
	b.buildTypes = NewBuildTypes(b.BaseBuild, outFile)
}

func (b *BuildInfrastructureLayer) initUtils() {
	outFile := fmt.Sprintf("%s/utils/utils.go", b.outDir)
	b.buildUtils = NewBuildUtils(b.BaseBuild, outFile)
}

func (b *BuildInfrastructureLayer) initBaseApi() {
	outFile := fmt.Sprintf("%s/base/userinterface/rest/facade/base_api.go", b.outDir)
	b.buildBaseApi = NewBuildBaseApi(b.BaseBuild, outFile)
}

func (b *BuildInfrastructureLayer) initBaseAssembler() {
	outFile := fmt.Sprintf("%s/base/userinterface/rest/assembler/base_assembler.go", b.outDir)
	b.buildBaseAssembler = NewBuildBaseAssembler(b.BaseBuild, outFile)
}

func (b *BuildInfrastructureLayer) initBaseView() {
	outFile := fmt.Sprintf("%s/base/domain/view/base_view.go", b.outDir)
	b.buildBaseView = NewBuildBaseView(b.BaseBuild, b.aggregate, outFile)
}

func (b *BuildInfrastructureLayer) initBaseDto() {
	outFile := fmt.Sprintf("%s/base/userinterface/rest/dto/base_dto.go", b.outDir)
	b.buildBaseDto = NewBuildBaseDto(b.BaseBuild, b.aggregate, outFile)
}
