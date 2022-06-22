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

	buildRestApiAggregate *BuildRestApiAggregate
	buildRestApiEntities  []*BuildRestApiEntity

	buildDtoAggregate *BuildDtoAggregate
	buildDtoEntities  *[]*BuildDtoEntity

	buildAssemblerAggregate *BuildAssemblerAggregate
	buildAssemblerEntities  *[]*BuildAssemblerEntity

	buildSwagger *BuildSwagger
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

func (b *BuildUserInterfaceLayer) Build() error {
	var list []builds.Build

	// aggregate
	list = append(list, b.buildRestApiAggregate)

	// entityObject
	for _, item := range b.buildRestApiEntities {
		list = append(list, item)
	}

	// dto
	list = append(list, b.buildDtoAggregate)
	for _, item := range *b.buildDtoEntities {
		list = append(list, item)
	}

	// assembler
	list = append(list, b.buildAssemblerAggregate)
	for _, item := range *b.buildAssemblerEntities {
		list = append(list, item)
	}

	// swagger
	list = append(list, b.buildSwagger)

	return b.DoBuild(list...)
}

func (b *BuildUserInterfaceLayer) init() {
	outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildRestApiAggregate = NewBuildRestApiAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	var dirs []string
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/dto", b.outDir, b.aggregate.FileName()))
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/assembler", b.outDir, b.aggregate.FileName()))
	b.Mkdir(dirs...)

	b.buildRestApiEntities = []*BuildRestApiEntity{}
	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRestApiEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.buildRestApiEntities = append(b.buildRestApiEntities, build)
	}

	outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildDtoAggregate = NewBuildDtoAggregate(b.BaseBuild, b.aggregate, outFile)

	var buildDtoEntities []*BuildDtoEntity
	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildDtoEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		buildDtoEntities = append(buildDtoEntities, build)
	}
	b.buildDtoEntities = &buildDtoEntities

	outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildAssemblerAggregate = NewBuildAssemblerAggregate(b.BaseBuild, b.aggregate, outFile)

	var buildAssemblerEntities []*BuildAssemblerEntity
	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildAssemblerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		buildAssemblerEntities = append(buildAssemblerEntities, build)
	}
	b.buildAssemblerEntities = &buildAssemblerEntities

	// swagger
	outFile = fmt.Sprintf("%s/rest/swagger.go", b.outDir)
	b.buildSwagger = NewBuildSwagger(b.BaseBuild, outFile)

}
