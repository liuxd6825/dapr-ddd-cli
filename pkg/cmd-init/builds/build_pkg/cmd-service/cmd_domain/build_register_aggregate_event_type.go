package cmd_domain

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildRegisterAggregateEventType struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRegisterAggregateEventType(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildRegisterAggregateEventType {
	res := &BuildRegisterAggregateEventType{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/event/reg_event_type.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRegisterAggregateEventType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Events"] = b.aggregate.Events

	res["EventTypes"] = b.aggregate.Events.GetEventTypes()
	return res
}
