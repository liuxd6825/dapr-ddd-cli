package cmd_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
)

type BuildRegisterEventType struct {
	builds.BaseBuild
}

func NewBuildRegisterEventType(base builds.BaseBuild, outFile string) *BuildRegisterEventType {
	res := &BuildRegisterEventType{
		BaseBuild: base,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/cmd-service/infrastructure/register/register_event_type.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRegisterEventType) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
