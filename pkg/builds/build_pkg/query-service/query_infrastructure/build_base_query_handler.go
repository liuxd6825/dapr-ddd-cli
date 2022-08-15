package query_infrastructure

import (
	"github.com/liuxd6825/dapr-ddd-cli/pkg/builds"
	"github.com/liuxd6825/dapr-ddd-cli/pkg/config"
)

type BuildBaseQueryHandler struct {
	builds.BaseBuild
	aggregate *config.Aggregate
}

func NewBuildBaseQueryHandler(base builds.BaseBuild, outFile string) *BuildBaseQueryHandler {
	res := &BuildBaseQueryHandler{
		BaseBuild: base,
	}
	res.ValuesFunc = res.Values
	res.TmplFile = "static/tmpl/go/init/pkg/query-service/infrastructure/base/application/handler/base_query_handler.go.tpl"
	res.OutFile = outFile
	return res
}

func (b *BuildBaseQueryHandler) Values() map[string]interface{} {
	res := b.BaseBuild.Values()
	res["Properties"] = b.Config.Configuration.DefaultReservedProperties.ViewProperties
	return res
}
