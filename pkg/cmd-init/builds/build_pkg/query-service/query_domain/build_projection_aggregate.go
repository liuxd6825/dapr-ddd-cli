package query_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
)

type BuildProjectionAggregate struct {
	builds.BaseBuild
	name      string
	aggregate *config.Aggregate
	values    interface{}
}

func NewBuildProjectionAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildProjectionAggregate {
	res := &BuildProjectionAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/view/aggregate_view.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildProjectionAggregate) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	defaultProperties := config.NewProperties(b.Aggregate, b.Config.GetDefaultAggregateProperties(), &b.aggregate.Properties)
	res["name"] = utils.FirstLower(b.aggregate.Name)
	res["Name"] = utils.FirstUpper(b.aggregate.Name)
	res["ClassName"] = fmt.Sprintf("%sView", utils.FirstUpper(b.aggregate.Name))
	res["Properties"] = b.aggregate.Properties
	res["DefaultProperties"] = defaultProperties
	res["Description"] = b.aggregate.Description
	res["Aggregate"] = b.aggregate
	return res
}
