package cmd_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRegisterController struct {
	builds.BaseBuild
}

func NewBuildRegisterController(base builds.BaseBuild, outFile string) *BuildRegisterController {
	res := &BuildRegisterController{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/controller/register_controller.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRegisterController) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}

func (b *BuildRegisterController) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "Controller")
}
