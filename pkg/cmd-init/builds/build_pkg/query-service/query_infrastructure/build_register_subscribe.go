package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildRegisterSubscribe struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildRegisterSubscribe(base builds.BaseBuild, outFile string) *BuildRegisterSubscribe {
	res := &BuildRegisterSubscribe{
		BaseBuild: base,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/register/register_subscribe.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildRegisterSubscribe) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
