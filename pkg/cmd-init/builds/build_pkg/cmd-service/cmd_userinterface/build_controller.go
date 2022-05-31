package cmd_userinterface

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestController struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRestController(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRestController {
	res := &BuildRestController{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/userinterface/rest/controller/controller.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestController) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["ClassName"] = b.ClassName()
	res["Events"] = b.aggregate.Events
	res["Commands"] = b.aggregate.Commands
	res["Resource"] = utils.MidlineString(b.aggregate.Name)

	return res
}

func (b *BuildRestController) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "Controller")
}

func (b *BuildRestController) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
