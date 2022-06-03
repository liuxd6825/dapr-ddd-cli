package cmd_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildRestControllerLayer struct {
	builds.BaseBuild
	aggregate             *config.Aggregate
	outDir                string
	buildRestAggregateApi *BuildRestApiAggregate
	buildRestEntityApis   []*BuildRestApiEntity
	buildDtoCommands      []*BuildDtoCommand
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

	var buildDtoCommands []*BuildDtoCommand
	for _, command := range b.aggregate.Commands {
		fileName := strings.Replace(command.FileName(), "_command", "", 1)
		outFile := fmt.Sprintf("%s/rest/%s/dto/%s_dto.go", b.outDir, b.aggregate.FileName(), fileName)
		buildDtoCommand := NewBuildDtoCommand(b.BaseBuild, b.aggregate, command, utils.ToLower(outFile))
		buildDtoCommands = append(buildDtoCommands, buildDtoCommand)
	}
	b.buildDtoCommands = buildDtoCommands

}

func (b *BuildRestControllerLayer) Build() error {
	var list []builds.Build
	list = append(list, b.buildRestAggregateApi)

	// api
	for _, item := range b.buildRestEntityApis {
		list = append(list, item)
	}

	// dto
	for _, dto := range b.buildDtoCommands {
		list = append(list, dto)
	}

	return b.DoBuild(list...)
}
