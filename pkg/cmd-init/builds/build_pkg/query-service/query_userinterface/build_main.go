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

	buildRestControllerAggregate *BuildRestControllerAggregate
	buildRestControllerEntities  []*BuildRestControllerEntity
	buildRegisterController      *BuildRegisterController
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

	// registerController
	list = append(list, b.buildRegisterController)
	return b.DoBuild(list...)
}

func (b *BuildUserInterfaceLayer) init() {
	outFile := fmt.Sprintf("%s/rest/controller/%s_controller/%s_controller.go", b.outDir, b.aggregate.FileName(), b.aggregate.FileName())
	b.buildRestControllerAggregate = NewBuildRestControllerAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	outFile = fmt.Sprintf("%s/rest/controller/register_controller.go", b.outDir)
	b.buildRegisterController = NewBuildRegisterController(b.BaseBuild, outFile)

	b.buildRestControllerEntities = []*BuildRestControllerEntity{}
	for _, item := range b.aggregate.Entities {
		outFile = fmt.Sprintf("%s/rest/controller/%s_controller/%s_controller.go", b.outDir, b.aggregate.FileName(), item.FileName())
		build := NewBuildRestControllerEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRestControllerEntities = append(b.buildRestControllerEntities, build)
	}
}
