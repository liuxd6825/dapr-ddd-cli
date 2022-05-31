package cmd_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
)

type BuildRegisterAllEventType struct {
	builds.BaseBuild
}

func NewBuildRegisterAllEventType(base builds.BaseBuild, outFile string) *BuildRegisterAllEventType {
	res := &BuildRegisterAllEventType{
		BaseBuild: base,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/domain/event/register_event_type.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRegisterAllEventType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
