package cmd_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildEventType struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildEventType(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildEventType {
	res := &BuildEventType{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/event_type.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildEventType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Events"] = b.aggregate.Events
	res["EventTypes"] = b.aggregate.Events.GetEventTypes()
	return res
}
