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
	res["AggregateType"] = b.AggregateType()
	res["Properties"] = b.aggregate.Properties
	res["Events"] = b.aggregate.Events
	res["Commands"] = b.aggregate.Commands
	res["Description"] = b.aggregate.Description
	res["EnumObjects"] = b.aggregate.EnumObjects
	res["Id"] = b.aggregate.Id
	res["FieldsObjects"] = b.aggregate.FieldsObjects
	res["Aggregate"] = b.aggregate
	res["CommandPackage"] = fmt.Sprintf("%s_command", utils.ToLower(b.aggregate.Name))
	res["EventPackage"] = fmt.Sprintf("%s_event", utils.ToLower(b.aggregate.Name))
	res["Package"] = fmt.Sprintf("%s_model", utils.ToLower(b.aggregate.Name))
	res["Version"] = b.aggregate.Version
	res["AppService"] = fmt.Sprintf("%sAppService", b.aggregate.LowerName())

	return res
}

func (b *BuildRestController) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "Controller")
}

func (b *BuildRestController) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
