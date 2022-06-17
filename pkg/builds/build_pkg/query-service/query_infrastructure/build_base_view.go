package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildViewBase struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildViewBase(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildViewBase {
	res := &BuildViewBase{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/domain/view/base_view.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildViewBase) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Properties"] = b.Config.Configuration.DefaultReservedProperties.ViewProperties
	return res
}
