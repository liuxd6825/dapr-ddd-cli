package query_domain

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildQueryHandlerRegisterSubscribe struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildQueryHandlerRegisterSubscribe(base builds.BaseBuild, outFile string) *BuildQueryHandlerRegisterSubscribe {
	res := &BuildQueryHandlerRegisterSubscribe{
		BaseBuild: base,
	}
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/domain/handler/register_subscribe.go.tpl"
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}

func (b *BuildQueryHandlerRegisterSubscribe) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
