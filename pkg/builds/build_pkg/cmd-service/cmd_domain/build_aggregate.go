package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildAggregate struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildAggregate(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildAggregate {
	res := &BuildAggregate{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/aggregate.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAggregate) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	values["ClassName"] = b.ClassName()
	values["AggregateType"] = b.AggregateType()
	values["Properties"] = b.aggregate.Properties
	values["Commands"] = b.aggregate.Commands
	values["Events"] = b.aggregate.Events
	b.AddTimePackageValue(values, &b.aggregate.Properties)
	return values
}

func (b *BuildAggregate) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "Aggregate")
}

func (b *BuildAggregate) AggregateType() string {
	return fmt.Sprintf("%s.%s", strings.ToLower(b.Config.Configuration.ServiceName), b.ClassName())
}
