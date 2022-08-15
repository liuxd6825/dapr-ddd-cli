package cmd_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildApplicationLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildApplicationLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildApplicationLayer {
	res := &BuildApplicationLayer{
		BaseBuild: builds.BaseBuild{
			Config:    cfg,
			Aggregate: aggregate,
		},
		aggregate: aggregate,
		outDir:    outDir,
	}
	res.init()
	return res
}

func (b *BuildApplicationLayer) init() {
	aggregateName := b.aggregate.SnakeName()

	// command service
	outFile := fmt.Sprintf("%s/internals/%s/service/%s_command_app_service.go", b.outDir, aggregateName, aggregateName)
	buildCmdAppService := NewBuildCmdAppService(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildCmdAppService)

	// dto
	outFile = fmt.Sprintf("%s/internals/%s/appcmd/%s_appcmd.go", b.outDir, aggregateName, aggregateName)
	buildDtoAggregate := NewBuildAppCmd(b.BaseBuild, nil, outFile)
	b.AddBuild(buildDtoAggregate)

	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/internals/%s/appcmd/%s_appcmd.go", b.outDir, aggregateName, entity.SnakeName())
		build := NewBuildAppCmd(b.BaseBuild, entity, outFile)
		b.AddBuild(build)
	}

	// assembler
	outFile = fmt.Sprintf("%s/internals/%s/assembler/%s_assembler.go", b.outDir, aggregateName, aggregateName)
	buildAssemblerAggregate := NewBuildAssembler(b.BaseBuild, nil, outFile)
	b.AddBuild(buildAssemblerAggregate)

	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/internals/%s/assembler/%s_assembler.go", b.outDir, aggregateName, entity.SnakeName())
		build := NewBuildAssembler(b.BaseBuild, entity, outFile)
		b.AddBuild(build)
	}

	for _, command := range b.aggregate.Commands {
		outFile = fmt.Sprintf("%s/internals/%s/executor/%s_executor.go", b.outDir, aggregateName, utils.SnakeString(command.Name))
		build := NewBuildExecutor(b.BaseBuild, b.aggregate, command, outFile)
		b.AddBuild(build)
	}

	outFile = fmt.Sprintf("%s/internals/%s/executor/x_init.go", b.outDir, aggregateName)
	buildExecutorInit := NewBuildExecutorInit(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildExecutorInit)

	outFile = fmt.Sprintf("%s/internals/%s/executor/%s_executor.go", b.outDir, aggregateName, "find_aggreage_id")
	buildFindAggregateByIdExecutor := NewBuildFindAggregateByIdExecutor(b.BaseBuild, b.Aggregate, outFile)
	b.AddBuild(buildFindAggregateByIdExecutor)
}
