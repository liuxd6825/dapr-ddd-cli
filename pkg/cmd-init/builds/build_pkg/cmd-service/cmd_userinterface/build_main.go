package cmd_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerLayer struct {
	builds.BaseBuild
	aggregate             *config.Aggregate
	outDir                string
	buildRestAggregateApi *BuildRestAggregateApi
	buildRestEntityApis   *[]BuildRestEntityApi
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
	b.buildRestAggregateApi = NewBuildRestAggregateApi(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	var buildEntityApis []BuildRestEntityApi
	for _, entity := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), entity.FileName())
		entityApi := NewBuildRestEntityApi(b.BaseBuild, b.aggregate, entity, utils.ToLower(outFile))
		buildEntityApis = append(buildEntityApis, *entityApi)
	}
	b.buildRestEntityApis = &buildEntityApis
}

func (b *BuildRestControllerLayer) Build() error {
	var list []builds.Build
	list = append(list, b.buildRestAggregateApi)

	// valueObject
	buildEntityApis := func() []builds.Build {
		var res []builds.Build
		for _, item := range *b.buildRestEntityApis {
			res = append(res, &item)
		}
		return res
	}
	list = append(list, buildEntityApis()...)

	return b.DoBuild(list...)
}
