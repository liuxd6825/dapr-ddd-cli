package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseApi struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseApi(base builds.BaseBuild, outFile string) *BuildBaseApi {
	res := &BuildBaseApi{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/userinterface/rest/facade/base_api.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseApi) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	return res
}
