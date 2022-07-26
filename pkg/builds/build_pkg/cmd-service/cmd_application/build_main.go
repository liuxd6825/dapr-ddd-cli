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

	buildCmdAppServiceAggregate *BuildCmdAppServiceAggregate
	//buildCmdAppServiceEntities  []*BuildCmdAppServiceEntity

	buildQueryAppServiceAggregate *BuildQueryAppServiceAggregate
	buildQueryAppServiceEntities  []*BuildQueryAppServiceEntity

	buildAssemblerAggregate *BuildAssemblerAggregate
	buildAssemblerEntities  []*BuildAssemblerEntity

	buildExecutorInit *BuildExecutorInit

	buildFindAggregateByIdExecutor *BuildFindAggregateByIdExecutor
	buildCommandExecutors          []*BuildCommandExecutor

	buildDtoAggregate *BuildDtoAggregate
	buildDtoEntities  []*BuildDtoEntity
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
	b.buildCmdAppServiceAggregate = NewBuildCmdAppServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	/*	b.buildCmdAppServiceEntities = []*BuildCmdAppServiceEntity{}
		for _, entity := range b.aggregate.Entities {
			outFile = fmt.Sprintf("%s/internals/%s/service/%s_command_app_service.go", b.outDir, aggregateName, entity.SnakeName())
			build := NewBuildCmdAppServiceEntity(b.BaseBuild, b.aggregate, entity, outFile)
			b.buildCmdAppServiceEntities = append(b.buildCmdAppServiceEntities, build)
		}*/

	// query service
	outFile = fmt.Sprintf("%s/internals/%s/service/%s_query_app_service.go", b.outDir, aggregateName, aggregateName)
	b.buildQueryAppServiceAggregate = NewBuildQueryAppServiceAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	/*	b.buildQueryAppServiceEntities = []*BuildQueryAppServiceEntity{}
		for _, entity := range b.aggregate.Entities {
			outFile = fmt.Sprintf("%s/internals/%s/service/%s_query_app_service.go", b.outDir, aggregateName, entity.SnakeName())
			build := NewBuildQueryAppServiceEntity(b.BaseBuild, entity, outFile)
			b.buildQueryAppServiceEntities = append(b.buildQueryAppServiceEntities, build)
		}*/

	// dto
	outFile = fmt.Sprintf("%s/internals/%s/dto/%s_dto.go", b.outDir, aggregateName, aggregateName)
	b.buildDtoAggregate = NewBuildDtoAggregate(b.BaseBuild, b.aggregate, outFile)

	b.buildDtoEntities = []*BuildDtoEntity{}
	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/internals/%s/dto/%s_dto.go", b.outDir, aggregateName, entity.SnakeName())
		build := NewBuildDtoEntity(b.BaseBuild, b.aggregate, entity, outFile)
		b.buildDtoEntities = append(b.buildDtoEntities, build)
	}

	// assembler
	outFile = fmt.Sprintf("%s/internals/%s/assembler/%s_assembler.go", b.outDir, aggregateName, aggregateName)
	b.buildAssemblerAggregate = NewBuildAssemblerAggregate(b.BaseBuild, b.aggregate, outFile)

	b.buildAssemblerEntities = []*BuildAssemblerEntity{}
	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/internals/%s/assembler/%s_assembler.go", b.outDir, aggregateName, entity.SnakeName())
		build := NewBuildAssemblerEntity(b.BaseBuild, b.aggregate, entity, outFile)
		b.buildAssemblerEntities = append(b.buildAssemblerEntities, build)
	}

	b.buildCommandExecutors = []*BuildCommandExecutor{}
	for _, command := range b.aggregate.Commands {
		outFile = fmt.Sprintf("%s/internals/%s/executor/%s_executor.go", b.outDir, aggregateName, utils.SnakeString(command.Name))
		build := NewBuildExecutor(b.BaseBuild, b.aggregate, command, outFile)
		b.buildCommandExecutors = append(b.buildCommandExecutors, build)
	}
	outFile = fmt.Sprintf("%s/internals/%s/executor/x_init.go", b.outDir, aggregateName)
	b.buildExecutorInit = NewBuildExecutorInit(b.BaseBuild, b.aggregate, outFile)

	outFile = fmt.Sprintf("%s/internals/%s/executor/%s_executor.go", b.outDir, aggregateName, "find_aggreage_id")
	b.buildFindAggregateByIdExecutor = NewBuildFindAggregateByIdExecutor(b.BaseBuild, b.Aggregate, outFile)
}

func (b *BuildApplicationLayer) Build() error {
	var list []builds.Build

	// command service
	list = append(list, b.buildCmdAppServiceAggregate)
	/*	for _, item := range b.buildCmdAppServiceEntities {
		list = append(list, item)
	}*/

	// executor
	list = append(list, b.buildFindAggregateByIdExecutor)
	list = append(list, b.buildExecutorInit)
	for _, item := range b.buildCommandExecutors {
		list = append(list, item)
	}

	// assembler
	list = append(list, b.buildAssemblerAggregate)
	for _, item := range b.buildAssemblerEntities {
		list = append(list, item)
	}

	// dto
	list = append(list, b.buildDtoAggregate)
	for _, item := range b.buildDtoEntities {
		list = append(list, item)
	}
	return b.DoBuild(list...)
}
