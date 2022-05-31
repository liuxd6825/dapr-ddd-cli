package query_userinterface

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildRestControllerAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRestControllerAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRestControllerAggregate {
	res := &BuildRestControllerAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/userinterface/rest/controller/controller_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildRestControllerAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.aggregate.Name
	res["ResourceName"] = utils.MidlineString(b.aggregate.Name)
	res["ServiceName"] = b.aggregate.FirstUpperName() + "AppService"

	return res
}
