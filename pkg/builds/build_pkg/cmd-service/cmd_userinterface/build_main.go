package cmd_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildRestControllerLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildRestControllerLayer {
	res := &BuildRestControllerLayer{
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

func (b *BuildRestControllerLayer) init() {
	var dirs []string
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/dto", b.outDir, b.aggregate.FileName()))
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/assembler", b.outDir, b.aggregate.FileName()))
	b.Mkdir(dirs...)

	outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildRestAggregateApi := NewBuildRestApiAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildRestAggregateApi)

	for _, entity := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		entityApi := NewBuildRestApiEntity(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		b.AddBuild(entityApi)
	}

	// dto
	outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildDtoAggregate := NewBuildDtoAggregate(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildDtoAggregate)

	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		buildDtoEntity := NewBuildDtoCommand(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		b.AddBuild(buildDtoEntity)
	}

	// assembler
	outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildAssemblerAggregate := NewBuildAssemblerAggregate(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildAssemblerAggregate)

	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		buildAssemblerEntity := NewBuildAssemblerEntity(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		b.AddBuild(buildAssemblerEntity)
	}

	// swagger
	outFile = fmt.Sprintf("%s/rest/swagger.go", b.outDir)
	buildSwagger := NewBuildSwagger(b.BaseBuild, outFile)
	b.AddBuild(buildSwagger)

}
