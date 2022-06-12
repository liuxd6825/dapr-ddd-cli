package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/cmd-init/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildApiBase struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildApiBase(base builds.BaseBuild, outFile string) *BuildApiBase {
	res := &BuildApiBase{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/userinterface/rest/facade/base_api.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildApiBase) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
