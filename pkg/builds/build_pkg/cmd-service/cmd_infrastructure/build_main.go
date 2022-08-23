package cmd_infrastructure

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
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
	res.initRegisterEventType()
	res.initRegisterAggregateType()
	res.initBuildDomainBaseCmdService()
	res.initBuildAppBaseQueryService()
	res.initRegisterRestController()
	res.initBuildUtilsAssembler()
	res.initBaseEvent()
	res.initLogs()
	return res
}

func (b *BuildInfrastructureLayer) initLogs() {
	outFile := fmt.Sprintf("%s/logs/logs.go", b.outDir)
	build := NewBuildLogs(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initRegisterRestController() {
	outFile := fmt.Sprintf("%s/register/register_res_api.go", b.outDir)
	build := NewBuildRegisterRestController(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initRegisterEventType() {
	outFile := fmt.Sprintf("%s/register/register_event_type.go", b.outDir)
	build := NewBuildRegisterEventType(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initRegisterAggregateType() {
	outFile := fmt.Sprintf("%s/register/register_aggregate_type.go", b.outDir)
	build := NewBuildRegisterAggregateType(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initBaseEvent() {
	outFile := fmt.Sprintf("%s/base/domain/event/base_event.go", b.outDir)
	build := NewBuildBaseEvent(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initBuildDomainBaseCmdService() {
	outFile := fmt.Sprintf("%s/base/domain/service/base_command_service.go", b.outDir)
	build := NewBuildDomainBaseCmdService(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initBuildAppBaseQueryService() {
	outFile := fmt.Sprintf("%s/base/application/service/base_query_service.go", b.outDir)
	build := NewBuildAppBaseQueryService(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}

func (b *BuildInfrastructureLayer) initBuildUtilsAssembler() {
	outFile := fmt.Sprintf("%s/utils/assembler_util.go", b.outDir)
	build := NewBuildUtilsAssembler(b.BaseBuild, utils.ToLower(outFile))
	b.AddBuild(build)
}
