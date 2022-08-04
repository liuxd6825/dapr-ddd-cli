package query_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildUserInterfaceLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string
}

func NewBuildUserInterfaceLayer(cfg *config.Config, aggregate *config.Aggregate, outDir string) *BuildUserInterfaceLayer {
	res := &BuildUserInterfaceLayer{
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

func (b *BuildUserInterfaceLayer) init() {
	outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildRestApiAggregate := NewBuildRestApiAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
	b.AddBuild(buildRestApiAggregate)

	var dirs []string
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/dto", b.outDir, b.aggregate.FileName()))
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/assembler", b.outDir, b.aggregate.FileName()))
	b.Mkdir(dirs...)

	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRestApiEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}

	outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildDtoAggregate := NewBuildDtoAggregate(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildDtoAggregate)

	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildDtoEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}

	outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	buildAssemblerAggregate := NewBuildAssemblerAggregate(b.BaseBuild, b.aggregate, outFile)
	b.AddBuild(buildAssemblerAggregate)

	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildAssemblerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.AddBuild(build)
	}

	// swagger
	outFile = fmt.Sprintf("%s/rest/swagger.go", b.outDir)
	buildSwagger := NewBuildSwagger(b.BaseBuild, outFile)
	b.AddBuild(buildSwagger)
}
