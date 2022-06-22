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

	buildAssemblerAggregate *BuildAssemblerAggregate
	buildAssemblerEntities  []*BuildAssemblerEntity

	buildRestAggregateApi *BuildRestApiAggregate
	buildRestEntityApis   []*BuildRestApiEntity

	buildDtoAggregate *BuildDtoAggregate
	buildDtoEntities  []*BuildDtoEntity

	buildSwagger *BuildSwagger
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
	b.buildRestAggregateApi = NewBuildRestApiAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	var buildEntityApis []*BuildRestApiEntity
	for _, entity := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		entityApi := NewBuildRestApiEntity(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		buildEntityApis = append(buildEntityApis, entityApi)
	}
	b.buildRestEntityApis = buildEntityApis

	// dto
	outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildDtoAggregate = NewBuildDtoAggregate(b.BaseBuild, b.aggregate, outFile)

	var buildDtoEntities []*BuildDtoEntity
	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		buildDtoEntity := NewBuildDtoCommand(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		buildDtoEntities = append(buildDtoEntities, buildDtoEntity)
	}
	b.buildDtoEntities = buildDtoEntities

	// assembler
	outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildAssemblerAggregate = NewBuildAssemblerAggregate(b.BaseBuild, b.aggregate, outFile)

	var buildAssemblerEntities []*BuildAssemblerEntity
	for _, entity := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/assembler/%s_assembler.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		buildAssemblerEntity := NewBuildAssemblerEntity(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		buildAssemblerEntities = append(buildAssemblerEntities, buildAssemblerEntity)
	}
	b.buildAssemblerEntities = buildAssemblerEntities

	// swagger
	outFile = fmt.Sprintf("%s/rest/swagger.go", b.outDir)
	b.buildSwagger = NewBuildSwagger(b.BaseBuild, outFile)

}

func (b *BuildRestControllerLayer) Build() error {
	var list []builds.Build

	// api
	list = append(list, b.buildRestAggregateApi)
	for _, item := range b.buildRestEntityApis {
		list = append(list, item)
	}

	// dto
	for _, dto := range b.buildDtoEntities {
		list = append(list, dto)
	}
	list = append(list, b.buildDtoAggregate)

	// assembler
	for _, entity := range b.buildAssemblerEntities {
		list = append(list, entity)
	}
	list = append(list, b.buildAssemblerAggregate)

	// swagger
	list = append(list, b.buildSwagger)

	return b.DoBuild(list...)
}
