package cmd_domain

import (
	"fmt"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/utils"
	"strings"
)

type BuildAggregateEvent struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildAggregateEvent(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildAggregateEvent {
	res := &BuildAggregateEvent{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/model/aggregate/aggregate_event.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildAggregateEvent) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	values["ClassName"] = b.ClassName()
	values["AggregateType"] = b.AggregateType()
	values["Properties"] = b.aggregate.Properties
	values["Commands"] = b.aggregate.Commands
	values["Events"] = b.aggregate.Events
	b.AddTimePackageValue(values, &b.aggregate.Properties)
	return values
}

func (b *BuildAggregateEvent) ClassName() string {
	return utils.FirstUpper(b.AggregateName() + "Aggregate")
}

func (b *BuildAggregateEvent) AggregateType() string {
	return fmt.Sprintf("%s.%s", strings.ToLower(b.Config.Configuration.ServiceName), b.ClassName())
}
