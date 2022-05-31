package cmd_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerLayer struct {
	builds.BaseBuild
	aggregate               *config.Aggregate
	outDir                  string
	buildRestController     *BuildRestController
	buildRegisterController *BuildRegisterController
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
	outFile := fmt.Sprintf("%s/rest/controller/%s_controller.go", b.outDir, b.aggregate.FileName())
	b.buildRestController = NewBuildRestController(b.BaseBuild, b.aggregate, utils.ToLower(outFile))

	outFile = fmt.Sprintf("%s/rest/controller/register_controller.go", b.outDir)
	b.buildRegisterController = NewBuildRegisterController(b.BaseBuild, utils.ToLower(outFile))
}

func (b *BuildRestControllerLayer) Build() error {
	var list []builds.Build

	list = append(list, b.buildRestController)
	list = append(list, b.buildRegisterController)

	return b.DoBuild(list...)
}
