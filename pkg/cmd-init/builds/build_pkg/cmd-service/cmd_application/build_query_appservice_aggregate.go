package cmd_application

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildQueryAppServiceAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildQueryAppServiceAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildQueryAppServiceAggregate {
	res := &BuildQueryAppServiceAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/application/internals/service/query_appservice_aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildQueryAppServiceAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Name"] = b.aggregate.Name
	res["name"] = utils.FirstLower(b.aggregate.Name)
	res["ResourceName"] = b.aggregate.SnakeName()
	res["Description"] = b.aggregate.Description
	return res
}
