package cmd_service

import (
	"github.com/dapr/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/dapr/dapr-ddd-cli/pkg/config"
)

type BuildRegisterEventType struct {
	builds.BaseBuild
	name      string
	aggregate *config.Aggregate
	dir       string
}

func NewBuildRegisterAllEventType(base builds.BaseBuild, aggregate *config.Aggregate, dir string) *BuildRegisterEventType {
	res := &BuildRegisterEventType{
		BaseBuild: base,
		aggregate: aggregate,
		dir:       dir,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/reg_all_event_type.go.tpl"
	res.OutFile = ""
	res.ValuesFunc = res.Values
	return res
}

func NewBuildRegisterAggregateEventType(base builds.BaseBuild, aggregate *config.Aggregate, dir string) *BuildRegisterEventType {
	res := &BuildRegisterEventType{
		BaseBuild: base,
		aggregate: aggregate,
		dir:       dir,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/events/reg_event_type.go.tpl"
	res.OutFile = ""
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRegisterEventType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Events"] = b.aggregate.Events
	return res
}
