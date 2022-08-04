package cmd_application

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildCmdAppService struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildCmdAppService(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildCmdAppService {
	res := &BuildCmdAppService{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/service/cmd_app_service.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildCmdAppService) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Commands"] = b.aggregate.Commands
	res["ClassName"] = b.ClassName()
	return res
}

func (b *BuildCmdAppService) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "CommandAppService")
}

func (b *BuildCmdAppService) AggregateType() string {
	return utils.FirstUpper(fmt.Sprintf("%s.%s", b.Namespace(), b.ClassName()))
}
