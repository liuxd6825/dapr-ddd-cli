package cmd_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
)

type BuildEventFactory struct {
	builds.BaseBuild
}

func NewBuildEventFactory(base builds.BaseBuild, outFile string) *BuildEventFactory {
	res := &BuildEventFactory{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/factory/event_factory.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildEventFactory) Values() map[string]interface{} {
	values := b.BaseBuild.Values()
	values["Commands"] = b.Aggregate.Commands
	return values
}
