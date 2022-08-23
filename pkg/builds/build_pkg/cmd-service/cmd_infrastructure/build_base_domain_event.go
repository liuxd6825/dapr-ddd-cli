package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseEvent struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseEvent(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildBaseEvent {
	res := &BuildBaseEvent{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/base/domain/event/base_event.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseEvent) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
