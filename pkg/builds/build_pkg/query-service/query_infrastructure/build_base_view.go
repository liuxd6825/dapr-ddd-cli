package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseView struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseView(base builds.BaseBuild, aggregate *config.Aggregate, outFile string) *BuildBaseView {
	res := &BuildBaseView{
		BaseBuild: base,
		aggregate: aggregate,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/domain/view/base_view.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseView) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Properties"] = b.Config.Configuration.DefaultReservedProperties.ViewProperties
	return res
}
