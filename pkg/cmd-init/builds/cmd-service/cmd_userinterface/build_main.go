package cmd_userinterface

import (
	"fmt"
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
	"github.com/dapr/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerLayer struct {
	builds.BaseBuild
	aggregate           *config.Aggregate
	outDir              string
	buildRestController *BuildRestController
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
	outFile := fmt.Sprintf("%s/rest/controller/%s_controller.go", b.outDir, b.aggregate.Name)
	b.buildRestController = NewBuildRestController(b.BaseBuild, b.aggregate, utils.ToLower(outFile))
}

func (b *BuildRestControllerLayer) Build() error {
	list := []builds.Build{}

	list = append(list, b.buildRestController)
	return b.doBuild(list...)
}

func (b *BuildRestControllerLayer) doBuild(builds ...builds.Build) error {
	if builds == nil {
		return nil
	}
	for _, build := range builds {
		if err := build.Build(); err != nil {
			return err
		}
	}
	return nil
}
