package query_userinterface

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildUserInterfaceLayer struct {
	builds.BaseBuild
	aggregate *config.Aggregate
	outDir    string

	buildRestControllerAggregate *BuildRestControllerAggregate
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

	res.initBuildRestControllerAggregate()
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

func (b *BuildUserInterfaceLayer) initBuildRestControllerAggregate() {
	outFile := fmt.Sprintf("%s/rest/controller/%s_controller.go", b.outDir, b.aggregate.Name)
	b.buildRestControllerAggregate = NewBuildRestControllerAggregate(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildUserInterfaceLayer) initQueryServiceEntities() {
	b.buildRestControllerEntities = []*BuildRestControllerEntity{}
	for _, item := range b.aggregate.Entities {
		outFile := fmt.Sprintf("%s/rest/controller/%s_controller.go", b.outDir, item.Name)
		build := NewBuildRestControllerEntity(b.BaseBuild, item, utils.ToLower(outFile))
		b.buildRestControllerEntities = append(b.buildRestControllerEntities, build)
	}
}
