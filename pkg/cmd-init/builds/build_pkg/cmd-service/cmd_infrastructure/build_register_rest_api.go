package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRegisterRestController struct {
	builds.BaseBuild
}

func NewBuildRegisterRestController(base builds.BaseBuild, outFile string) *BuildRegisterRestController {
	res := &BuildRegisterRestController{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/register/register_rest_api.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRegisterRestController) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}

func (b *BuildRegisterRestController) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "RestApi")
}
