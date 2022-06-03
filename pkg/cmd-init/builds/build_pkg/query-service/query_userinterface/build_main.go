package query_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildUserInterfaceLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string

	buildRestControllerAggregate *BuildRestApiAggregate
	buildRestControllerEntities  []*BuildRestControllerEntity
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
	list = append(list, b.buildRestControllerAggregate)

	// entityObject
	buildQueryServiceImplEntities := func() []builds.Build {
		var res []builds.Build
		for _, item := range b.buildRestControllerEntities {
			res = append(res, item)
		}
		return res
	}
	list = append(list, buildQueryServiceImplEntities()...)

	return b.DoBuild(list...)
}

func (b *BuildUserInterfaceLayer) init() {
	outFile := fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildRestControllerAggregate = NewBuildRestApiAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	var dirs []string
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/dto", b.outDir, b.aggregate.FileName()))
	dirs = append(dirs, fmt.Sprintf("%s/rest/%s/assembler", b.outDir, b.aggregate.FileName()))
	b.Mkdir(dirs...)

	b.buildRestControllerEntities = []*BuildRestControllerEntity{}
	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/%s/facade/%s_api.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRestControllerEntity(b.BaseBuild, b.aggregate, item, utils.ToLower(outFile))
		b.buildRestControllerEntities = append(b.buildRestControllerEntities, build)
	}
}
