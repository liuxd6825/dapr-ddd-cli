package cmd_infrastructure

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

	buildRegisterAllEventType  *BuildRegisterEventType
	buildRegisterAggregateType *BuildRegisterAggregateType

	buildDomainBaseCmdService   *BuildDomainBaseCmdService
	buildAppBaseQueryService    *BuildAppBaseQueryService
	buildRegisterRestController *BuildRegisterRestController
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
	res.initRegisterEventType()
	res.initRegisterAggregateType()
	res.initBuildDomainBaseCmdService()
	res.initBuildAppBaseQueryService()
	res.initRegisterRestController()
	return res
}

func (b *BuildInfrastructureLayer) Build() error {
	var list []builds.Build

	// register
	list = append(list, b.buildRegisterAllEventType)

	list = append(list, b.buildRegisterAggregateType)

	list = append(list, b.buildDomainBaseCmdService)

	list = append(list, b.buildAppBaseQueryService)

	list = append(list, b.buildRegisterRestController)

	return b.DoBuild(list...)
}

func (b *BuildInfrastructureLayer) initRegisterRestController() {
	outFile := fmt.Sprintf("%s/register/register_res_api.go", b.outDir)
	b.buildRegisterRestController = NewBuildRegisterRestController(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initRegisterEventType() {
	outFile := fmt.Sprintf("%s/register/register_event_type.go", b.outDir)
	b.buildRegisterAllEventType = NewBuildRegisterEventType(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initRegisterAggregateType() {
	outFile := fmt.Sprintf("%s/register/register_aggregate_type.go", b.outDir)
	b.buildRegisterAggregateType = NewBuildRegisterAggregateType(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initBuildDomainBaseCmdService() {
	outFile := fmt.Sprintf("%s/domain/service/base_command_service.go", b.outDir)
	b.buildDomainBaseCmdService = NewBuildDomainBaseCmdService(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildInfrastructureLayer) initBuildAppBaseQueryService() {
	outFile := fmt.Sprintf("%s/application/service/base_query_service.go", b.outDir)
	b.buildAppBaseQueryService = NewBuildAppBaseQueryService(b.BaseBuild, utils.ToLower(outFile))
}
